package constant

const (
	UserIdKey   = "userId"
	UserNameKey = "userName"
)

const (
	ShardingFloor float64 = 10485760 // 需要分片起始大小：10MB
	ShardingSize  float64 = 5242880  // 分片大小： 5MB
)

const (
	SmallFileFlag int8 = iota
	BigFileFlag
)

// 0：待合并/未上传，1：上传成功
const (
	StatusFileUnuploaded int8 = iota
	StatusFileUploaded
	StatusFileUndeleted
	StatusFileDeleted
)

const (
	ConfirmNotShard = iota
	ConfirmShard
)

// 0：大文件未上传，1：大文件待合并，2：小文件未上传，3：上传成功
const (
	StatusFsBigFileUnuploaded int8 = iota
	StatusFsBigFileNeedMerge
	StatusFsFileUnuploaded
	StatusFsUploaded
)

const (
	StatusFolderUndeleted int8 = iota
	StatusFolderDeleted
)

// context keys
const (
	CtxFolderIdsKey = "folderIds"
	CtxFileIdsKey   = "fileIds"
)

const (
	TimeFormat1 = "2006-01-02 15:04:05"
	TimeFormat2 = "2006-01-02/15:04:05"
)
