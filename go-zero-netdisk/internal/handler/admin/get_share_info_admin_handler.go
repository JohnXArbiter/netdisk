package admin

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lc/netdisk/internal/logic/admin"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
)

func GetShareInfoAdminHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdStrReq
		if err := httpx.ParsePath(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := admin.NewGetShareInfoAdminLogic(r.Context(), svcCtx)
		if resp, err := l.GetShareInfoAdmin(&req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
