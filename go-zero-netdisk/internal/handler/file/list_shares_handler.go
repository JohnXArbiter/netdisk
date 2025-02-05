package file

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"lc/netdisk/internal/logic/file"
	"lc/netdisk/internal/svc"
)

func ListSharesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := file.NewListSharesLogic(r.Context(), svcCtx)
		if resp, err := l.ListShares(); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
