package admin

import (
	"context"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUsersLogic {
	return &ListUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUsersLogic) ListUsers(req *types.PageReq) ([]*model.User, error) {
	var (
		engine   = l.svcCtx.Xorm
		minioSvc = l.svcCtx.Minio.NewService()
		users    []*model.User
	)

	offset := int((req.Page - 1) * req.Size)
	if err := engine.Limit(int(req.Size), offset).Find(&users); err != nil {
		logx.Errorf("获取用户列表，查询users失败，ERR: [%v]", err)
		return nil, err
	}

	//var resp []*types.UserInfo
	for _, user := range users {
		url, err := minioSvc.GenUrl(user.Avatar, "", false)
		if err != nil {
			continue
		}
		user.Avatar = url
		//resp = append(resp, &types.UserInfo{
		//	UserId:    user.Id,
		//	Name:      user.Name,
		//	Avatar:    user.Avatar,
		//	Email:     user.Email,
		//	Signature: user.Signature,
		//	Status:    user.Status,
		//})
	}

	return users, nil
}
