<template>
    <el-row>
        <el-col :span="24">
            <div class="file-table">
                <el-upload v-if="fileButtonsState === 0"
                           ref="uploadFiless"
                           class="upload-demo"
                           action="actionUrl"
                           multiple
                           :limit="20"
                           :auto-upload="false"
                           :on-change="change"
                           :http-request="a"
                >
                    <template #trigger>
                        <el-button type="primary" :icon="Select" round>选择文件</el-button>
                    </template>
                    <el-button type="success" :icon="Upload" round @click="">
                        确认上传
                    </el-button>
                </el-upload>

                <div class="button-group">
                    <template v-if="fileButtonsState !== 0">
                        <el-button-group>
                            <el-button type="primary" round plain :icon="Download" @click="fileButton(0)">下载
                            </el-button>
                            <template v-if="fileButtonsState === 1">
                                <el-button type="primary" round plain :icon="EditPen" @click="fileButton(1)">重命名
                                </el-button>
                            </template>
                            <el-button type="primary" round plain :icon="Rank" @click="fileButton(2)">移动</el-button>
                            <el-button type="primary" round plain :icon="CopyDocument" @click="fileButton(3)">复制
                            </el-button>
                            <el-button type="danger" round plain :icon="DeleteFilled" @click="fileButton(4)">删除
                            </el-button>
                        </el-button-group>
                    </template>
                </div>

                <el-empty v-if="!fileList.arr || fileList.arr.length==0"
                          description="文件列表为空，上传你的第一个文件吧！😺"/>

                <el-table v-if="fileList && fileList.arr.length!=0"
                          ref="fileTableRef"
                          :data="fileList.arr" style="width: 100%"
                          @selection-change="fileSelectionChange"
                >
                    <el-table-column type="selection" width="55"/>
                    <el-table-column label="文件名" width="180">
                        <template #default="scope">
                            <div class="file-folder-row" style="display: flex; align-items: center">
                                <span>{{ scope.row.name }}</span>
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
                            <template>
                                <div>{{ scope.row.size }}</div>
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
        </el-col>
    </el-row>

    <el-dialog v-model="fileDialogVisible[1]" title="输入要更改的文件名">
        <el-form :model="renamingFile">
            <el-form-item label="文件名">
                <el-input v-model="renamingFile.name"/>
            </el-form-item>
        </el-form>
        <template #footer>
      <span class="dialog-footer">
        <el-button @click="fileDialogVisible[1] = false">取消</el-button>
        <el-button type="primary" @click="renameFile(1)">
          确定
        </el-button>
      </span>
        </template>
    </el-dialog>

    <el-dialog v-model="fileCopyAndMoveDialog" title="选择文件夹">
        <el-table :data="fileMovableFolderList.arr" highlight-current-row>
            <el-table-column label="" width="180">
                <template #default="scope">
                    <div style="display: flex; align-items: center">
                        <div @click="toFolder( scope.row.id)">
                            <el-icon>
                                <FolderOpened/>
                            </el-icon>
                            <span style="margin-left: 10px">{{ scope.row.name }}</span>
                        </div>
                    </div>
                </template>
            </el-table-column>
        </el-table>
        <template #footer>
      <span class="dialog-footer">
        <el-button @click="fileCopyAndMoveDialog = false">取消</el-button>
        <el-button type="primary" @click="fileCopyAndMoveConfirm()">
          确定
        </el-button>
      </span>
        </template>
    </el-dialog>

    <el-dialog v-model="fileDialogVisible[4]" title="删除文件">
        <h3>
            <el-icon>
                <Warning/>
            </el-icon>
            确定要<span style="color: red"> 删除 {{ selectedFiles.map(file => file.name).join('，') }} </span>吗？
            你可以在回收站中找到他们。
        </h3>
        <template #footer>
      <span class="dialog-footer">
        <el-button @click="fileDialogVisible[4] = false">取消</el-button>
        <el-button type="primary" @click="deleteFilesConfirm()">
          确定
        </el-button>
      </span>
        </template>
    </el-dialog>
</template>

<script lang="ts" setup>
import {ElTable, UploadFile, UploadFiles, UploadInstance} from "element-plus";
import {
    CopyDocument,
    DeleteFilled,
    Download,
    EditPen,
    FolderOpened,
    Rank,
    Select,
    Upload, Warning
} from "@element-plus/icons-vue";
import {onMounted, reactive, ref} from "vue";
import {
    copyFiles, deleteFiles,
    listFilesByFolderId,
    listFileMovableFolders, moveFiles,
    updateFileName, listFilesByFileType
} from "./file.ts";
import type {File} from '/file.ts'
import axios, {AxiosProgressEvent} from "axios";
import type {Folder} from "./folder.ts";
import {codeOk, promptSuccess, Resp} from "../../utils/apis/base.ts";

let forFolder = false
let folderId: number
let fileType: number
const props = defineProps(['fileType', 'folderId']);


let fileButtonsState = ref(0)
const fileTableRef = ref<InstanceType<typeof ElTable>>()

let folderList = reactive<{ arr: Folder[] }>({
    arr: [{
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
})
let fileList = reactive<{ arr: File[] }>({
    arr: [
        {
            id: 4444,
            name: '43',
            size: 43,
            url: 'qwe',
            status: 2,
            updated: '2016-05-07'
        },
        {
            id: 44,
            name: 'adsasd',
            size: 423,
            url: 'qwe',
            status: 2,
            updated: '2016-05-07'
        }
    ]
})

const listFiles = async () => {
    let resp: Resp<any>
    if (forFolder) {
        resp = await listFilesByFolderId(folderId)
    } else {
        resp = await listFilesByFileType(fileType)
    }
    if (resp.code === 0 && resp.data) {
        fileList.arr = resp.data
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
//     console.log(options.files)
//     options.files
//     return XMLHttpRequest
// }

const fileDialogVisible = reactive([false, false, false, false, false])
let renamingFile = reactive<any>({})
let listFoldersCurrentFolderId = 0
let fileCopyAndMoveDialog = ref(false)
let fileCopyAndMoveFlag: number
let selectedFiles: File[]
let fileMovableFolderList = reactive<{ arr: Folder[] }>({arr: folderList.arr})

// 对话框
function fileButton(option: number) {
    selectedFiles = fileTableRef.value!.getSelectionRows()
    if (!selectedFiles) {
        return
    }
    if (option === 1) {
        Object.assign(renamingFile, selectedFiles[0])
    } else if (option === 2 || option === 3) {
        toFolder(0)
        fileCopyAndMoveDialog.value = true
        fileCopyAndMoveFlag = option
        return
    }
    fileDialogVisible[option] = true
}

async function renameFile(option: number) {
    const resp = await updateFileName(renamingFile)
    if (resp && resp.code === codeOk) {
        for (const idx in folderList.arr) {
            if (fileList.arr[idx].id == renamingFile.id) {
                fileList.arr[idx].name = renamingFile.name
                break
            }
        }
        promptSuccess()
        fileDialogVisible[option] = false
    }
}

async function toFolder(folderId: number) {
    const resp = await listFileMovableFolders(folderId)
    if (resp && resp.code === codeOk) {
        fileMovableFolderList.arr = resp.data
        listFoldersCurrentFolderId = folderId
    }
}

async function fileCopyAndMoveConfirm() {
    const fileIds = selectedFiles.map(file => file.id);
    let resp: Resp<any>
    if (fileCopyAndMoveFlag === 2) {
        resp = await moveFiles(listFoldersCurrentFolderId, fileIds)
    } else if (fileCopyAndMoveFlag === 3) {
        resp = await copyFiles(listFoldersCurrentFolderId, fileIds)
    }
    if (resp && resp.code == codeOk) {
        promptSuccess()
        fileCopyAndMoveDialog.value = false
        listFoldersCurrentFolderId = 0
    }
}

async function deleteFilesConfirm() {
    await deleteFiles(selectedFiles.map(file => file.id))
    await listFiles()
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
    if (props.folderId != undefined) {
        folderId = props.folderId
        forFolder = true
    } else {
        fileType = props.fileType
    }

    listFiles()
})

</script>

<style scoped>
.button-group {
    margin-bottom: 15px;
}
</style>