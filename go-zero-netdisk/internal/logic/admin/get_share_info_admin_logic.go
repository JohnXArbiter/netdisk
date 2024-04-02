package admin

import (
	"context"
	"lc/netdisk/common/constant"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShareInfoAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShareInfoAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShareInfoAdminLogic {
	return &GetShareInfoAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShareInfoAdminLogic) GetShareInfoAdmin(req *types.IdStrReq) (*types.GetShareInfoResp, error) {
	var (
		engine   = l.svcCtx.Xorm
		minioSvc = l.svcCtx.MinioSvc
		files    []*model.File
	)

	if err := engine.Where(
		"id in (select file_id from share_file where share_id = ?)",
		req.Id).Find(&files); err != nil {
		logx.Errorf("GetShareInfoAdmin，查询失败，ERR: [%v]", err)
		return nil, err
	}

	resp := &types.GetShareInfoResp{}
	for _, file := range files {
		url, err := minioSvc.GenUrl(file.ObjectName, file.Name, true)
		if err != nil {
			logx.Errorf("GetShareInfoAdmin，生成文件url出错：ERR: [%v]", err)
			url = ""
		}
		resp.Items = append(resp.Items, &types.ListShareItemStruct{
			Id:      file.Id,
			Name:    file.Name,
			Updated: file.Updated.Format(constant.TimeFormat1),
			Size:    file.Size,
			Url:     url,
			Status:  file.Status,
			Type:    file.Type,
			DelFlag: file.DelFlag,
		})
	}

	return resp, nil
}
