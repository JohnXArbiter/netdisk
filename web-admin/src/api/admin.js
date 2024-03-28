import http from "@/utils/http/http.js";

function getAdminInfo(id) {
    return http.get(`/admin/info`)
}

function getAdminList(page) {
    return http.post('/admin/list', page)
}

export default {
    getAdminInfo, getAdminList
}