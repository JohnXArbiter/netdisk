import http from "@/utils/http/http.js";

function getAdminInfo(id) {
    return http.get(`/admin/info`)
}

function getAdminList(page) {
    return http.post('/admin/list', page)
}

function add(form) {
    return http.post('/admin/add', form)
}

export default {
    getAdminInfo, getAdminList, add
}