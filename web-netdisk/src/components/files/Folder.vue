<template>
    <el-row>
        <el-col :span="14">
            <file :folderId="folderId"/>
        </el-col>

        <!--  folder  -->
        <el-col :span="10">
            <div class="folder-table">
                <el-upload v-if="folderButtonsState === 0">
                    <el-button type="primary" :icon="FolderAdd" round @click="folderButton(0)">新建文件夹</el-button>
                    <template #trigger v-show="false"></template>
                </el-upload>
                <div class="button-group">
                    <template v-if="folderButtonsState !== 0">
                        <el-button-group>
                            <el-button type="primary" round plain :icon="Download" @click="folderButton(1)">下载
                            </el-button>
                            <template v-if="folderButtonsState === 1">
                                <el-button type="primary" round plain :icon="EditPen" @click="folderButton(2)">重命名
                                </el-button>
                            </template>
                            <el-button type="primary" round plain :icon="Rank" @click="folderButton(3)">移动</el-button>
                            <el-button type="primary" round plain :icon="CopyDocument" @click="folderButton(4)">复制
                            </el-button>
                            <el-button type="danger" round plain :icon="DeleteFilled" @click="folderButton(5)">删除
                            </el-button>
                        </el-button-group>
                    </template>
                </div>

                <el-table ref="folderTableRef"
                          :data="folderList" style="width: 100%"
                          @selection-change="folderSelectionChange"
                >
                    <el-table-column type="selection" width="55"/>
                    <el-table-column label="文件夹名" width="180">
                        <template #default="scope">
                            <div class="file-folder-row" style="display: flex; align-items: center">
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

  <!--  folder  -->
    <el-dialog v-model="folderDialogVisible[0]" title="创建文件夹">
        <el-form :model="selectedFolders[0]" label-width="120px">
            <el-form-item label="文件夹名">
                <el-input v-model="selectedFolders[0].name"/>
            </el-form-item>
        </el-form>
        <template #footer>
      <span class="dialog-footer">
        <el-button @click="folderDialogVisible[0] = false">取消</el-button>
        <el-button type="primary" @click="createFolderConfirm()">
          确定
        </el-button>
      </span>
        </template>
    </el-dialog>

    <el-dialog v-model="folderDialogVisible[2]" title="输入要更改的文件夹名">
        <el-form :model="selectedFolders[0]">
            <el-form-item label="文件名">
                <el-input v-model="selectedFolders[0].name"/>
            </el-form-item>
        </el-form>
        <template #footer>
      <span class="dialog-footer">
        <el-button @click="folderDialogVisible[2] = false">取消</el-button>
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
        <el-button @click="folderCopyAndMoveDialog = false">取消</el-button>
        <el-button type="primary" @click="folderCopyAndMoveConfirm()">
          确定
        </el-button>
      </span>
        </template>
    </el-dialog>

    <el-dialog v-model="folderDialogVisible[5]" title="删除文件夹">
        <h3>
            <el-icon>
                <Warning/>
            </el-icon>
            确定要<span style="color: red"> 删除 {{ selectedFolders.map(folder => folder.name).join('，') }} </span>吗？
            你可以在回收站中找到他们。
        </h3>
        <template #footer>
      <span class="dialog-footer">
        <el-button @click="folderDialogVisible[5] = false">取消</el-button>
        <el-button type="primary" @click="deleteFoldersConfirm()">
          确定
        </el-button>
      </span>
        </template>
    </el-dialog>
</template>

<script lang="ts" setup>
import {onMounted, reactive, ref} from 'vue'
import {ElMessage, ElTable} from 'element-plus'
import type {Folder} from './folder.ts'
import {
    FolderOpened, FolderAdd, Download, CopyDocument,
    EditPen, DeleteFilled, Rank, Warning
} from '@element-plus/icons-vue'
import {
    updateFolderName,
    createFolder,
    listFolderMovableFolders,
    moveFolders,
    copyFolders,
    deleteFolders, listFoldersByParentFolderId
} from "./folder.ts";
import File from './File.vue'

let props = defineProps(["folderId"]);
let folderId = props.folderId
let folderButtonsState = ref(0)
const folderTableRef = ref<InstanceType<typeof ElTable>>()

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

const listFolders = async () => {
    const resp = await listFoldersByParentFolderId(folderId)
    if (resp.code === 0 && resp.data) {
        Object.assign(folderList, resp.data)
    }
}


let listFoldersCurrentFolderId = 0

const folderDialogVisible = reactive([false, false, false, false, false, false])
let folderCopyAndMoveDialog = ref(false)
let folderCopyAndMoveFlag
let selectedFolders: Folder[]
let folderMovableFolderList = reactive<Folder[]>(folderList)

async function toFolder(folderId: number) {
    const resp = await listFolderMovableFolders(folderId)
    if (resp && resp.code === 0) {
        Object.assign(folderMovableFolderList, resp.data)
        listFoldersCurrentFolderId = folderId
    }
}

// folder
function folderButton(option: number) {
    selectedFolders = folderTableRef.value!.getSelectionRows()
    if (option === 0) {
        folderDialogVisible[option] = true
        return
    }
    if (!selectedFolders) {
        return
    }
    if (option === 3 || option === 4) {
        toFolder(0)
        folderCopyAndMoveDialog.value = true
        folderCopyAndMoveFlag = option
        return
    }
    folderDialogVisible[option] = true
}

async function createFolderConfirm() {
    await createFolder(selectedFolders[0])
    await listFolders()
}

async function renameFolder() {
    await updateFolderName(selectedFolders[0])
    await listFolders()
}

async function folderCopyAndMoveConfirm() {
    const folderIds = selectedFolders.map(folder => folder.id);
    if (folderCopyAndMoveFlag === 3) {
        await moveFolders(listFoldersCurrentFolderId, folderIds)
    } else if (folderCopyAndMoveFlag === 4) {
        await copyFolders(listFoldersCurrentFolderId, folderIds)
    }
    listFoldersCurrentFolderId = 0
}

async function deleteFoldersConfirm() {
    await deleteFolders(selectedFolders.map(folder => folder.id))
    await listFolders()
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


onMounted(() => {
    listFolders()
})


</script>


<style scoped>
.button-group {
    margin-bottom: 13px;
}

.file-folder-row:hover {
    cursor: pointer;
    background-color: rgb(230, 230, 245);
    border-radius: 5px;
}
</style>