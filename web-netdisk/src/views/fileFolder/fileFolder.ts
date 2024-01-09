import {Resp} from "../../utils/apis/base.ts";
import api from "../../utils/apis/request.ts";

export interface Folder {
    id: number
    name: string
    updated: string
}

export interface File {
    id: number
    name: string
    size: number
    url: string
    status: number
    updated: string
}

export interface listFolderItemsResp {
    folders: Folder[]
    files: File[]
}

export function getFolderItems(parentFolderId: number) {
    return api.get<any, Resp<listFolderItemsResp>>('/file/folder/' + parentFolderId)
}

export function updateFileName(file: File) {
    return api.put<any, Resp<any>>('/file', {
        'id': file.id,
        'name': file.name
    })
}

export function listFileMovableFolders(folderId: number) {
    return api.get<any, Resp<{ id: number, name: string }[]>>(`/move/${folderId}`)
}

export function moveFiles(parentFolderId: number, fileIds: number[]) {
    return api.put<any, Resp<any>>('/move', {
        'parentFolderId': parentFolderId,
        'fileIds': fileIds
    })
}

export function copyFiles(parentFolderId: number, fileIds: number[]) {
    return api.post<any, Resp<any>>('/copy', {
        'parentFolderId': parentFolderId,
        'fileIds': fileIds
    })
}

export function deleteFiles(ids: number[]) {
    return api.delete<any, Resp<any>>('', {'id': ids})
}

// folder
export function createFolder(folder: Folder) {
    return api.post<any, Resp<any>>('/folder', {
        'id': folder.id,
        'name': folder.name
    })
}

export function updateFolderName(folder: Folder) {
    return api.put<any, Resp<any>>('/folder', {
        'id': folder.id,
        'name': folder.name
    })
}

export function listFolderMovableFolders(folderId: number) {
    return api.get<any, Resp<{ id: number, name: string }[]>>(`/folder-move/${folderId}`)
}

export function moveFolders(parentFolderId: number, folderIds: number[]) {
    return api.put<any, Resp<any>>('/folder-move', {
        'parentFolderId': parentFolderId,
        'folderIds': folderIds
    })
}

export function copyFolders(parentFolderId: number, folderIds: number[]) {
    return api.post<any, Resp<any>>('/folder-copy', {
        'parentFolderId': parentFolderId,
        'folderIds': folderIds
    })
}

export function deleteFolders(ids: number[]) {
    return api.delete<any, Resp<any>>('/folder', {'id': ids})
}