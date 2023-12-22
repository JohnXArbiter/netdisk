import api from "../utils/apis/request.ts";
import {Resp} from "../utils/apis/base.ts";

export interface LoginReq {
    username: string,
    password: string
}

export interface LoginResp {
    token: string
}

export const loginPost = (loginReq: LoginReq): Resp<LoginResp> => {
   return api.post<any, Resp<LoginResp>>("/login", loginReq)
}