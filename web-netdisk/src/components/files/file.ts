import api from "../../utils/apis/request.ts";
import {Resp} from "../../utils/apis/base.ts";

export interface File {
    id: number
    name: string
    size: number
    url: string
    status: number
    updated: string
}

export function listFilesByFolderId(parentFolderId: number) {
    return api.get<any, Resp<File[]>>(`/file/list/${parentFolderId}`)
}

export function listFilesByFileType(fileType: number) {
    return api.get<any, Resp<File[]>>(`/file/type/${fileType}`)
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
    return api.post<any, Resp<any>>('', {'id': ids})
}
