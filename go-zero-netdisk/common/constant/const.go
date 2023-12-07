package constant

const (
	UserIdKey        = "userId"
	NeedShardingSize = 67108864 // 需要分片起始大小：64MB
)

// 0：待合并/未上传，1：上传成功
const (
	StatusNetdiskUnuploaded int8 = iota
	StatusNetdiskUploaded
)

// 0：大文件未上传，1：大文件待合并，2：小文件未上传，3：上传成功
const (
	StatusFsBigFileUnuploaded int8 = iota
	StatusFsBigFileNeedMerge
	StatusFsFileUnuploaded
	StatusFsUploaded
)
