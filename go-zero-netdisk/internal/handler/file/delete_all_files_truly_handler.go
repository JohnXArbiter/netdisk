package file

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"lc/netdisk/internal/logic/file"
	"lc/netdisk/internal/svc"
)

func DeleteAllFilesTrulyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := file.NewDeleteAllFilesTrulyLogic(r.Context(), svcCtx)
		if err := l.DeleteAllFilesTruly(); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		}
	}
}
