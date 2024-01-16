package file

import (
	"context"
	"lc/netdisk/common/constant"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFileLogic {
	return &ListFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFileLogic) ListFile(req *types.ParentFolderIdReq) ([]*types.ListFileStruct, error) {
	var (
		userId = l.ctx.Value(constant.UserIdKey).(int64)
		engine = l.svcCtx.Xorm
		files  []*model.File
		resp   []*types.ListFileStruct
	)

	if err := engine.Where("folder_id = ?", req.ParentFolderId).
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
