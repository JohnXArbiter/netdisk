package handler

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"lc/netdisk/internal/logic"
	"lc/netdisk/internal/svc"
)

func pingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPingLogic(r.Context(), svcCtx)
		xhttp.JsonBaseResponseCtx(r.Context(), w, l.Ping())
	}
}
