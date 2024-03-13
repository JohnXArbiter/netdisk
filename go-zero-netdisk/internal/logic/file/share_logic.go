package file

import (
	"context"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/variable"
	"lc/netdisk/common/xorm"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"
	"strconv"
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
		userId         = l.ctx.Value(constant.UserIdKey).(int64)
		engine         = l.svcCtx.Xorm
		shareType int8 = constant.TypeShareMulti
	)

	if req.Pwd == "" {
		return errors.New("出错啦，请重试")
	}

	var file model.File
	if has, err := engine.Select("name, ext").
		ID(req.FileIds[0]).Get(&file); err != nil {
		logx.Errorf("分享多文件，查询file失败，ERR: [%v]", err)
		return errors.New("出错了")
	} else if !has {
		return errors.New("信息有误")
	}

	shareName := file.Name
	if len(req.FileIds) == 1 {
		shareType = variable.GetTypeByBruteForce(file.Ext)
	} else {
		shareName += "等..."
	}

	id := strconv.FormatInt(idgen.NextId(), 10)
	url := req.Prefix + id
	if req.Auto == 1 {
		url += "?pwd=" + req.Pwd
	}

	_, err := engine.DoTransaction(func(session *xorm.Session) (interface{}, error) {
		var shareFile []*model.ShareFile
		for _, fileId := range req.FileIds {
			shareFile = append(shareFile, &model.ShareFile{
				ShareId: id,
				FileId:  fileId,
			})
		}
		if _, err := session.Insert(shareFile); err != nil {
			logx.Errorf("分享多文件，插入shareFile失败，ERR: [%v]", err)
			return nil, err
		}

		created := time.Now().Local()
		expired := created.Unix() + variable.ShareExpireType[req.ExpireType]
		share := &model.Share{}
		share.Id = id
		share.Pwd = req.Pwd
		share.Name = shareName
		share.UserId = userId
		share.Created = created
		share.Expired = expired
		share.Type = shareType
		share.Url = url
		if req.ExpireType == constant.ShareExpireForever {
			share.Status = constant.StatusShareForever
		}
		if _, err := session.Insert(share); err != nil {
			logx.Errorf("分享多文件，插入share失败，ERR: [%v]", err)
			return nil, err
		}
		return nil, nil
	})

	return err
}
