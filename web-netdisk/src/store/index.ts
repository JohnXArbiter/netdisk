import {defineStore} from "pinia";
import api from "../utils/apis/request.ts";
import {Resp} from "../utils/apis/base.ts";

export interface UserInfo {
    id: number
    name: string
    avatar: string
    email: string
    signature: string
    capacity: number
    status: number
}

export const useBaseStore = defineStore('base', () => {
    let token = localStorage.getItem("token") || '',
        userInfoInit = {data: {id: 0, name: '', avatar: '', email: '', signature: '', capacity: 0, status: 0}},
        user: { data: UserInfo } = userInfoInit

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

    function clearToken() {
        token = ''
        localStorage.removeItem('token')
        window.location.href = '/login'
    }

    async function getUserInfo() {
        if (user.data.id == 0) {
            const resp = await api.get<any, Resp<UserInfo>>(`/user/detail/0`)
            if (resp.code === 0) {
                user.data = resp.data
            }
        }
        return user.data
    }

    async function updateUserInfo(userInfo: UserInfo, post: boolean) {
        if (post) {
            const resp = await api.post<any, Resp<UserInfo>>(`/user/detail`, userInfo)
            if (resp.code === 0) {
                user.data = userInfo
                return
            }
        }
        user.data = userInfo
    }

    return {
        userInfoInit,
        updateToken, getToken, clearToken,
        getUserInfo, updateUserInfo
    }
})
