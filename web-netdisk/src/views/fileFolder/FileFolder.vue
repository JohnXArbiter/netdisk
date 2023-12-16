<template>
    <el-upload
            ref="upload"
            class="upload-demo"
            action="actionUrl"
            :limit="20"
            :on-exceed="handleExceed"
            :auto-upload="false"
    >
        <template #trigger>
            <el-button type="primary">select file</el-button>
        </template>
        <el-button class="ml-3" type="success" @click="submitUpload">
            upload to server
        </el-button>
    </el-upload>
    <el-table
            :data="folderList"
            style="width: 100%"
    >
        <el-table-column type="selection" width="55"/>
        <el-table-column property="name" label="文件夹名" width="120">
        </el-table-column>
        <el-table-column property="updated" label="修改时间" width="120"/>
        <el-table-column property="size" label="大小" show-overflow-tooltip/>
    </el-table>
    <el-table
            :data="fileList"
            style="width: 100%"
    >
        <el-table-column type="selection" width="55"/>
        <el-table-column property="name" label="文件名" width="120">
        </el-table-column>
        <el-table-column property="updated" label="修改时间" width="120"/>
        <el-table-column property="size" label="大小" show-overflow-tooltip/>
    </el-table>
</template>

<script lang="ts" setup>
import {onMounted, reactive, ref} from 'vue'
import {ElMessage, ElTable} from 'element-plus'
import {listFolderItemsResp} from "./fileFolder.ts";
import {getFolderItems} from "./fileFolder.ts";

const actionUrl = `/upload`

let folderList = reactive([])
let fileList = reactive([])
const parentFolderId = 0

const listFolderItems = async () => {
    let res: listFolderItemsResp
    res = await getFolderItems("/file/folder/" + parentFolderId)
    if (res.code === 0 && res.data) {
        folderList = res.data.folders
        fileList = res.data.files
    } else {
        ElMessage({
            type: 'error',
            message: res.data.msg,
        })
    }
}
const tableData = [
    {
        updated: '2016-05-03',
        name: 'Tom',
        size: '38MB',
    },
    {
        updated: '2016-05-02',
        name: 'Tom',
        size: 'No. 189, Grove St, Los Angeles',
    },
    {
        updated: '2016-05-04',
        name: 'Tom',
        size: 'No. 189, Grove St, Los Angeles',
    },
    {
        updated: '2016-05-01',
        name: 'Tom',
        size: 'No. 189, Grove St, Los Angeles',
    },
]
fileList = tableData

onMounted(() => {
    listFolderItems()
})


</script>


<style scoped>

</style>