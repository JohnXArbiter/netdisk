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

export function updateFilename(file: File) {
    return api.put<any, Resp<any>>('/file', {'id': file.id, 'name': file.name})
}

export function listFileMovableFolders(folderId: number) {
    return api.get<any, Resp<{ id: number, name: string }[]>>(`/move/${folderId}`)
}

export function deleteFiles(ids: number[]) {
    return api.put<any, Resp<any>>('', {"id": ids})
}