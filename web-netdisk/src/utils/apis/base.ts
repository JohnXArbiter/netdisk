import {ElMessage} from "element-plus";

 interface Resp<T> {
    code: number
    msg: string
    data: T
}

const codeOk = 0
const codeError = -1
const msgOk = '操作成功！😻'

function promptSuccess() {
    ElMessage({
        type: "success",
        message: msgOk,
    })
}

export type {
    Resp
}

export {
    codeOk, codeError, msgOk,
    promptSuccess
}