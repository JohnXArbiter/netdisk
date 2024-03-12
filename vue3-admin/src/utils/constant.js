export const
    TypeDocs = 0,
    typeImage = 1,
    typeVideo = 2,
    typeAudio = 3,
    typeOther = 4,
    typeMulti = 5

export const typeMap = {
    0: '文档',
    1: '图片',
    2: '视频',
    3: '音频',
    4: '其他',
    5: '多文件',
}

export const
    shareNotExpired = 0,
    shareExpired = 1,
    shareIllegal = 2,
    shareNotExistOrDeleted = 3,

    userOk = 0,
    userBannedByAvatar = 1,
    userBannedByUsername = 2,
    userBannedByName = 3,
    userBannedBySignature = 4,
    userBannedByShare = 5

export const userMap = {
    0: '正常',
    1: '头像违规，暂时封禁',
    2: '帐号包含违规信息，暂时封禁',
    3: '昵称包含违规信息，暂时封禁',
    4: '签名包含违规信息，暂时封禁',
    5: '分享内容违规，暂时封禁',
}