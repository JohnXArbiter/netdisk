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
  <el-table :data="tb" style="width: 100%">
    <el-table-column type="selection" width="55"/>
    <el-table-column label="文件夹名" width="180">
      <template #default="scope">
        <div style="display: flex; align-items: center">
          <template v-if="(scope.row as File).size != undefined">
            <span>{{ scope.row.name }}</span>
          </template>
          <template v-else>
            <el-icon>
              <FolderOpened/>
            </el-icon>
            <span style="margin-left: 10px">{{ scope.row.name }}</span>
          </template>
        </div>
      </template>
    </el-table-column>
    <el-table-column label="修改时间" width="180">
      <template #default="scope">
        <div>{{ scope.row.updated }}</div>
      </template>
    </el-table-column>
    <el-table-column label="大小" width="180">
      <template #default="scope">
        <template v-if="(scope.row as File).size != undefined ">
          <div>{{ scope.row.size }}</div>
        </template>
        <template v-else>
          <div> - </div>
        </template>
      </template>
    </el-table-column>
    <el-table-column label="Operations">
      <template #default="scope">
        <el-button size="small" @click="handleEdit(scope.$index, scope.row)"
        >Edit
        </el-button>
        <el-button
            size="small"
            type="danger"
            @click="handleDelete(scope.$index, scope.row)"
        >Delete
        </el-button>
      </template>
    </el-table-column>
  </el-table>

  <!--  <el-table :data="fileList" style="width: 100%">-->
  <!--    <el-table-column type="selection" width="55" />-->
  <!--    <el-table-column label="文件名" width="180">-->
  <!--      <template #default="scope">-->
  <!--        <div style="display: flex; align-items: center">-->
  <!--          <template v-if=""></template>-->
  <!--          <span style="margin-left: 10px">{{ scope.row.name }}</span>-->
  <!--        </div>-->
  <!--      </template>-->
  <!--    </el-table-column>-->
  <!--    <el-table-column label="修改时间" width="180">-->
  <!--      <template #default="scope">-->
  <!--        <div>{{ scope.row.updated }}</div>-->
  <!--      </template>-->
  <!--    </el-table-column>-->
  <!--    <el-table-column label="大小" width="180">-->
  <!--      <template #default="scope">-->
  <!--        <div>{{ scope.row.size }}</div>-->
  <!--      </template>-->
  <!--    </el-table-column>-->
  <!--    <el-table-column label="Operations">-->
  <!--      <template #default="scope">-->
  <!--        <el-button size="small" @click="handleEdit(scope.$index, scope.row)"-->
  <!--        >Edit-->
  <!--        </el-button>-->
  <!--        <el-button-->
  <!--            size="small"-->
  <!--            type="danger"-->
  <!--            @click="handleDelete(scope.$index, scope.row)"-->
  <!--        >Delete-->
  <!--        </el-button>-->
  <!--      </template>-->
  <!--    </el-table-column>-->
  <!--  </el-table>-->

</template>

<script lang="ts" setup>
import {onMounted, reactive, ref} from 'vue'
import {ElMessage, ElTable, UploadFile, UploadFiles, UploadInstance} from 'element-plus'
import type {Folder, File} from './fileFolder.ts';
import {getFolderItems} from './fileFolder.ts';
import axios, {AxiosProgressEvent} from "axios";
import {FolderOpened} from '@element-plus/icons-vue'

var {parentFolderId} = defineProps(["parentFolderId"]);

let folderList: Folder[] = [
  {
    id: 111,
    updated: '2016-05-03',
    name: 'Jerry',
  },
  {
    id: 222,
    updated: '2016-05-02',
    name: 'Tom',
  },
  {
    id: 333,
    updated: '2016-05-04',
    name: 'Sam',
  }
]
let fileList: File[] = [
  {
    id: 4444,
    name: '43',
    size: 43,
    url: 'qwe',
    status: 2,
    updated: '2016-05-07'
  },
  {
    id: 4444,
    name: 'adsasd',
    size: 423,
    url: 'qwe',
    status: 2,
    updated: '2016-05-07'
  }
]

const tb: any[] = reactive([])
tb.push(...folderList)
tb.push(...fileList)

const listFolderItems = async () => {
  const res = await getFolderItems(parentFolderId)
  if (res.code === 0 && res.data) {
    folderList = res.data.folders
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
        onUploadProgress: (progressEvent: AxiosProgressEvent) => {
          Math.round((progressEvent.loaded / (progressEvent.total as number) * 100))
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