package model

import "time"

type Model struct {
	Id       int64     `xorm:"pk autoincr"`
	CreateAt time.Time `xorm:"create_at"`
	UpdateAt time.Time `xorm:"update_at"`
	DeleteAt time.Time `xorm:"delete_at"`
}

// 未命名数据模型
type User struct {
	Model     `xorm:"extends"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Signature string `json:"signature"`
}

// FileNetdisk 用户视角存储
type FileNetdisk struct {
	Model    `xorm:"extends"`
	UserId   int64     `xorm:"user_id" json:"userId"`
	FsId     int64     `xorm:"fs_id" json:"fsId"`
	FolderId int64     `xorm:"folder_id" json:"folderId"`
	Name     string    `json:"name"`     // 用户视角文件名
	Url      string    `json:"url"`      // 访问地址
	Status   int8      `json:"status"`   // 文件状态，0：待合并/未上传，1：上传成功
	IsBig    int8      `json:"is_big"`   // 是否大文件，0：不是，1：是
	DoneAt   time.Time `json:"done_at"`  // 大文件合并完成时间
	DelFlag  int8      `json:"del_flag"` // 文件删除状态：0：未删除，1：删除
}

// FileFs 实际存储
type FileFs struct {
	Model      `xorm:"extends"`
	Bucket     string    `json:"bucket"`      // 桶名
	Ext        string    `json:"ext"`         // 文件扩展名
	ObjectName string    `json:"object_name"` // 存储路径名
	Hash       string    `json:"hash"`        // 哈希值
	Name       string    `json:"name"`        // 实际文件名
	Size       int64     `json:"size"`        // 文件大小
	ChunkNum   int64     `json:"chunkNum"`    // 分片数量
	Url        string    `json:"url"`         // 访问地址
	Status     int8      `json:"status"`      // 文件状态，0：大文件未上传，1：大文件待合并，2：小文件未上传，3：上传成功
	DoneAt     time.Time `json:"done_at"`     // 大文件合并完成时间
}

// FileUploading 上传中间态
type FileUploading struct {
	Model     `xorm:"extends"`
	NetdiskId int64 `xorm:"netdisk_id" json:"netdiskId"`
	FsId      int64 `xorm:"repository_id" json:"repositoryId"`
	ChunkNum  int   `xorm:"chunk_num" json:"chunkNum"`
}

// FileFolder 网盘文件夹
type FileFolder struct {
	Model    `xorm:"extends"`
	ParentId int64  `json:"parent_id"` // 父文件夹id
	Name     string `json:"name"`      // 文件夹名
	UserId   int64  `json:"user_id"`
}

func (*FileNetdisk) TableName() string {
	return "file_netdisk"
}

func (*FileFs) TableName() string {
	return "file_fs"
}
