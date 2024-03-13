import http from "@/utils/http/http.js";

const getShareList = (data) => {
    return http.post("/admin/share-list", data)
}

function getUrl(id) {
    return http.get(`/admin/file-url/${id}`)
}

function getShareFilesByShareId(id) {
    return http.get(`/admin/share-file/${id}`)
}

function setStatus(obj) {
    return http.post('/admin/share-status', {
        'id': obj.id,
        'status': obj.status,
        'type': obj.type
    })
}

export default {
    getShareList, getUrl, getShareFilesByShareId,
    setStatus
}