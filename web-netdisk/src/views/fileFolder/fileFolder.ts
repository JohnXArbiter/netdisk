import {Resp} from "../../utils/apis/base.ts";
import api from "../../utils/apis/request.ts";

interface folder {
    id: number
    name: string
    updated: string
}

interface file {
    id: number
    name: string
    size: number
    url: string
    status: number
    updated: string
}

export interface listFolderItemsResp {
    data: {
        folders: folder[]
        files: file[]
    }
}

export const getFolderItems = (parentFolderId: number) => {
    return api.get<any, Resp<listFolderItemsResp>>("/file/folder/" + parentFolderId)
}

