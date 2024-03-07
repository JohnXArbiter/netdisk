package model

import "time"

type Model struct {
	Id      int64     `xorm:"pk autoincr"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

// User 用户
type User struct {
	Model     `xorm:"extends"`
	Username  string `xorm:"varchar(20) notnull unique 'username' comment('账号')" redis:"capacity"`
	Password  string `xorm:"varchar(255) notnull default '' 'password'"`
	Name      string `xorm:"varchar(20) notnull default '' 'name'" redis:"capacity"`
	Avatar    string `xorm:"varchar(255) notnull default '' 'avatar'" redis:"capacity"`
	Email     string `xorm:"varchar(127) notnull default '' 'email'" redis:"capacity"`
	Signature string `xorm:"varchar(255) notnull default '' 'signature'" redis:"capacity"`
	Status    int8   `xorm:"tinyint notnull default 0 'status'" redis:"capacity"`
	Used      int64  `xorm:"bigint notnull default 0 'used' comment('已用容量')" redis:"capacity"`
	Capacity  int64  `xorm:"bigint notnull default 0 'capacity' comment('空间容量')" redis:"capacity"`
}

type (
	// File 用户视角存储
	File struct {
		Model      `xorm:"extends"`
		UserId     int64     `xorm:"bigint notnull default 0 'user_id'"`
		FsId       int64     `xorm:"bigint notnull default 0 'fs_id'"`
		FolderId   int64     `xorm:"bigint notnull default 0 'folder_id'"`
		Name       string    `xorm:"varchar(255) notnull default '' 'name' comment('用户视角文件名')"`
		ObjectName string    `xorm:"varchar(255) notnull default '' 'object_name' comment('存储路径名')"`
		Url        string    `xorm:"varchar(255) notnull default '' 'url' comment('')"`
		Ext        string    `xorm:"varchar(255) notnull default '' 'ext' comment('文件扩展名')"`
		Size       int64     `xorm:"bigint notnull default 0 'size' comment('文件大小')"`
		Type       int8      `xorm:"tinyint notnull default 0 comment('类别')"`
		Status     int8      `xorm:"tinyint notnull default 0 'status' comment('文件状态，0：待合并/未上传，1：上传成功')"`
		IsBig      int8      `xorm:"tinyint notnull default 0 'is_big' comment('是否大文件，0：不是，1：是')"`
		DoneAt     time.Time `xorm:"datetime 'done_at' comment('大文件合并完成时间')"`
		DelFlag    int8      `xorm:"tinyint notnull default 0 'del_flag' comment('文件删除状态：2：未删除，3：删除')"`
		DelTime    int64     `xorm:"bigint notnull default 0 'del_time' comment('')"`
	}

	// FileFs 实际存储
	FileFs struct {
		Model      `xorm:"extends"`
		Bucket     string    `xorm:"varchar(255) notnull default '' bucket comment('桶名')"`
		Ext        string    `xorm:"varchar(255) notnull default '' 'ext' comment('文件扩展名')"`
		ObjectName string    `xorm:"varchar(255) notnull default '' 'object_name' comment('存储路径名')"`
		Hash       string    `xorm:"varchar(255) notnull default '' 'hash' comment('哈希值')"`
		Name       string    `xorm:"varchar(255) notnull default '' 'name' comment('实际文件名')"`
		Size       int64     `xorm:"bigint notnull default 0 'size' comment('文件大小')"`
		ChunkNum   int64     `xorm:"bigint notnull default 0 'chunk_num' comment('分片数量')"`
		Url        string    `xorm:"varchar(255) notnull default '' 'url' comment('访问地址')"`
		Status     int8      `xorm:"tinyint notnull default 0 'status' comment('文件状态，0：大文件未上传，1：大文件待合并，2：小文件未上传，3：上传成功')"` //
		DoneAt     time.Time `xorm:"datetime 'done_at' comment('大文件合并完成时间')"`
	}

	// FileSchedule 任务表
	FileSchedule struct {
		Model    `xorm:"extends"`
		FileId   int64 `xorm:"bigint notnull default 0 'file_id'"`
		FsId     int64 `xorm:"bigint notnull default 0 'fs_id'"`
		ChunkNum int64 `xorm:"bigint notnull default 0 'chunk_num'"`
	}
)

// Folder 网盘文件夹
type (
	Folder struct {
		Model    `xorm:"extends"`
		ParentId int64  `xorm:"bigint notnull default 0 'parent_id' comment('父文件夹id')"`
		Name     string `xorm:"varchar(64) notnull default '' 'name' comment('文件夹名')"`
		UserId   int64  `xorm:"bigint notnull default 0 'user_id'"`
		DelFlag  int8   `xorm:"tinyint notnull default 0 'del_flag' comment('文件删除状态：0：未删除，1：删除')"`
		DelTime  int64  `xorm:"bigint notnull default 0 'del_time'"`
	}
)

type (
	Share struct {
		Id          string    `xorm:"varchar(255) notnull default '' 'id' comment('分享id')"`
		Pwd         string    `xorm:"varchar(8) notnull default '' 'pwd' comment('分享密码')"`
		Name        string    `xorm:"varchar(64) notnull default '' 'name'"`
		UserId      int64     `xorm:"bigint notnull default 0 'user_id'"`
		Created     time.Time `xorm:"created"`
		Expired     int64     `xorm:"bigint notnull default 0 'expired' comment('到期时间')"`
		DownloadNum int64     `xorm:"bigint notnull default 0 'download_num'"`
		ClickNum    int64     `xorm:"bigint notnull default 0 'click_num'"`
		Status      int8      `xorm:"tinyint notnull default 0 'status'"`
		Type        int8      `xorm:"tinyint notnull default 0 'type'"`
	}

	ShareFile struct {
		Model   `xorm:"extends"`
		ShareId string `xorm:"varchar(255) notnull default '' 'share_id' comment('分享id')"`
		FileId  int64  `xorm:"bigint notnull default 0 'file_id'"`
	}
)

func (*User) TableName() string {
	return "user"
}

func (*File) TableName() string {
	return "file"
}

func (*FileFs) TableName() string {
	return "file_fs"
}

func (*Folder) TableName() string {
	return "folder"
}

func (*Share) TableName() string {
	return "share"
}

func (*ShareFile) TableName() string {
	return "share_file"
}
