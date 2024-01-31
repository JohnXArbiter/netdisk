package handler

import (
	"fmt"
	"net/http"

	"lc/netdisk/internal/svc"
)

func testHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.Header)

		w.Header().Set("Content-Disposition", `attachment; filename="example.txt"`)
		w.Header().Set("Content-Type", "application/octet-stream")

		http.ServeFile(w, r, "D:\\ubuntu-20.04.4-live-server-amd64.iso")
	}
}
