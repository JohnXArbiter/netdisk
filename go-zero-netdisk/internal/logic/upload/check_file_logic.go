package upload

import (
	"context"
	"github.com/yitter/idgenerator-go/idgen"
	"lc/netdisk/common/xorm"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"lc/netdisk/model"

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

const fileRepositoryKey = "fileRepository"

func (l *CheckFileLogic) CheckFile(req *types.CheckFileReq) (*types.CheckFileResp, error) {
	var (
		userId int64 = 123
		ext          = req.Ext
		engine       = l.svcCtx.Xorm
		resp   *types.CheckFileResp
	)

	fileRepository := &model.FileRepository{
		Hash: req.Hash,
		Ext:  ext,
	}
	has, err := engine.Get(fileRepository)
	if err != nil {
		// TODO: log
		return nil, err
	}

	// 文件不存在时
	if !has {
		fileRepository.Size = req.Size
		ctx := context.WithValue(l.ctx, fileRepositoryKey, fileRepository)
		_, err := xorm.Transaction(ctx, engine, l.createFile)

		return nil, err
	}

	// 文件存在时
	if has {
		// 判断该用户是否上传过
		fileNetdisk := model.FileNetdisk{
			UserId:       userId,
			RepositoryId: fileRepository.Id,
		}
		if has, err = engine.Get(fileNetdisk); err != nil {
			// TODO: log
			return nil, err
		}

		// 用户上传过，提示并返回
		if has {
			return &types.CheckFileResp{
				FileId: fileNetdisk.Id,
				Status: 1,
			}, nil
		} else {
			// 用户未上传，信息落库
			fileNetdisk.Id = idgen.NextId()
			fileNetdisk.Name = req.Name + ext
			fileNetdisk.FolderId = req.FolderId
			fileNetdisk.Status = 1
			fileNetdisk.Url = fileRepository.Url
			if _, err = engine.Insert(fileNetdisk); err != nil {
				// TODO: log
				return nil, err
			}
		}
		resp = &types.CheckFileResp{
			FileId: fileNetdisk.Id,
			Status: 1,
		}
	}

	return resp, nil
}

func (l *CheckFileLogic) createFile(ctx context.Context, session *xorm.Session) (interface{}, error) {
	var (
		fileRepository = ctx.Value(fileRepositoryKey).(*model.FileRepository)
	)

	id, err := session.Insert(fileRepository)
	return id, err
}
