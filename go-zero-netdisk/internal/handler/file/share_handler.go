package file

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lc/netdisk/internal/logic/file"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
)

func ShareHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShareReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewShareLogic(r.Context(), svcCtx)
		err := l.Share(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
