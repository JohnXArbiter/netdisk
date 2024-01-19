<template>
    <el-row>
        <el-col :span="19">
            <router-view :key="$route.fullPath" @selectedIds="selectedItems"/>
        </el-col>

        <el-col :span="5">
            <file-info v-model:idObj="idObj"/>
        </el-col>
    </el-row>
</template>

<script lang="ts" setup>
import FileInfo from "../../components/files/info/FileInfo.vue";
import {reactive, watch} from "vue";

const ids: {
  files: number[],
  folders: number[]
} = {
  files: [],
  folders: []
}

const idObj = reactive({
  data: {
    id: 0,
    type: 0
  }
})

function selectedItems(ids: number[], forFile: boolean) {
  console.log(794509453, ids)
  if (forFile) {
    ids.files = ids
  } else {
    ids.folders = ids
  }
}

watch(() => ids, (newV) => {
  if (newV.files.length === 1 && newV.folders.length < 1) {
    idObj.data.id = newV.files[0]
    idObj.data.type = 0
  } else if (newV.folders.length == 1 && newV.files.length < 1) {
    idObj.data.id = newV.folders[0]
    idObj.data.type = 1
  } else {
    idObj.data.type = 2
  }
  console.log(idObj)
}, {deep: true})
</script>

<style scoped>

</style>