package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
	"lc/netdisk/internal/logic"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
	"net/http"
)

func getShareInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetShareInfoReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		//var userId int64
		//token := r.Header.Get("Authorization")
		//if token != "" {
		//	if redis.Redis == nil {
		//		return
		//	}
		//
		//	claim, err := utils.ParseToken(token)
		//	if err != nil {
		//		httpx.WriteJson(w, http.StatusUnauthorized, "身份认证错误或过期，请重新登录!")
		//		return
		//	}
		//
		//	userId = claim.Id
		//	key := redis.UserLogin + strconv.FormatInt(userId, 10)
		//
		//	redisToken, err := redis.Redis.Get(r.Context(), key).Result()
		//	if redisToken != token {
		//		httpx.WriteJson(w, http.StatusUnauthorized, "身份认证过期，请重新登录!")
		//		return
		//	}
		//}

		l := logic.NewGetShareInfoLogic(r.Context(), svcCtx)
		if resp, err := l.GetShareInfo(&req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
