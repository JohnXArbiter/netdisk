<template>
    <el-upload ref="uploadFiles"
               class="upload-demo"
               action="actionUrl"
               multiple
               :limit="20"
               :auto-upload="true"
               :http-request="handleUpload"
    >
        <!--               :on-change="change"-->

        <el-button type="primary" :icon="Upload" round>选择上传</el-button>
        <!--        <div v-if="!progress" class="el-upload__text">-->
        <!--            Drop file here or <em>click to upload</em>-->
        <!--        </div>-->

        <!--        <el-progress-->

        <!--                v-else-->
        <!--                :text-inside="true"-->
        <!--                :stroke-width="24"-->
        <!--                :percentage="progress"-->
        <!--                status="success"-->
        <!--        />-->

    </el-upload>
</template>

<script lang="ts" setup>
import {Upload} from "@element-plus/icons-vue";
import {ElMessage, UploadRequestOptions} from "element-plus";
import {UploadRawFile} from "element-plus/es/components/upload/src/upload";

const props = defineProps([]);

async function handleUpload(param: UploadRequestOptions) {
    console.log('file', param.file);
    console.log('data', param.data);
    console.log('action', param.action);
    console.log('filename', param.filename);
    console.log('headers', param.headers);
    console.log('action', param.method);

    let checkRes = await checkBeforeUpload(param)
    if (checkRes.success) {
        uploadSlice(param.file, 0);

    } else {
        ElMessage.error('请检查文件是否合法！');
    }
}

async function checkBeforeUpload(param: UploadRequestOptions) {
  const qwe: CheckReq = {
    filename: param.filename,
    size: param.file.size,
    type: param.file.type
    ext: param.filename.substring(param.filename.lastIndexOf('.') + 1)
  }

  var spark = new SparkMD5.ArrayBuffer();
  spark.append(file);
  var md5 = spark.end();


  const res: CheckRes = {
    success: false,
    type: 0
  }


  param.filename

  return res
  // const fileType = file.file.name.split('.')
  // if (fileType[fileType.length - 1] !== 'zip' && fileType[fileType.length - 1] !== 'tar') {
  //     ElMessage.warning('文件格式错误，仅支持 .zip/.tar')
  //     return res
  // }
  //
  // // 校验文件大小
  // const fileSize = file.file.size;
  // // 文件大小是否超出 2G
  // if (fileSize > 2 * 1024 * 1024 * 1024) {
  //     ElMessage.warning('上传文件大小不能超过 2G')
  //     return false
  // }
  //
  // // 调用接口校验文件合法性，比如判断磁盘空间大小是否足够
  // const res = await checkMirrorFileApi()
  // if (res.code !== 200) {
  //     ElMessage.warning('暂时无法查看磁盘可用空间，请重试')
  //     return false
  // }
  // // 查看磁盘容量大小
  // if (res.data.diskDevInfos && res.data.diskDevInfos.length > 0) {
  //     let saveSize = 0
  //     res.data.diskDevInfos.forEach(i => {
  //         // 磁盘空间赋值
  //         if (i.devName === '/dev/mapper/centos-root') {
  //             // 返回值为GB，转为字节B
  //             saveSize = i.free * 1024 * 1024 * 1024
  //         }
  //     })
  //     // 上传的文件大小没有超出磁盘可用空间
  //     if (fileSize < saveSize) {
  //         return true
  //     } else {
  //         ElMessage.warning('文件大小超出磁盘可用空间容量')
  //         return false
  //     }
  // } else {
  //     ElMessage.warning('文件大小超出磁盘可用空间容量')
  //     return false
  // }

}

function uploadSlice(file: UploadRawFile, idx: number) {
}

// function asd(e: Event) {
//     const target = e.target
//     if (target instanceof HTMLInputElement) {
//         const file = target.files
//         if (file) {
//             const form = new FormData()
//             for (let i = 1; i < file.length; i++) {
//                 form.append("file", file[i])
//             }
//             axios.post("/", form, {
//                 onUploadProgress: (progressEvent: AxiosProgressEvent) => {
//                     Math.round((progressEvent.loaded / (progressEvent.total as number) * 100))
//                 }
//             })
//         }
//     }
// }

// const uploadProcedure = (options: UploadRequestOptions) => {
//     console.log(options.files)
//     options.files
//     return XMLHttpRequest
// }
</script>

<style scoped>

</style>