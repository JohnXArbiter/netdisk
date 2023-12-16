import baseResp from "../../utils/apis/base.ts";
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

export interface listFolderItemsResp extends baseResp {
    data: {
        folders: folder[]
        files: file[]
    }
}

export const getFolderItems = (url: string) => {
    return api.get<any, listFolderItemsResp>(url)
}

