package file

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"lc/netdisk/internal/logic/file"
	"lc/netdisk/internal/svc"
)

func ListDeletedFilesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := file.NewListDeletedFilesLogic(r.Context(), svcCtx)
		if resp, err := l.ListDeletedFiles(); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
