<template>
    <el-card>
        <el-table
                ref="fileTableRef" border
                :data="list.items" style="width: 100%"
        >
            <el-table-column type="selection" width="55"/>
            <el-table-column label="ID" min-width="200">
                <template #default="scope">
                    <div style="display: flex; align-items: center">
                        <el-image :src="`/src/assets/img/alt_type${scope.row.type}.jpg`"
                                  class="small-pic"
                                  :fit="'cover'"/>
                        <span style="margin-left: 5px">{{ scope.row.id }}</span>
                    </div>
                </template>
            </el-table-column>
            <el-table-column prop="name" label="文件名" min-width="400"/>
            <el-table-column label="修改时间" min-width="200">
                <template #default="scope">
                    <div>{{ scope.row.updated }}</div>
                </template>
            </el-table-column>
            <el-table-column label="大小" min-width="200">
                <template #default="scope">
                    <div>{{ scope.row.sizeStr }}</div>
                </template>
            </el-table-column>
            <el-table-column label="类型" min-width="200">
                <template #default="scope">
                    <div>{{ typeMap[scope.row.type] }}</div>
                </template>
            </el-table-column>
            <el-table-column prop="state" label="状态" min-width="100"/>
            <el-table-column label="操作" min-width="100">
                <template #default="scope">
                    <el-button type="primary" plain size="small">恢复</el-button>
                    <template v-if="scope.row.status === fileStatus.ok">
                        <el-button type="danger" size="small">封禁</el-button>
                    </template>
                </template>
            </el-table-column>
        </el-table>
    </el-card>

<!--    <el-dialog v-model="dialogVisible.option[1]" title="封禁文件">-->
<!--        <h3>-->
<!--            <el-icon>-->
<!--                <Warning/>-->
<!--            </el-icon>-->
<!--            确定删除这个分享吗😶-->
<!--        </h3>-->
<!--        <template #footer>-->
<!--      <span class="dialog-footer">-->
<!--        <el-button @click="dialogVisible.option[1]=false; cancelId=''">取消</el-button>-->
<!--        <el-button type="primary" @click="cancel()">-->
<!--          确定-->
<!--        </el-button>-->
<!--      </span>-->
<!--        </template>-->
<!--    </el-dialog>-->
</template>

<script lang="js" setup>
import {onBeforeMount, reactive} from "vue";
import {formatSize} from "@/utils/util.js";
import shareApi from "@/api/share.js";
import {fileStatus, fileStatusMap, typeMap} from "@/utils/constant.js";

const props = defineProps(["shareId"]),
    list = reactive({
        name: '',
        created: '',
        expired: 0,
        owner: 0,
        items: []
    }),
    dialogVisible = reactive({
        option: [false, false]
    })

const listItems = async () => {
    let resp = await shareApi.getShareFilesByShareId(props.shareId)
    if (resp.data.code === 0 && resp.data.data) {
        // list.name = resp.data.data.name
        // list.created = resp.data.data.created
        // list.items = resp.data.data.items
        // list.owner = resp.data.data.owner
        // state = formatState(resp.data.expired)
        list.items = resp.data.data.items
        list.items.forEach(item => {
            console.log(typeMap[item.type])
            item.state = fileStatusMap[item.status]
            item.sizeStr = formatSize(item.size)
        })
    }
}

onBeforeMount(async () => {
    await listItems()
})

</script>

<style scoped>
.small-pic {
    width: 40px;
    height: 40px;
    border-radius: 5px;
}</style>