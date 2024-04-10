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
		userId = l.ctx.Value(constant.UserIdKey).(int64)
		engine = l.svcCtx.Xorm
		es     = l.svcCtx.Es
		err    error
	)

	defer mqs.LogSend(l.ctx, err, "Search", req.Phrase)

	bq := elastic.NewBoolQuery()

	bq1 := elastic.NewBoolQuery().Should(
		elastic.NewMatchQuery("Name", req.Phrase),
		elastic.NewMatchQuery("Ext", req.Phrase),
	)

	bq = bq.Must(
		elastic.NewTermQuery("UserId", userId),
		bq1)

	do, err := es.Search().
		Index("netdisk_file").
		Query(bq).
		Do(context.TODO())

	var files []*model.File
	if len(do.Hits.Hits) > 0 {
		for _, hit := range do.Hits.Hits {
			var file model.File
			err = json.Unmarshal(hit.Source, &file)
			if err != nil {
				logx.Errorf("Search，file: [%v] 反序列化失败，ERR: [%v]", file.Id, err)
				continue
			}
			files = append(files, &file)
		}
	} else {
		if err = engine.Where("user_id = ?", userId).
			And("name like ?", "%"+req.Phrase+"%").
			Find(&files); err != nil {
			err = errors.New("Search，查询数据库失败，ERR: " + err.Error())
			logx.Errorf(err.Error())
		}
	}

	return files, nil
}
