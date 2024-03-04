package user

import (
	"context"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAvatarLogic {
	return &UpdateAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAvatarLogic) UpdateAvatar(req *types.UpdateAvatarReq, fileParam *types.FileParam) (interface{}, error) {
	//var (
	//	loginUserId = l.ctx.Value(constant.UserIdKey).(int64)
	//	engine      = l.svcCtx.Xorm
	//	minioSvc    = l.svcCtx.Minio.NewService()
	//)
	//
	//index := strings.LastIndex(fileParam.FileHeader.Filename, ",")
	//ext := fileParam.FileHeader.Filename[index+1:]
	//objectName := "/avatar/" + strconv.FormatInt(loginUserId, 10) + ext
	//
	//_, err := engine.DoTransaction(func(session *xorm.Session) (interface{}, error) {
	//	if _, err := session.Where("id = ?", loginUserId).
	//		Update(&model.User{Avatar: objectName}); err != nil {
	//		return nil, err
	//	}
	//
	//	if err := minioSvc.Upload(l.ctx, objectName, fileParam.File); err != nil {
	//		return nil, err
	//	}
	//
	//})

	return nil, nil
}
