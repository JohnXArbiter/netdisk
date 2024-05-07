package file

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
	"lc/netdisk/common/constant"
	"lc/netdisk/internal/logic/mqs"
	"lc/netdisk/model"

	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (interface{}, error) {
	var (
		userId   = l.ctx.Value(constant.UserIdKey).(int64)
		engine   = l.svcCtx.Xorm
		minioSvc = l.svcCtx.MinioSvc
		es       = l.svcCtx.Es
		err      error
	)

	defer mqs.LogSend(l.ctx, err, "Search", req.Phrase)

	var files []*model.File
	bq := elastic.NewBoolQuery().Must(
		elastic.NewTermQuery("UserId", userId),
		elastic.NewBoolQuery().Should(
			elastic.NewMatchQuery("Name", req.Phrase),
			elastic.NewMatchQuery("Ext", req.Phrase),
		))

	do, err := es.Search().
		Index("netdisk_file").
		Query(bq).
		Do(context.TODO())
	if err != nil {
		err = errors.New("Search，搜索ES失败，ERR: " + err.Error())
	}

	var list []map[string]interface{}
	if err == nil && len(do.Hits.Hits) > 0 {
		for _, hit := range do.Hits.Hits {
			m := map[string]interface{}{}
			var file model.File
			if err = json.Unmarshal(hit.Source, &file); err != nil {
				logx.Errorf("Search，file: [%v] 反序列化失败，ERR: [%v]", file.Id, err)
				continue
			}
			url, err2 := minioSvc.GenUrl(file.ObjectName, file.Name, true)
			if err2 == nil {
				m["url"] = url
			}
			m["name"] = file.Name
			m["folderId"] = file.FolderId
			m["type"] = file.Type
			list = append(list, m)
		}
		return list, nil
	}

	if err = engine.Where("user_id = ?", userId).
		And("name like ?", "%"+req.Phrase+"%").
		Find(&files); err != nil {
		err = errors.New("Search，查询数据库失败，ERR: " + err.Error())
	}

	for _, file := range files {
		m := map[string]interface{}{}

		url, err2 := minioSvc.GenUrl(file.ObjectName, file.Name, true)
		if err2 == nil {
			m["url"] = url
		}
		m["name"] = file.Name
		m["folderId"] = file.FolderId
		m["type"] = file.Type
		list = append(list, m)
	}

	return list, nil
}
