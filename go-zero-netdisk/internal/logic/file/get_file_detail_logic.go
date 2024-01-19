package file

import (
	"context"
	"errors"
	"lc/netdisk/common/constant"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileDetailLogic {
	return &GetFileDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFileDetailLogic) GetFileDetail(req *types.IdPathReq) (*types.FileResp, error) {
	var (
		userId = l.ctx.Value(constant.UserIdKey).(int64)
		engine = l.svcCtx.Xorm
		file   = &model.File{}
	)

	has, err := engine.ID(req.Id).And("user_id = ?", userId).Get(file)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("未能找到该文件信息！😿")
	}

	resp := &types.FileResp{}
	resp.Id = file.Id
	resp.Name = file.Name
	resp.Url = file.Url
	resp.Size = file.Id
	resp.Status = file.Status
	resp.FolderId = file.FolderId
	resp.Created = file.Created.Format(constant.TimeFormat1)
	resp.Updated = file.Updated.Format(constant.TimeFormat1)
	return resp, nil
}
