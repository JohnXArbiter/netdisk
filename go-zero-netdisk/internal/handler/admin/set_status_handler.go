package admin

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lc/netdisk/internal/logic/admin"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
)

func SetStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetStatusReq
		if err := httpx.ParsePath(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := admin.NewSetStatusLogic(r.Context(), svcCtx)
		if err := l.SetStatus(&req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
