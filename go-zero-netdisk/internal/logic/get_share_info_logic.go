package logic

import (
	"context"
	"errors"
	"lc/netdisk/common/constant"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShareInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShareInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShareInfoLogic {
	return &GetShareInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShareInfoLogic) GetShareInfo(req *types.GetShareInfoReq) (*types.GetShareInfoResp, error) {
	var (
		engine   = l.svcCtx.Xorm
		minioSvc = l.svcCtx.Minio.NewService()
		files    []*model.File
	)

	var share model.Share
	if has, err := engine.Where("id = ?", req.Id).Get(&share); err != nil {
		logx.Errorf("获取分享文件列表，查询share出错：ERR: [%v]", err)
		return nil, err
	} else if !has {
		return nil, errors.New("找不到分享列表")
	}

	if share.Pwd != req.Pwd {
		return nil, errors.New("提取码错误")
	}

	resp := &types.GetShareInfoResp{}
	resp.Name = share.Name
	resp.Created = share.Created.Format(constant.TimeFormat1)
	resp.Expired = share.Expired
	resp.Owner = share.UserId

	if err := engine.Where(
		"id in (select file_id from share_file where share_id = ?)",
		req.Id).Find(&files); err != nil {
		logx.Errorf("获取分享文件列表，查询文件出错：ERR: [%v]", err)
		return nil, err
	}

	for _, file := range files {
		url, err := minioSvc.GenUrl(file.ObjectName, true)
		if err != nil {
			logx.Errorf("获取分享文件列表，生成文件url出错：ERR: [%v]", err)
			url = ""
		}
		resp.Items = append(resp.Items, &types.ListShareItemStruct{
			Id:      file.Id,
			Name:    file.Name,
			Updated: file.Updated.Format(constant.TimeFormat1),
			Size:    file.Size,
			Url:     url,
			Status:  file.Status,
		})
	}

	return resp, nil
}
