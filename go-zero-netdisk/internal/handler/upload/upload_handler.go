package upload

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lc/netdisk/internal/logic/upload"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		}

		fileParam := &types.FileParam{
			File:       file,
			FileHeader: header,
		}

		l := upload.NewUploadLogic(r.Context(), svcCtx)
		if resp, err := l.Upload(&req, fileParam); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
