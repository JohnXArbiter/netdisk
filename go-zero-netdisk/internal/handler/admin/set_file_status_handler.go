package admin

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lc/netdisk/internal/logic/admin"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
)

func SetFileStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetFileStatusReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := admin.NewSetFileStatusLogic(r.Context(), svcCtx)
		if err := l.SetFileStatus(&req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
