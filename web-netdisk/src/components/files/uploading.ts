interface CheckReq {
    folderId: number
    hash: string
    size: number
    name: string
    ext: string
}

interface CheckRes {
    success: boolean
    type: number
}
