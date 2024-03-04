import {Resp} from "@/utils/apis/base.ts";
import api from "@/utils/apis/request.ts";


export async function updateAvatar(formData: FormData) {
    return await api.post<any, Resp<any>>('/user/avatar', formData, {headers: {'Content-Type': 'multipart/form-data'}});
}