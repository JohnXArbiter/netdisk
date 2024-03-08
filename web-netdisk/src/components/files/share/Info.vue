<template>
    <div class="form-div">
        <el-row>
            <el-col :span="24">
                <el-empty v-if="!list.items || list.items.length==0"
                          description="当前分享文件夹没有文件 😺"/>

                <div v-if="list.items && list.items.length!=0"
                     style="margin: 20px 0">
                    <div style="font-size: 2rem; font-weight: 700;
                        margin-bottom: 10px">
                        {{ list.name }} 等文件...
                        <span style="float: right;">
                            <el-button v-if="userId != list.owner"
                                       size="large" type="primary"
                                       @click="downloadFiles">
                                <el-icon>
                                    <Download/>
                                </el-icon>&nbsp;
                                下载
                            </el-button>
                            <el-button-group v-if="userId == list.owner">
                                <el-button plain size="large" type="primary"
                                           @click="downloadFiles">
                                <el-icon>
                                    <Download/>
                                </el-icon>&nbsp;
                                下载
                            </el-button>
                                <el-button plain size="large" type="danger"
                                @click="dialogVisible = true">
                                <el-icon><CircleClose/></el-icon>&nbsp;
                                取消分享
                            </el-button>
                            </el-button-group>
                        </span>
                    </div>
                    <div style="font-size: 1.4rem">
                        <el-icon>
                            <Clock/>
                        </el-icon>&nbsp;
                        <span>{{ list.created }}</span>
                        <span style="margin-left: 50px;">{{ state }}</span>
                    </div>
                </div>

                <el-table v-if="list.items && list.items.length!=0"
                          ref="fileTableRef"
                          :data="list.items" style="width: 100%"
                          @selection-change="fileSelectionChange"
                >
                    <el-table-column type="selection" width="55"/>
                    <el-table-column label="文件名" min-width="500">
                        <template #default="scope">
                            <div style="display: flex; align-items: center">
                                <el-image :src="`/src/assets/alt_type${scope.row.type}.jpg`"
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
                    <el-table-column label="大小" min-width="200">
                        <template #default="scope">
                            <div>{{ scope.row.sizeStr }}</div>
                        </template>
                    </el-table-column>
                </el-table>
            </el-col>
        </el-row>

        <el-dialog v-model="dialogVisible" title="删除分享">
            <h3>
                <el-icon>
                    <Warning/>
                </el-icon>
                确定删除这个分享吗😶
            </h3>
            <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false;">取消</el-button>
        <el-button type="primary" @click="cancelShare()">
          确定
        </el-button>
      </span>
            </template>
        </el-dialog>
    </div>
</template>

<script lang="ts" setup>
import {ElTable} from "element-plus";
import {onMounted, reactive, ref} from "vue";
import {listFilesByShareId, ShareItem} from "@/components/files/share/Info.ts";
import {formatSize, formatState} from "@/utils/util.ts";
import {
    Clock, Download, CircleClose, Warning
} from "@element-plus/icons-vue";
import {useBaseStore} from "@/store";
import {useRoute} from "vue-router";

const route = useRoute()

let pwd = route.query.pwd | ''

const props = defineProps(['shareId']),
    fileTableRef = ref<InstanceType<typeof ElTable>>(),
    baseStore = useBaseStore(),
    list = reactive<{
        name: string
        created: string
        expired: number
        owner: number
        items: ShareItem[]
    }>({
        name: '',
        created: '',
        expired: 0,
        owner: 0,
        items: []
    })

let state = '',
    userId: number = -1,
    selected = [],
    dialogVisible = ref(false)

const listItems = async () => {
    console.log(props.shareId)
    let resp = await listFilesByShareId(props.shareId, pwd.toString())
    if (resp.code === 0 && resp.data) {
        list.name = resp.data.name
        list.created = resp.data.created
        list.items = resp.data.items
        list.owner = resp.data.owner
        state = formatState(resp.data.expired)
        list.items.forEach(item => {
            item.sizeStr = formatSize(item.size)
        })
    }
}

async function downloadFiles() {
    selected = fileTableRef.value!.getSelectionRows()
    selected.forEach(item => {
        window.open(item.url)
    })
}

onMounted(async () => {
    if (baseStore.getToken() !== '') {
        let userInfo = await baseStore.getUserInfo()
        userId = userInfo.id
    }
    await listItems()
})
</script>

<style scoped>
.small-pic {
    width: 40px;
    height: 40px;
    border-radius: 5px;
}

.form-div {
    background: rgba(255, 194, 133, 0.5);
    padding: 10px;
    border: 1px solid lightgray;
    border-radius: 10px;
    display: flex;
    justify-content: center;
}
</style>