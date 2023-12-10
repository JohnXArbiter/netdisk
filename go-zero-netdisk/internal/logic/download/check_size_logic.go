package download

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lc/netdisk/common/constant"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"
)

type CheckSizeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckSizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckSizeLogic {
	return &CheckSizeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckSizeLogic) CheckSize(req *types.CheckSizeReq) (*types.CheckSizeResp, error) {
	var (
		userId = l.ctx.Value(constant.UserIdKey).(int64)
		engine = l.svcCtx.Xorm
		file   model.File
		fileFs model.FileFs
		resp   types.CheckSizeResp
	)

	if has, err := engine.ID(req.FileId).
		And("user_id = ?", userId).
		Get(&file); err != nil || !has {
		return nil, err
	}

	if file.IsBig == constant.SmallFileFlag {
		resp.IsBig = file.IsBig
		return &resp, nil
	}

	if has, err := engine.ID(file.FsId).
		Get(&fileFs); err != nil || !has {
		return nil, err
	}

	resp.IsBig = file.IsBig
	resp.ChunkNum = fileFs.ChunkNum
	return &resp, nil
}
