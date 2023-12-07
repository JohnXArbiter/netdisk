package download

import (
	"context"
	"lc/netdisk/common/constant"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"
	"math"

	"github.com/zeromicro/go-zero/core/logx"
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
		userId      = l.ctx.Value(constant.UserIdKey).(int64)
		engine      = l.svcCtx.Xorm
		fileNetdisk model.FileNetdisk
		fileFs      model.FileFs
		resp        types.CheckSizeResp
	)

	if has, err := engine.ID(req.FileNetdiskId).
		And("user_id = ?", userId).
		Get(&fileNetdisk); err != nil || !has {
		return nil, err
	}

	if fileNetdisk.IsBig == constant.SmallFileFlag {
		resp.IsBig = fileNetdisk.IsBig
		return &resp, nil
	}

	if has, err := engine.ID(fileNetdisk.FsId).
		Get(&fileFs); err != nil || !has {
		return nil, err
	}

	chunkCount := math.Ceil(float64(fileFs.Size / constant.NeedShardingSize))
	resp.IsBig = fileNetdisk.IsBig
	resp.ChunkCount = int64(chunkCount)
	return &resp, nil
}
