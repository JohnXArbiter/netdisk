import {Resp} from "../../utils/apis/base.ts";
import api from "../../utils/apis/request.ts";

export interface Folder {
    id: number
    name: string
    updated: string
}

// folder
export function listFoldersByParentFolderId(parentFolderId: number) {
    return api.get<any, Resp<Folder[]>>(`/file/folder-list/${parentFolderId}`)
}

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
    return api.post<any, Resp<{ id: number, name: string }[]>>(`/file/folder-move`)
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
    return api.post<any, Resp<any>>('/folder', {'id': ids})
}