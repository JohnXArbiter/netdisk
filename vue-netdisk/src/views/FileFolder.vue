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
    <el-table-column property="name" label="文件名" width="120">
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

<script setup>
import {onMounted, reactive, ref} from 'vue'
import {ElTable} from 'element-plus'
import {getAPI} from '../utils/apis/request.js'

const actionUrl = `/upload`
const folder = {
  id: 0,
  name: "",
  updated: "",
}

const file = {
  id: 0,
  name: "",
  size: 0,
  url: "",
  status: 0,
  updated: "",
}
let folderList = reactive([])
let fileList = reactive([])
const parentFolderId = 0

const listFolderItems = async () => {
  let res = {}
  try {
    res = await getAPI({url: "/file/folder/" + parentFolderId}).catch((e) => e)
  } catch (error) {
  }
  if (res?.code === 0 && res.data) {
    folderList = res.data.folders
    fileList = res.data.files
  }
}

onMounted(() => {
  listFolderItems()
})


</script>


<style scoped>

</style>