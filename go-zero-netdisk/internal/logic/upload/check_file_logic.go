package upload

import (
	"context"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckFileLogic {
	return &CheckFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckFileLogic) CheckFile(req *types.CheckFileReq) (*types.CheckFileResp, error) {
	var (
		userId int64 = 123
		xorm         = l.svcCtx.Xorm
	)

	fileRepository := &model.FileRepository{
		Hash:   req.Hash,
		UserId: userId,
	}
	has, err := xorm.Get(fileRepository)
	if err != nil {
		// TODO: log
		return nil, err
	}
	if !has {

	}

	resp := &types.CheckFileResp{
		FileId: 0,
		Status: 0,
	}
	return resp, nil
}
