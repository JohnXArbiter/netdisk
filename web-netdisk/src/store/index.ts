import {defineStore} from "pinia";
import api from "../utils/apis/request.ts";
import {Resp} from "../utils/apis/base.ts";

export interface UserInfo {
    id: number
    name: string
    avatar: string
    email: string
    signature: string
    status: number
}

export const useBaseStore = defineStore('base', () => {
    let token = localStorage.getItem("token") || ''
    let user: UserInfo | null = null

    function updateToken(tokenStr: string) {
        token = tokenStr
        localStorage.setItem("token", tokenStr)
    }

    function getToken() {
        if (!token || token === '') {
            token = localStorage.getItem("token") || ''
        }
        return token
    }

    async function getUserInfo() {
        if (user === null) {
            const resp = await api.get<any, Resp<UserInfo>>('/detail')
            if (resp.code === 0) {
                user = resp.data
            }
        }
        return user
    }

    function updateUserInfo(userInfo: UserInfo) {
        user = userInfo
    }

    return {
        updateToken, getToken,
        getUserInfo, updateUserInfo
    }
})
