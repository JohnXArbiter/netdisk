import api from "../../utils/apis/request.ts";
import {Resp} from "../../utils/apis/base.ts";

export interface CheckReq {
    folderId: number
    hash: string
    size: number
    name: string
    ext: string
}

export interface CheckResp {
    fileId: number
    status: number
}

export interface CheckRes extends CheckResp {
    success: boolean
}

export const sliceSize = 4194304

export function checkFile(req: CheckReq) {
    return api.post<any, Resp<CheckResp>>('/upload/check', req)
}

export function upload(formData: FormData) {
    return api.post<any, Resp<any>>('/upload', formData, {headers: {'Content-Type': 'multipart/form-data'}})
}