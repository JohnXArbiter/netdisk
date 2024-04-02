package logic

import (
	"context"
	"errors"
	"lc/netdisk/common/constant"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByShareIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoByShareIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByShareIdLogic {
	return &GetUserInfoByShareIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoByShareIdLogic) GetUserInfoByShareId(req *types.IdPathReq) (map[string]interface{}, error) {
	var (
		engine   = l.svcCtx.Xorm
		minioSvc = l.svcCtx.MinioSvc
		share    model.Share
		user     model.User
	)

	if has, err := engine.ID(req.Id).
		Get(&share); err != nil {
		return nil, err
	} else if !has {
		return map[string]interface{}{
			"shareStatus": constant.StatusShareNotExistOrDeleted,
		}, nil
	}

	if has, err := engine.ID(share.UserId).
		Get(&user); err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("找不到用户信息！")
	}

	url, err := minioSvc.GenUrl(user.Avatar, "", false)
	if err != nil {
		logx.Errorf("GetUserInfoByShareId，获取头像url失败，ERR: [%v]", err)
	}
	resp := map[string]interface{}{
		"userId":      user.Id,
		"name":        user.Name,
		"avatar":      url,
		"signature":   user.Signature,
		"userStatus":  user.Status,
		"shareStatus": share.Status,
	}
	return resp, nil
}
