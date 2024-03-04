<template>
  <div class="main-outer">
    <div class="container-body article-list-body">
      <div class="article-panel">

        <el-form
            label-position="top"
            label-width="auto"
            :model="infoInputs"
            style="max-width: 600px;
                    padding: 50px 0 0 100px;"
            class="demo-ruleForm"
            :show-file-list="false"
        >
          <el-form-item >
            <el-avatar :size="100"
                       :src="`https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png`"
                       style="margin-right:50px"

            />

            <el-upload
                v-model:file-list="fileList"
                class="upload-demo"
                :limit="1"
                :http-request="changeAvatar"
            >
              <el-button type="primary" style="margin-top: 40px">上传头像</el-button>
              <template #tip>
                <div class="el-upload__tip">
                  接受 jpg/png 格式，并且大小≤1MB
                </div>
              </template>
            </el-upload>

            <div class="progress">
              <proggress-dashboard/>
            </div>
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
            <el-button type="warning" size="large" style="margin-top: 30px">确定</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>

import {onMounted, reactive} from "vue";
import {UploadRawFile} from "element-plus/es/components/upload/src/upload";
import {updateAvatar} from "@/components/user/info.ts";
import {codeOk} from "@/utils/apis/base.ts";
import {useBaseStore} from "@/store";
import ProggressDashboard from "@/components/ProggressDashboard.vue";


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
.main-outer{
  //background-color: lightyellow;
  background: url("@/assets/background_grid.png");
  width: 100%;
  height: 94vh;
  margin: 0;
  padding: 0;
  position: absolute;
  top: 6vh;
  overflow: hidden;
}

.container-body {
    margin: 5vh auto;
    width: 45%;
}
.progress{
  position: absolute;
  float: right;
  right: 1px;

}

.article-panel {
    background: rgba(255,194,133, 0.5);
  padding: 10px;
  border: 1px solid lightgray;
  border-radius: 10px;
}
</style>