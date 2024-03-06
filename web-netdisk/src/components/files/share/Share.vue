<template>
    <el-row>
        <el-col :span="24">
            <div class="file-table">

                <el-empty v-if="!shareList.data || shareList.data.length==0"
                          description="文件列表为空，上传你的第一个文件吧！😺"/>

                <el-table v-if="shareList && shareList.data.length!=0"
                          ref="fileTableRef"
                          :data="shareList.data" style="width: 100%"
                          @selection-change="fileSelectionChange"
                >
                    <el-table-column type="selection" width="55"/>
                    <el-table-column label="文件名" min-width="500">
                        <template #default="scope">
                            <div style="display: flex; align-items: center">
                                <el-image v-if="scope.row.type === typeImage"
                                          class="small-pic"
                                          :src="scope.row.url"
                                          alt="../../assets/alt_type1.jpg"
                                          :fit="'cover'"/>
                                <el-image v-else
                                          :src="`/src/assets/alt_type${scope.row.type}.jpg`"
                                          alt=""
                                          class="small-pic"
                                          :fit="'cover'"/>
                                <span style="margin-left: 5px">{{ scope.row.name }}</span>
                            </div>
                        </template>
                    </el-table-column>
                    <el-table-column label="修改时间" min-width="200">
                        <template #default="scope">
                            <div>{{ scope.row.updated }}</div>
                        </template>
                    </el-table-column>
                    <el-table-column label="大小" min-width="100">
                        <template #default="scope">
                            <div>{{ scope.row.sizeStr }}</div>
                        </template>
                    </el-table-column>
                </el-table>
            </div>
        </el-col>
    </el-row>

</template>

<script lang="ts" setup>
import {CopyDocument, DeleteFilled, Download, EditPen, Rank, Share} from "@element-plus/icons-vue";
import {typeImage} from "@/utils/constant.ts";
import {ElTable} from "element-plus";
import {reactive} from "vue";
import {Folder} from "@/components/files/folder.ts";

let folderList = reactive<{ data: Folder[] }>({
    data: []
})

</script>

<style scoped>

</style>