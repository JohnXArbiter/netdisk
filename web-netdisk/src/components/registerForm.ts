import api from "../utils/apis/request.ts";
import {Resp} from "../utils/apis/base.ts";

export interface RegisterReq {
    username: string,
    password: string,
    passwordRepeat: string
}

export interface RegisterResp {
}

export const registerPost = (registerReq: RegisterReq): Resp<RegisterResp> => {
    return api.post<any, Resp<RegisterResp>>("/register", registerReq)
}