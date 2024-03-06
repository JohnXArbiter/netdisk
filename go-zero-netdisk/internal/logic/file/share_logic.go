package file

import (
	"context"
	"errors"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/variable"
	"lc/netdisk/common/xorm"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareLogic {
	return &ShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareLogic) Share(req *types.ShareReq) error {
	var (
		userId = l.ctx.Value(constant.UserIdKey).(int64)
		engine = l.svcCtx.Xorm
	)

	if req.Id == "" {
		return errors.New("出错啦，请重试")
	}

	_, err := engine.DoTransaction(func(session *xorm.Session) (interface{}, error) {
		var shareFile []*model.ShareFile
		for _, fileId := range req.FileIds {
			shareFile = append(shareFile, &model.ShareFile{
				ShareId: req.Id,
				FileId:  fileId,
			})
		}
		if _, err := session.Insert(shareFile); err != nil {
			logx.Errorf("分享多文件，插入shareFile失败，ERR: [%v]", err)
			return nil, err
		}

		created := time.Now().Unix()
		expired := created + variable.ShareExpireType[req.ExpireType]
		share := &model.Share{}
		share.Id = req.Id
		share.UserId = userId
		share.Created = created
		share.Expired = expired
		if _, err := session.Insert(share); err != nil {
			logx.Errorf("分享多文件，插入share失败，ERR: [%v]", err)
			return nil, err
		}
		return nil, nil
	})

	return err
}
