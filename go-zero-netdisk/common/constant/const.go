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

	BanStr = "您因%v违规而被封禁，暂时无法登录"
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

const (
	StatusAdminNormal = iota
	StatusAdminSuper
	StatusAdminBanned
)

const (
	Operation = "用户，ID：[%v]，昵称：[%v]，%v"
)

const (
	FlagSyncDone int8 = iota
	FlagSyncWrite
	FlagSyncDelete
)
