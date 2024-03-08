import http from "@/utils/http/http.js";

const getShareList = (data) => {
    return http.post("/admin/share-list", data);
};
export default {
    getShareList
}