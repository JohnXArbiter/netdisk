package constant

const (
	UserIdKey   = "userId"
	UserNameKey = "userName"
)

const (
	ShardingFloor float64 = 20971520 // 需要分片起始大小：20MB
	ShardingSize  float64 = 5242880  // 分片大小： 5MB

	DefaultCapacity = 1099511627776
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
	StatusFileIllegal
	StatusFileNeedMerge
)

const (
	ConfirmNotShard = iota
	ConfirmShard
)

// 0：大文件未上传，1：大文件待合并，2：小文件未上传，3：上传成功
const (
	StatusFsFileUnuploaded int8 = iota
	StatusFsFileNeedMerge
	StatusFsUploaded
)

const (
	StatusFolderUndeleted int8 = iota
	StatusFolderDeleted
)

const (
	StatusShareNotExpired = iota
	StatusShareExpired
	StatusShareIllegal
	StatusShareNotExistOrDeleted
	StatusShareForever
)

const (
	ShareExpireDay = iota
	ShareExpireWeek
	ShareExpireMonth
	ShareExpireForever
)

const (
	StatusUserOk = iota
	StatusUserBannedByAvatar
	StatusUserBannedByUsername
	StatusUserBannedByName
	StatusUserBannedBySignature
	StatusUserBannedByShare
)

const (
	StageMerging = iota
	StageNeedMerge
	StageMergeDone
)

const (
	TimeFormat1 = "2006-01-02 15:04:05"
	TimeFormat2 = "2006-01-02/15:04:05"
)

const (
	TypeDocs = iota
	TypeImage
	TypeVideo
	TypeAudio
	TypeOther
	TypeShareMulti
)
