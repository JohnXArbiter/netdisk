// Code generated by goctl. DO NOT EDIT.
package types

type IdPathReq struct {
	Id int64 `path:"id"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"passwordRepeat"`
	Email          string `json:"email"`
	Code           string `json:"code"`
}

type SendCodeReq struct {
	Email string `json:"email"`
}

type LoginResp struct {
	Token    string   `json:"token"`
	UserInfo UserInfo `json:"userInfo"`
}

type UserInfo struct {
	UserId    int64  `json:"userId"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Signature string `json:"signature"`
	Status    int8   `json:"status"`
}

type CheckFileReq struct {
	FolderId int64  `json:"folderId"`
	Hash     string `json:"hash"`
	Size     int64  `json:"size"`
	Name     string `json:"name"`
	Ext      string `json:"ext"`
}

type UploadReq struct {
	FileId int64 `form:"fileId"`
}

type CheckChunkReq struct {
	FileId   int64  `json:"fileId"`
	Hash     string `json:"hash"`
	ChunkSeq int64  `json:"chunkSeq"`
}

type UploadChunkReq struct {
	FileId   int64 `form:"fileId"`
	ChunkSeq int64 `form:"chunkSeq"`
}

type CheckFileResp struct {
	FileId       int64 `json:"fileId"`
	Status       int8  `json:"status"`       // 0：文件未上传，1：文件已存在
	ConfirmShard int8  `json:"confirmShard"` // 0：不分片，1：分片
}

type CheckChunkResp struct {
	Status int8 `json:"status"`
}

type IdsReq struct {
	Ids []int64 `json:"ids"`
}

type UpdateNameReq struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type CopyFilesReq struct {
	FolderId int64 `path:"folderId"`
	FileIdsStruct
}

type DeleteFilesReq struct {
	FileIdsStruct
	FolderId int64 `json:"folderId"`
}

type RecoverReq struct {
	FolderIdsStruct
	FileIdsStruct
}

type DeleteBatchTrulyReq struct {
	FolderIdsStruct
	FileIdsStruct
}

type FileTypeReq struct {
	FileType int64 `path:"fileType"`
}

type ParentFolderIdReq struct {
	ParentFolderId int64 `path:"parentFolderId"`
}

type CreateFolderReq struct {
	Name           string `json:"name"`
	ParentFolderId int64  `json:"parentFolderId"`
}

type MoveFilesReq struct {
	FolderId int64 `json:"folderId"`
	FileIdsStruct
}

type MoveFoldersReq struct {
	ParentFolderId int64 `json:"parentFolderId"`
	FolderIdsStruct
}

type ListFolderMovableFolderReq struct {
	ParentFolderId    int64   `json:"parentFolderId"`
	SelectedFolderIds []int64 `json:"selectedFolderIds"`
}

type RecoverFilesReq struct {
	Files []*RecoverFilesStruct `json:"files"`
}

type ShareReq struct {
	FileId int64  `json:"fileId"`
	Url    string `json:"url"`
}

type DeletedFilesResp struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	Status     int8   `json:"status"`
	Size       int64  `json:"size"`
	FolderId   int64  `json:"folderId,omitempty"`
	FolderName string `json:"folderName"`
	DelTime    int64  `json:"delTime"`
}

type FileResp struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	Type     int8   `json:"type"`
	Ext      string `json:"ext"`
	Status   int8   `json:"status"`
	Size     int64  `json:"size"`
	FolderId int64  `json:"folderId,omitempty"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

type ListFolderStruct struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Updated string `json:"updated,omitempty"`
}

type FolderIdsStruct struct {
	FolderIds []int64 `json:"folderIds"`
}

type FileIdsStruct struct {
	FileIds []int64 `json:"fileIds"`
}

type RecoverFilesStruct struct {
	FileId   int64 `json:"fileId"`
	FolderId int64 `json:"folderId"`
}

type ShareFileStruct struct {
	Id     int64  `json:"id"`
	FileId int64  `json:"fileId"`
	Name   string `json:"name"`
}

type CheckSizeReq struct {
	FileId int64 `json:"fileId"`
}

type DownloadReq struct {
	FileId int64 `json:"fileId"`
}

type ChunkDownloadReq struct {
	FileId   int64 `json:"fileId"`
	ChunkSeq int64 `json:"chunkSeq"`
}

type CheckSizeResp struct {
	IsBig    int8  `json:"status"`
	ChunkNum int64 `json:"chunkNum"`
}

type UpdateUserDetailReq struct {
	Name      string `json:"name"`
	Signature string `json:"signature"`
	Email     string `json:"email"`
	Code      string `json:"code"`
}
