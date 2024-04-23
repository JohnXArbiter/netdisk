package user

import (
	"context"
	"errors"
	"lc/netdisk/common/constant"
	"lc/netdisk/common/redis"
	"lc/netdisk/common/xorm"
	"lc/netdisk/internal/logic/mqs"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"
	"strconv"
	"strings"

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

func (l *UpdateAvatarLogic) UpdateAvatar(fileParam *types.FileParam) (interface{}, error) {
	var (
		loginUserId = l.ctx.Value(constant.UserIdKey).(int64)
		userIdStr   = strconv.FormatInt(loginUserId, 10)
		engine      = l.svcCtx.Xorm
		rdb         = l.svcCtx.Redis
		minioSvc    = l.svcCtx.Minio.NewService()
		objectName  string
		err         error
	)

	defer mqs.LogSend(l.ctx, err, "UpdateAvatar", objectName)

	index := strings.LastIndex(fileParam.FileHeader.Filename, ",")
	ext := fileParam.FileHeader.Filename[index+1:]
	objectName = "/avatar/" + userIdStr + ext

	return engine.DoTransaction(func(session *xorm.Session) (interface{}, error) {
		user := &model.User{Avatar: objectName}
		if _, err = session.ID(loginUserId).
			Update(user); err != nil {
			err = errors.New("更新头像，更新数据库失败，ERR: " + err.Error())
			return nil, err
		}

		if err = minioSvc.Upload(l.ctx, objectName, fileParam.File); err != nil {
			err = errors.New("更新头像，上传头像失败，ERR: " + err.Error())
			return nil, err
		}

		key := redis.UserInfoKey + userIdStr
		if err = rdb.Del(l.ctx, key).Err(); err != nil {
			err = errors.New("更新头像，上传头像失败了，ERR: " + err.Error())
			return nil, err
		}
		return nil, nil
	})
}
