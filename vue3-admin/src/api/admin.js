import http from "@/utils/http/http.js";

function getAdminInfo(id) {
    return http.get(`/admin/info`)
}

export default {
    getAdminInfo
}