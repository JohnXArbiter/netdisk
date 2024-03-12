import http from "@/utils/http/http.js";

const getShareList = (data) => {
    return http.post("/admin/share-list", data)
}

function getUrl(id) {
    return http.get(`/admin/file-url/${id}`)
}

export default {
    getShareList, getUrl
}