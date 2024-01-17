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
					Path:    "/check_file",
					Handler: upload.CheckFileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/",
					Handler: upload.UploadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/check_chunk",
					Handler: upload.CheckChunkHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/chunk",
					Handler: upload.UploadChunkHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/file/upload"),
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
		rest.WithPrefix("/file/download"),
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
					Path:    "/detail/:userId",
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
					Method:  http.MethodDelete,
					Path:    "/",
					Handler: file.DeleteFilesHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/delete",
					Handler: file.DeleteFilesTrulyHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/folder",
					Handler: file.DeleteFoldersHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/folder-delete",
					Handler: file.DeleteFoldersTrulyHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/recover",
					Handler: file.RecoverFilesHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/folder-recover",
					Handler: file.RecoverFoldersHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/folder",
					Handler: file.CreateFolderHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/delete",
					Handler: file.ListDeletedItemsHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/file"),
	)
}
