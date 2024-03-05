// 负责用户登入进来后，说 上午|中午|下午|晚上|半夜|凌晨|早上 好
import {RegisterResp} from "@/components/registerForm.ts";
import {promptSuccess, Resp} from "@/utils/apis/base.ts";

import api from "@/utils/apis/request.ts";

export const util = () => {
    const hours = new Date().getHours()
    if (hours < 3)
        return '深夜'
    else if (hours < 6)
        return '凌晨'
    else if (hours < 9)
        return '早上'
    else if (hours < 12)
        return '上午'
    else if (hours < 14)
        return '中午'
    else if (hours < 18)
        return '下午'
    else
        return '晚上'
}

export function formatSize(size: number) {
    const units = ['B', 'K', 'M', 'G', 'T', 'P']
    while (size > 1024 && units.length > 0) {
        size /= 1024
        units.shift()
    }
    return (units[0] === 'B' ? size : size.toFixed(2)) + units[0]
}

export async function sendCode2Email(email: string) {
    const resp = await sendCode(email)
    if (resp.code === 0) {
        promptSuccess('验证码已发送至邮件😊')
    }
}

export function sendCode(email: string) {
    return api.post<any, Resp<RegisterResp>>("/code", {
        'email': email
    })
}