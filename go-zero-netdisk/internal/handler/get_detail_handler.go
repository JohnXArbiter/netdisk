package handler

import (
	xhttp "github.com/zeromicro/x/http"
	"lc/netdisk/common/redis"
	"lc/netdisk/common/utils"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lc/netdisk/internal/logic"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
)

func getDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdPathReq
		if err := httpx.ParsePath(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		var userId int64
		token := r.Header.Get("Authorization")
		if token != "" {
			if redis.Redis == nil {
				return
			}

			claim, err := utils.ParseToken(token)
			if err != nil {
				httpx.WriteJson(w, http.StatusUnauthorized, "身份认证错误或过期，请重新登录!")
				return
			}

			userId = claim.Id
			key := redis.UserLogin + strconv.FormatInt(userId, 10)

			redisToken, err := redis.Redis.Get(r.Context(), key).Result()
			if redisToken != token {
				httpx.WriteJson(w, http.StatusUnauthorized, "身份认证过期，请重新登录!")
				return
			}
		}

		l := logic.NewGetDetailLogic(r.Context(), svcCtx)
		if resp, err := l.GetDetail(&req, userId); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
