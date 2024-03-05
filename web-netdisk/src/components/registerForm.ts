import api from "../utils/apis/request.ts";
import {Resp} from "../utils/apis/base.ts";

export interface RegisterReq {
    username: string,
    password: string,
    passwordRepeat: string
    email: string
    code: string
}

export interface RegisterResp {
}

export function registerPost(registerReq: RegisterReq) {
    return api.post<any, Resp<RegisterResp>>("/register", registerReq)
}
