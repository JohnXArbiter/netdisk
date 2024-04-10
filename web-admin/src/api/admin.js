import http from "@/utils/http/http.js";

function getAdminInfo(id) {
    return http.get(`/admin/info`)
}

function getAdminList(page) {
    return http.post('/admin/list', page)
}

function setStatus(obj) {
    return http.post('/admin/status', obj)
}

function add(form) {
    return http.post('/admin/add', form)
}

export default {
    getAdminInfo, getAdminList,
    setStatus, add
}