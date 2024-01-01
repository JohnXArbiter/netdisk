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

    const updateToken = (tokenStr: string) => {
        localStorage.setItem("token", tokenStr)
    }

    const getToken = () => {
        if (!token || token === '') {
            token = localStorage.getItem("token") || ''
        }
        return token
    }

    const getUserInfo = async () => {
        if (user === null) {
            const resp = await api.get<any, Resp<UserInfo>>('/detail')
            if (resp.code === 0) {
                user = resp.data
            }
        }
        return user
    }
    return {
        updateToken, getToken,
        getUserInfo
    }
})
