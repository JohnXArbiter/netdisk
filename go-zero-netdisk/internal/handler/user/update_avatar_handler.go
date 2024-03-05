package user

import (
	xhttp "github.com/zeromicro/x/http"
	"lc/netdisk/common/utils"
	"net/http"

	"lc/netdisk/internal/logic/user"
	"lc/netdisk/internal/svc"
)

func UpdateAvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fileParam, err := utils.GetSingleFile(r)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		}

		l := user.NewUpdateAvatarLogic(r.Context(), svcCtx)
		if resp, err := l.UpdateAvatar(fileParam); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
