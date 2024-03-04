package redis

import "time"

const (
	UserLogin    = "user:login:"
	RegisterCode = "user:code:"

	UploadCheckKey        = "upload:file:"
	UploadCheckBigFileKey = "upload:file-shard:"
	UploadCheckChunkKeyF  = "upload:file-shard:%d:chunk:%d"

	FileFolderDownloadUrlKey = "download:folder:%d:%d" // download:folder:{userId}:{folderId}
	FileTypeDownloadUrlKey   = "download:type:%d:%d"   // download:type:{userId}:{type}
)

const (
	RegisterCodeExpire = 5*time.Minute + 10*time.Minute

	UploadCheckExpire      = 24 * time.Hour
	UploadCheckChunkExpire = 10 * time.Minute
)
