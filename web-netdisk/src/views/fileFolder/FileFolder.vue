<template>

    <el-row>
        <el-col :span="11">
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
                </el-upload>
            </div>
            <div class="file-table">
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

                <el-table ref="fileTableRef"
                          :data="fileList" style="width: 100%"
                          @selection-change="fileSelectionChange"
                >
                    <el-table-column type="selection" width="55"/>
                    <el-table-column label="文件名" width="180">
                        <template #default="scope">
                            <div style="display: flex; align-items: center">
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
        <el-col :span="2"></el-col>
        <el-col :span="11">
            <el-upload>
                <el-button type="primary" :icon="FolderAdd" plain round>新建文件夹</el-button>
                <template #trigger v-show="false"></template>
            </el-upload>

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
                                <el-icon>
                                    <FolderOpened/>
                                </el-icon>
                                <span style="margin-left: 10px">{{ scope.row.name }}</span>
                            </div>
                        </template>
                    </el-table-column>
                    <el-table-column label="修改时间" width="180">
                        <template #default="scope">
                            <div>{{ scope.row.updated }}</div>
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

  <!--  file  -->
    <el-dialog v-model="fileDialogVisible[1]" title="输入要更改的文件名">
        <el-form :model="selectedFile[0]">
            <el-form-item label="文件名">
                <el-input v-model="selectedFile[0].name"/>
            </el-form-item>
        </el-form>
        <template #footer>
      <span class="dialog-footer">
        <el-button @click="fileDialogVisible[1] = false">取消</el-button>
        <el-button type="primary" @click="renameFile()">
          确定
        </el-button>
      </span>
        </template>
    </el-dialog>

    <el-dialog v-model="fileCopyAndMoveDialog" title="选择文件夹">
        <el-table :data="fileMovableFolderList" highlight-current-row>
            <el-table-column label="" width="180">
                <template #default="scope">
                    <div style="display: flex; align-items: center">
                        <div @click="toFolder(true, scope.row.id)">
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
            确定要<span style="color: red"> 删除 {{ selectedFile.map(file => file.name).join('，') }} </span>吗？
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

  <!--  folder  -->
    <el-dialog v-model="folderDialogVisible[1]" title="输入要更改的文件夹名">
        <el-form :model="selectedFolders[0]">
            <el-form-item label="文件名">
                <el-input v-model="selectedFolders[0].name"/>
            </el-form-item>
        </el-form>
        <template #footer>
      <span class="dialog-footer">
        <el-button @click="fileDialogVisible[1] = false">取消</el-button>
        <el-button type="primary" @click="renameFolder()">
          确定
        </el-button>
      </span>
        </template>
    </el-dialog>

    <el-dialog v-model="folderCopyAndMoveDialog" title="选择文件夹">
        <el-table :data="folderMovableFolderList" highlight-current-row>
            <el-table-column label="" width="180">
                <template #default="scope">
                    <div style="display: flex; align-items: center">
                        <div @click="toFolder(false, scope.row.id)">
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
        <el-button @click="folderCopyAndMoveDialog = false">取消</el-button>
        <el-button type="primary" @click="fileCopyAndMoveConfirm()">
          确定
        </el-button>
      </span>
        </template>
    </el-dialog>

    <el-dialog v-model="folderDialogVisible[4]" title="删除文件夹">
        <h3>
            <el-icon>
                <Warning/>
            </el-icon>
            确定要<span style="color: red"> 删除 {{ selectedFolders.map(folder => folder.name).join('，') }} </span>吗？
        </h3>
        <template #footer>
      <span class="dialog-footer">
        <el-button @click="folderDialogVisible[4] = false">取消</el-button>
        <el-button type="primary" @click="deleteFoldersConfirm()">
          确定
        </el-button>
      </span>
        </template>
    </el-dialog>
</template>

<script lang="ts" setup>
import {onMounted, reactive, ref} from 'vue'
import {ElMessage, ElTable, UploadFile, UploadFiles, UploadInstance} from 'element-plus'
import type {Folder, File} from './fileFolder.ts'
import axios, {AxiosProgressEvent} from "axios";
import {
    FolderOpened, Upload, Select, FolderAdd, Download, CopyDocument,
    EditPen, DeleteFilled, Rank, Warning
} from '@element-plus/icons-vue'
import {
    getFolderItems, updateFileName, listFileMovableFolders,
    deleteFiles, copyFiles, moveFiles, updateFolderName
} from "./fileFolder.ts";

let props = defineProps(["folderId"]);
let folderId = props.folderId
let fileButtonsState = ref(0)
let folderButtonsState = ref(0)
const fileTableRef = ref<InstanceType<typeof ElTable>>()

let folderList = reactive<Folder[]>([
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
])
let fileList = reactive<File[]>([
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
])
let fileMovableFolderList = reactive<Folder[]>(folderList)
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

const fileDialogVisible = reactive([false, false, false, false, false])
let listFoldersCurrentFolderId = 0
let fileCopyAndMoveDialog = ref(false)
let fileCopyAndMoveFlag
let selectedFile: File[]

const folderDialogVisible = reactive([false, false, false, false, false])
// let listFoldersCurrentFolderId = 0
let folderCopyAndMoveDialog = ref(false)
let folderCopyAndMoveFlag
let selectedFolders: Folder[]

// file
function fileButton(option: number) {
    selectedFile = fileTableRef.value!.getSelectionRows()
    if (!selectedFile) {
        return
    }
    if (option === 2 || option === 3) {
        toFolder(true, listFoldersCurrentFolderId)
        fileCopyAndMoveDialog.value = true
        fileCopyAndMoveFlag = option
        console.log(fileCopyAndMoveFlag)
        return
    }
    fileDialogVisible[option] = true
}

async function renameFile() {
    await updateFileName(selectedFile[0])
    await listFolderItems()
}

async function toFolder(isFile: boolean, folderId: number) {
    if (isFile) {
        const resp = await listFileMovableFolders(folderId)
        if (resp && resp.code === 0) {
            fileMovableFolderList = resp.data
        }
    } else {

    }

    listFoldersCurrentFolderId = folderId
}

async function fileCopyAndMoveConfirm() {
    console.log(fileCopyAndMoveFlag)
    const fileIds = selectedFile.map(file => file.id);
    if (fileCopyAndMoveFlag === 2) {
        await moveFiles(listFoldersCurrentFolderId, fileIds)
    } else if (fileCopyAndMoveFlag === 3) {
        await copyFiles(listFoldersCurrentFolderId, fileIds)
    }
    listFoldersCurrentFolderId = 0
}

async function deleteFilesConfirm() {
    await deleteFiles(selectedFile.map(file => file.id))
    await listFolderItems()
}

// folder
async function renameFolder() {
    await updateFolderName(selectedFile[0])
    await listFolderItems()
}



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
.button-group {
    margin-bottom: 13px;
}
</style>