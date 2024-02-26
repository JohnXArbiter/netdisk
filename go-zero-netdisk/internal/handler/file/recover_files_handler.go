package file

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lc/netdisk/internal/logic/file"
	"lc/netdisk/internal/svc"
	"lc/netdisk/internal/types"
)

func RecoverFilesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecoverFilesReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		var fileIds, folderIds []int64
		m := make(map[int64]struct{})
		for _, f := range req.Files {
			fileIds = append(fileIds, f.FileId)
			if _, ok := m[f.FolderId]; !ok {
				folderIds = append(folderIds, f.FolderId)
				m[f.FolderId] = struct{}{}
			}
		}

		l := file.NewRecoverFilesLogic(r.Context(), svcCtx)
		if err := l.RecoverFiles(fileIds, folderIds); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
