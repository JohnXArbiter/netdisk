// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	download "lc/netdisk/internal/handler/download"
	file "lc/netdisk/internal/handler/file"
	upload "lc/netdisk/internal/handler/upload"
	user "lc/netdisk/internal/handler/user"
	"lc/netdisk/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: loginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: registerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/code",
				Handler: sendCodeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: pingHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/check",
					Handler: upload.CheckFileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/",
					Handler: upload.UploadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/chunk-check",
					Handler: upload.CheckChunkHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/chunk",
					Handler: upload.UploadChunkHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/upload"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/check_size",
					Handler: download.CheckSizeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/",
					Handler: download.DownloadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/chunk",
					Handler: download.ChunkDownloadHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/download"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/detail",
					Handler: user.UpdateDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/detail/:id",
					Handler: user.GetDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/avatar",
					Handler: user.UpdateAvatarHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/:id",
					Handler: file.GetFileDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/list/:parentFolderId",
					Handler: file.ListFileHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/type/:fileType",
					Handler: file.ListFileByTypeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/folder-list/:parentFolderId",
					Handler: file.ListFolderHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/move/:parentFolderId",
					Handler: file.ListFileMovableFolderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/folder-move",
					Handler: file.ListFolderMovableFolderHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/move",
					Handler: file.MoveFilesHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/folder-move",
					Handler: file.MoveFoldersHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/",
					Handler: file.UpdateFileHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/folder",
					Handler: file.UpdateFolderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/copy",
					Handler: file.CopyFilesHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/delete",
					Handler: file.DeleteFilesHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: file.DeleteFilesTrulyHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/folder-delete",
					Handler: file.DeleteFoldersHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/clear",
					Handler: file.DeleteAllFilesTrulyHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/recover",
					Handler: file.RecoverFilesHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/folder",
					Handler: file.CreateFolderHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/delete",
					Handler: file.ListDeletedFilesHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/share",
					Handler: file.ListShareFilesHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/share",
					Handler: file.ShareHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/share-folder",
					Handler: file.ShareFolderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/share-cancel",
					Handler: file.CancelShareHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/file"),
	)
}
