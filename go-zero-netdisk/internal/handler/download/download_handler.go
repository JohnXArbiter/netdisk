package download

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lc/netdisk/internal/logic/download"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
)

func DownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := download.NewDownloadLogic(r.Context(), svcCtx)

		if filename, err := l.Download(&req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			http.ServeFile(w, r, filename)
		}
	}
}
