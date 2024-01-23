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
import {useFileFolderStore} from "../../store/fileFolder.ts";
import SparkMD5 from 'spark-md5'
import {checkFile, upload, uploadConst} from "./uploading.ts";
import {codeOk, promptError, promptSuccess} from "../../utils/apis/base.ts";

const fileFolderStore = useFileFolderStore()

async function handleUpload(param: UploadRequestOptions) {
    const checkRes = await checkBeforeUpload(param.file)
    if (checkRes.success) {
        if (checkRes.status === uploadConst.codeNeedUpload) {
            if (param.file.size > uploadConst.sliceSize) {
                uploadSlice(param.file, checkRes.fileId, 0);
            } else {
                await uploadSingle(param.file, checkRes.fileId)
            }
        } else {
            promptSuccess(param.file.name + ' 上传成功！😺')
        }
    } else {
        promptError('请检查文件是否合法！')
    }
}

async function checkBeforeUpload(file: UploadRawFile) {
    const size = file.size
    const checkReq = {
        folderId: fileFolderStore.folderId,
        name: file.name,
        size: size,
        ext: file.name.substring(file.name.lastIndexOf('.')),
        hash: genMd5(file)
    }

    const resp = await checkFile(checkReq)
    if (resp && resp.code === codeOk) {
        return {
            success: true,
            fileId: resp.data.fileId,
            status: resp.data.status
        }
    }
    return {
        success: false,
        fileId: 0,
        status: 0
    }
}

function genMd5(file: UploadRawFile) {
    const spark = new SparkMD5.ArrayBuffer()
    spark.append(file)
    return spark.end()
}

function uploadSlice(file: UploadRawFile, fileId: number, idx: number) {

}

async function uploadSingle(file: UploadRawFile, fileId: number) {
    const formData = new FormData();
    formData.append('file', file)
    formData.append('fileId', fileId.toString())
    const resp = await upload(formData)
    if (resp && resp.code === codeOk) {
        ElMessage.success('上传成功')
    }
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