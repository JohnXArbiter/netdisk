package file

import (
	"context"
	"lc/netdisk/common/constant"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFileByTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFileByTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFileByTypeLogic {
	return &ListFileByTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFileByTypeLogic) ListFileByType(req *types.FileTypeReq) ([]*types.ListFileStruct, error) {
	var (
		userId = l.ctx.Value(constant.UserIdKey).(int64)
		engine = l.svcCtx.Xorm
		files  []*model.File
		resp   []*types.ListFileStruct
	)

	if err := engine.Where("type = ?", req.FileType).
		And("user_id = ?", userId).And("del_flag = ?",
		constant.StatusFileUndeleted).Find(&files); err != nil {
		return nil, err
	}

	for _, file := range files {
		resp = append(resp, &types.ListFileStruct{
			Id:      file.Id,
			Name:    file.Name,
			Url:     file.Url,
			Size:    file.Size,
			Status:  file.Status,
			Updated: file.Updated.Format(constant.TimeFormat1),
		})
	}
	return resp, nil
}
