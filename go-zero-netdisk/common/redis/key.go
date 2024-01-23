package redis

import "time"

const (
	UserLogin = "user:login:"

	UploadCheckKey   = "upload:file:"
	DownloadGetFsKey = "download:fs:fn:"
)

const (
	UploadCheckExpire = 10 * time.Minute
)
