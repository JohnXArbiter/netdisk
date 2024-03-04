<template>
    <div class="container-body article-list-body" :style="{ width: '1000px' }">
        <div class="article-panel">

            <el-form
                    label-position="top"
                    label-width="auto"
                    :model="infoInputs"
                    style="max-width: 600px;
                    padding: 50px 0 0 100px;"
                    :show-file-list="false"
            >
                <el-form-item>
                    <el-avatar :size="200"
                               :src="`https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png`"/>
                </el-form-item>
                <el-form-item>

                    <el-upload
                            v-model:file-list="fileList"
                            class="upload-demo"
                            :limit="1"
                            :http-request="changeAvatar"
                    >
                        <el-button type="primary">上传头像</el-button>
                        <template #tip>
                            <div class="el-upload__tip">
                                接受 jpg/png 格式，并且大小小于等于1MB
                            </div>
                        </template>
                    </el-upload>
                </el-form-item>


                <el-form-item label="名称">
                    <el-input v-model="infoInputs.name"/>
                </el-form-item>
                <el-form-item label="账号">
                    <el-input v-model="infoInputs.username"/>
                </el-form-item>
                <el-form-item label="个性签名">
                    <el-input v-model="infoInputs.signature"/>
                </el-form-item>
                <el-form-item>
                    <el-button>确定</el-button>
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>

<script lang="ts" setup>

import {onMounted, reactive} from "vue";
import {UploadRawFile} from "element-plus/es/components/upload/src/upload";
import {updateAvatar} from "@/components/user/info.ts";
import {codeOk} from "@/utils/apis/base.ts";
import {useBaseStore} from "@/store";

const baseStore = useBaseStore(),
    infoInputs: { data: {} } = reactive({data: {}})

let user = baseStore.userInfoInit

async function showUserInfo() {
    if (user.data.id === 0) {
        user.data = await useBaseStore().getUserInfo()
        console.log(user.data.name)
    }
}

async function changeAvatar(file: UploadRawFile) {
    // if (file.name)
    const formData = new FormData();
    formData.append('file', file)
    const resp = await updateAvatar(formData)
    if (resp.code === codeOk) {
        user.data.avatar = resp.data.url
        await baseStore.updateUserInfo(user.data, false)
    }
}

onMounted(() => {
    showUserInfo()
})

</script>

<style lang="scss" scoped>
.container-body {
    margin: 0 auto;
}

.article-panel {
    background: #fffced;
}
</style>