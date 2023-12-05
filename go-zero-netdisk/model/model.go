package model

import "time"

type Model struct {
	Id       int64     `xorm:"pk autoincr"`
	CreateAt time.Time `xorm:"create_at"`
	UpdateAt time.Time `xorm:"update_at"`
	DeleteAt time.Time `xorm:"delete_at"`
}

// FileNetdisk 用户视角存储
type FileNetdisk struct {
	Model        `xorm:"extends"`
	UserId       int64  `xorm:"user_id" json:"userId"`
	RepositoryId int64  `xorm:"repository_id" json:"repositoryId"`
	FolderId     int64  `xorm:"folder_id" json:"folderId"`
	Name         string `json:"name"`     // 用户视角文件名
	Url          string `json:"url"`      // 访问地址
	Status       int64  `json:"status"`   // 文件状态，1：上传成功（小文件为1），0：待合并（大文件）
	DoneAt       string `json:"done_at"`  // 大文件合并完成时间
	DelFlag      int64  `json:"del_flag"` // 文件删除状态：0：未删除，1：删除
}

// FileRepository 实际存储
type FileRepository struct {
	Model    `xorm:"extends"`
	Bucket   string    `json:"bucket"`    // 桶名
	CreateAt string    `json:"create_at"` // 创建时间
	Ext      string    `json:"ext"`       // 文件扩展名
	Filename string    `json:"filename"`  // 存储路径名
	Hash     string    `json:"hash"`      // 哈希值
	Name     string    `json:"name"`      // 实际文件名
	Size     int64     `json:"size"`      // 文件大小
	Url      string    `json:"url"`       // 访问地址
	Status   int64     `json:"status"`    // 文件状态，1：上传成功，0：待合并
	DoneAt   time.Time `json:"done_at"`   // 大文件合并完成时间
}

// FileUploading 上传中间态
type FileUploading struct {
	Model        `xorm:"extends"`
	NetdiskId    int64 `xorm:"netdisk_id" json:"netdiskId"`
	RepositoryId int64 `xorm:"repository_id" json:"repositoryId"`
	ChunkNum     int   `xorm:"chunk_num" json:"chunkNum"`
}

func (*FileNetdisk) TableName() string {
	return "file_netdisk"
}

func (*FileRepository) TableName() string {
	return "file_repository"
}
