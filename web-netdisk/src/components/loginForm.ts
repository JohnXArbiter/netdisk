import baseResp from "../utils/apis/base.ts";

export interface loginForm {
    username: string,
    password: string
}

export interface loginResp extends baseResp {
    data: {
        token: string
    }
}