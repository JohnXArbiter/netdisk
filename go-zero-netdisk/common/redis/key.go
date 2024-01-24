package redis

import "time"

const (
	UserLogin = "user:login:"

	UploadCheckKey        = "upload:file:"
	UploadCheckBigFileKey = "upload:file-shard:"

	DownloadGetFsKey = "download:fs:fn:"
)

const (
	UploadCheckExpire = 24 * time.Hour
)
