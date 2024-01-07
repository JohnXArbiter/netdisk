<template>
    <div class="upload-button">
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
                <el-button type="primary" :icon="Select" round>选择文件</el-button>
            </template>
            <el-button type="success" :icon="Upload" round @click="">
                确认上传
            </el-button>
            <el-button type="primary" :icon="FolderAdd" plain round>新建文件夹</el-button>
        </el-upload>
    </div>
    <div class="folder-table">
        <div class="button-group">
            <template v-if="folderButtonsState !== 0">
                <el-button-group>
                    <el-button type="primary" round plain :icon="Download">下载</el-button>
                    <template v-if="folderButtonsState === 1">
                        <el-button type="primary" round plain :icon="EditPen">重命名</el-button>
                    </template>
                    <el-button type="primary" round plain :icon="Rank">移动</el-button>
                    <el-button type="primary" round plain :icon="CopyDocument">复制</el-button>
                    <el-button type="danger" round plain :icon="DeleteFilled">删除</el-button>
                </el-button-group>
            </template>
        </div>

        <el-table :data="folderList" style="width: 100%"
                  @selection-change="folderSelectionChange"
        >
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
                        <div> -</div>
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
    </div>

    <div class="file-table">
        <div class="button-group">
            <template v-if="fileButtonsState !== 0">
                <el-button-group>
                    <el-button type="primary" round plain :icon="Download">下载</el-button>
                    <template v-if="fileButtonsState === 1">
                        <el-button type="primary" round plain :icon="EditPen">重命名</el-button>
                    </template>
                    <el-button type="primary" round plain :icon="Rank">移动</el-button>
                    <el-button type="primary" round plain :icon="CopyDocument">复制</el-button>
                    <el-button type="danger" round plain :icon="DeleteFilled">删除</el-button>
                </el-button-group>
            </template>
        </div>

        <el-table :data="fileList" style="width: 100%"
                  @selection-change="fileSelectionChange"
        >
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
                        <div> -</div>
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
    </div>
</template>

<script lang="ts" setup>
import {onMounted, reactive, ref} from 'vue'
import {ElMessage, ElTable, UploadFile, UploadFiles, UploadInstance} from 'element-plus'
import type {Folder, File} from '@/views/fileFolder/fileFolder.ts'
import {getFolderItems} from '@/views/fileFolder/fileFolder.ts'
import axios, {AxiosProgressEvent} from "axios";
import {FolderOpened} from '@element-plus/icons-vue'
import {Upload, Select, FolderAdd, Download, CopyDocument, EditPen, DeleteFilled, Rank} from '@element-plus/icons-vue'

let props = defineProps(["folderId"]);
let folderId = props.folderId
let fileButtonsState = ref(0)
let folderButtonsState = ref(0)

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

const listFolderItems = async () => {
    const res = await getFolderItems(folderId)
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

function folderSelectionChange(items: Folder[]) {
    if (!items || items.length == 0) {
        folderButtonsState.value = 0
    } else if (items) {
        if (items.length === 1) {
            folderButtonsState.value = 1
        } else {
            folderButtonsState.value = 2
        }
    }
}

function fileSelectionChange(items: File[]) {
    if (!items || items.length == 0) {
        fileButtonsState.value = 0
    } else if (items) {
        if (items.length === 1) {
            fileButtonsState.value = 1
        } else {
            fileButtonsState.value = 2
        }
    }
}

onMounted(() => {
    listFolderItems()
})


</script>


<style scoped>

</style>