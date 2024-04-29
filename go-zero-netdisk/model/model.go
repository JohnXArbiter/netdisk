package model

import "time"

type Model struct {
	Id      int64     `xorm:"pk autoincr" json:"id"`
	Created time.Time `xorm:"created" json:"created"`
	Updated time.Time `xorm:"updated" json:"updated"`
}

// User 用户
type (
	User struct {
		Model     `xorm:"extends"`
		Username  string `xorm:"varchar(20) notnull unique 'username' comment('账号')" redis:"capacity" json:"username"`
		Password  string `xorm:"varchar(255) notnull default '' 'password'" json:"password"`
		Name      string `xorm:"varchar(20) notnull default '' 'name'" json:"name"`
		Avatar    string `xorm:"varchar(255) notnull default '' 'avatar'" json:"avatar"`
		Email     string `xorm:"varchar(127) notnull default '' 'email'" json:"email"`
		Signature string `xorm:"varchar(255) notnull default '' 'signature'" json:"signature"`
		Status    int8   `xorm:"tinyint notnull default 0 'status'" json:"status"`
		Used      int64  `xorm:"bigint notnull default 0 'used' comment('已用容量')" json:"used"`
		Capacity  int64  `xorm:"bigint notnull default 0 'capacity' comment('空间容量')" json:"capacity"`
	}

	Admin struct {
		Model    `xorm:"extends"`
		Name     string `xorm:"varchar(20) notnull default '' 'name'" json:"name"`
		Username string `xorm:"varchar(20) notnull unique 'username' comment('账号')" json:"username"`
		Password string `xorm:"varchar(255) notnull default '' 'password'" json:"password"`
		Status   int8   `xorm:"tinyint notnull default 0 'status'" json:"status"`
	}
)

type (
	// File 用户视角存储
	File struct {
		Model      `xorm:"extends"`
		UserId     int64     `xorm:"bigint notnull default 0 'user_id'" json:"userId"`
		FsId       int64     `xorm:"bigint notnull default 0 'fs_id'" json:"fsId"`
		FolderId   int64     `xorm:"bigint notnull default 0 'folder_id'" json:"folderId"`
		Name       string    `xorm:"varchar(255) notnull default '' 'name' comment('用户视角文件名')" json:"name"`
		ObjectName string    `xorm:"varchar(255) notnull default '' 'object_name' comment('存储路径名')" json:"objectName"`
		Ext        string    `xorm:"varchar(255) notnull default '' 'ext' comment('文件扩展名')" json:"ext"`
		Size       int64     `xorm:"bigint notnull default 0 'size' comment('文件大小')" json:"size"`
		Type       int8      `xorm:"tinyint notnull default 0 comment('类别')" json:"type"`
		Status     int8      `xorm:"tinyint notnull default 0 'status' comment('文件状态，0：待合并/未上传，1：上传成功')" json:"status"`
		IsBig      int8      `xorm:"tinyint notnull default 0 'is_big' comment('是否大文件，0：不是，1：是')" json:"isBig"`
		DoneAt     time.Time `xorm:"datetime 'done_at' comment('大文件合并完成时间')" json:"doneAt"`
		DelFlag    int8      `xorm:"tinyint notnull default 0 'del_flag' comment('文件删除状态：2：未删除，3：删除')" json:"delFlag"`
		DelTime    int64     `xorm:"bigint notnull default 0 'del_time' comment('')" json:"delTime"`
		SyncFlag   int8      `xorm:"tinyint notnull default 0 'sync_flag' comment('同步es标志')" json:"syncFlag"`
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
		Status     int8      `xorm:"tinyint notnull default 0 'status' comment('文件状态，0：大文件未上传，1：大文件待合并，2：小文件未上传，3：上传成功')"` //
		DoneAt     time.Time `xorm:"datetime 'done_at' comment('大文件合并完成时间')"`
	}

	// FileSchedule 任务表
	FileSchedule struct {
		Model    `xorm:"extends"`
		FileId   int64 `xorm:"bigint notnull default 0 'file_id'"`
		FsId     int64 `xorm:"bigint notnull default 0 'fs_id'"`
		ChunkNum int64 `xorm:"bigint notnull default 0 'chunk_num'"`
		Stage    int8  `xorm:"tinyint notnull default 0 'stage'"`
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
		Id          string    `xorm:"pk varchar(255) notnull default '' 'id' comment('分享id')" json:"id"`
		Pwd         string    `xorm:"varchar(8) notnull default '' 'pwd' comment('分享密码')" json:"pwd"`
		Name        string    `xorm:"varchar(64) notnull default '' 'name'" json:"name"`
		UserId      int64     `xorm:"bigint notnull default 0 'user_id'" json:"userId"`
		Url         string    `xorm:"varchar(255) notnull default '' 'url'" json:"url"`
		Created     time.Time `xorm:"created" json:"created"`
		Expired     int64     `xorm:"bigint notnull default 0 'expired' comment('到期时间')" json:"expired"`
		DownloadNum int64     `xorm:"bigint notnull default 0 'download_num'" json:"downloadNum"`
		ClickNum    int64     `xorm:"bigint notnull default 0 'click_num'" json:"clickNum"`
		Status      int8      `xorm:"tinyint notnull default 0 'status'" json:"status"`
		Type        int8      `xorm:"tinyint notnull default 0 'type'" json:"type"`
		Reason      string    `xorm:"varchar(1023) notnull default '' 'reason'" json:"reason"`
	}

	ShareFile struct {
		Model   `xorm:"extends"`
		ShareId string `xorm:"varchar(255) notnull default '' 'share_id' comment('分享id')"`
		FileId  int64  `xorm:"bigint notnull default 0 'file_id'"`
	}
)

type Basic struct {
	Capacity int64 `xorm:"bigint notnull default 0 'capacity'"`
}

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

func (*Admin) TableName() string {
	return "admin"
}

func (*Basic) TableName() string {
	return "basic"
}
