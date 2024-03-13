package admin

import (
	"context"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/xorm"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetShareStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetShareStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetShareStatusLogic {
	return &SetShareStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetShareStatusLogic) SetShareStatus(req *types.SetShareStatusReq) error {
	var (
		engine = l.svcCtx.Xorm
	)

	_, err := engine.DoTransaction(func(session *xorm.Session) (interface{}, error) {
		if req.Type != constant.TypeShareMulti {
			bean := &model.File{Status: constant.StatusFileIllegal}
			if _, err := session.Where("id = "+
				"(select file_id from share_file where share_id = ?)",
				req.Id).Update(bean); err != nil {
				logx.Errorf("SetShareStatus，更新文件失败，ERR: [%v]", err)
				return nil, err
			}
		}

		bean := &model.Share{Status: req.Status}
		if _, err := session.ID(req.Id).
			Cols("status").
			Update(bean); err != nil {
			logx.Errorf("SetStatus，更新失败，ERR: [%v]", err)
			return nil, err
		}
		return nil, nil
	})
	return err
}
