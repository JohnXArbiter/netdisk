import api from "@/utils/apis/request.ts";
import {Resp} from "@/utils/apis/base.ts";

export interface ShareItem {
    id: number
    // type: number
    name: string
    updated: string
    size: number
    url?: string
    sizeStr?: string
}

export interface ListResp {
    name: string
    created: string
    expired: number
    owner: number
    items: ShareItem[]
}

export function listFilesByShareId(shareId: string, pwd: string) {
    return api.post<any, Resp<ListResp>>(`/file/share-info`, {
        'id': shareId,
        'pwd': pwd,
    })
}