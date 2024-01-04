<template>
    <el-upload
            ref="uploadFiless"
            class="upload-demo"
            action="actionUrl"
            multiple
            :limit="20"
            :auto-upload="false"
            :on-change="change"
            :http-request="uploadProcedure"
    >
        <template #trigger>
            <el-button type="primary">select file</el-button>
        </template>
        <el-button class="ml-3" type="success" @click="">
            upload to server
        </el-button>
    </el-upload>
    <el-table
            :data="folderList"
            style="width: 100%"
            empty-text="暂无文件夹"
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
            empty-text="暂无文件"
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
import {ElMessage, ElTable, UploadFile, UploadFiles, UploadInstance} from 'element-plus'
import {folder, file} from './fileFolder.ts';
import {getFolderItems} from './fileFolder.ts';
import axios, {AxiosProgressEvent} from "axios";

let folderList = reactive<folder[]>([
  {
    id: 111,
    updated: '2016-05-03',
    name: 'Jerry',
    size: '-',
  },
  {
    id: 222,
    updated: '2016-05-02',
    name: 'Tom',
    size: '-',
  },
  {
    id: 333,
    updated: '2016-05-04',
    name: 'Sam',
    size: '-',
  }
])
let fileList = reactive<file[]>([
  {
    id: 4444,
    name: '43',
    size: 43,
    url: 'qwe',
    status: 2,
    updated: '2016-05-07'
  }
])
const parentFolderId = 0

const listFolderItems = async () => {
    const res = await getFolderItems(parentFolderId)
    if (res.code === 0 && res.data) {
      folderList = res.data.folders
      folderList.forEach(folder => {
        folder.size = '-';
      })
      fileList = res.data.files

    } else {
        ElMessage({
            type: 'error',
            message: res.msg,
        })
    }
}

let uploadFiless = ref<UploadInstance[]>([])

function change(uploadFile: UploadFile, uploadFiles: UploadFiles) {
    console.log("111", uploadFiless)
    console.log("222", uploadFile)
    console.log("333", uploadFiles)

}

function asd(e: Event) {
  const target = e.target
  if (target instanceof HTMLInputElement) {
    const file = target.files
    if (file) {
      const form = new FormData()
      for (let i = 1; i < file.length; i++) {
        form.append("file", file[i])
      }
      axios.post("/", form, {
        onUploadProgress:(progressEvent: AxiosProgressEvent) => {
          Math.round((progressEvent.loaded/(progressEvent.total as number)*100))
        }
      })
    }
  }
}

// const uploadProcedure = (options: UploadRequestOptions) => {
//     console.log(options.file)
//     options.file
//     return XMLHttpRequest
// }

onMounted(() => {
    listFolderItems()
})


</script>


<style scoped>

</style>