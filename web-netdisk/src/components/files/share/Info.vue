<template>
    <div class="form-div">
        <el-row>

            <el-col :span="24" v-if="pwd === '0'">
                <div style="position: absolute; left: -60%;color: #adabab;
    font: 800 23px Arial, sans-serif; line-height: 100%;">提取文件
                </div>
                <div v-if="ownerInfo.data.shareStatus === shareNotExistOrDeleted"
                     class="small-zi">
                    当前分享已被删除或者不存在！😣
                </div>
                <div v-else-if="ownerInfo.data.shareStatus === shareIllegal"
                     class="small-zi">
                    当前分享已被违法封禁！😡
                </div>
                <div v-else-if="ownerInfo.data.userStatus === userBanned"
                     class="small-zi">
                    当前用户已被违法封禁！😡
                </div>

                <div v-else>
                    <div class="pwd-box">
                        <el-image class="big-pic"
                                  :src="ownerInfo.data.avatar"
                                  fit="cover"
                        />
                        <div style="position: relative; top: -100px; right: -115px">
                            <div style="font-size: 2rem; font-weight: 700">
                                {{ ownerInfo.data.name }}
                            </div>
                            <div style="position: absolute; margin-top: 20px">
                                <span v-if="ownerInfo.data.signature == ''">暂无签名</span>
                                <span v-else>{{ ownerInfo.data.signature }}</span>
                            </div>
                        </div>

                        <el-form label-position="top">
                            <el-form-item label="请先输入提取码：" size="large">
                                <el-input v-model="pwdInput"/>
                            </el-form-item>
                            <el-form-item>
                                <el-button size="large" style="width: 100%;"
                                           type="primary"
                                           @click="listItems(pwdInput)">提取文件
                                </el-button>
                            </el-form-item>
                        </el-form>
                    </div>
                </div>
            </el-col>

            <el-col v-else-if="list.items && list.items.length!=0"
                    :span="24" style="margin-bottom: 100px">
                <!--                <el-empty v-if="!list.items || list.items.length==0"-->
                <!--                          description="当前分享文件夹没有文件 😺"/>-->

                <div style="margin: 20px 0">
                    <div style="font-size: 2rem; font-weight: 700; margin-bottom: 10px">
                        {{ list.name }} 等文件...
                        <span style="float: right;">
                            <el-button v-if="ownerInfo.data.userId != list.owner"
                                       size="large" type="primary"
                                       @click="downloadFiles">
                                <el-icon>
                                    <Download/>
                                </el-icon>&nbsp;
                                下载
                            </el-button>
                            <el-button-group v-if="ownerInfo.data.userId == list.owner">
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

                <el-table
                        ref="fileTableRef"
                        :data="list.items" style="width: 100%"
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
            <el-footer>Copyright © 2024 咪咪网盘</el-footer>
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
import {getOwnerInfoByShareId, listFilesByShareId, ShareItem} from "@/components/files/share/Info.ts";
import {formatSize, formatState} from "@/utils/util.ts";
import {
    Clock, Download, CircleClose, Warning
} from "@element-plus/icons-vue";
import {useBaseStore} from "@/store";
import {useRoute} from "vue-router";
import {codeOk, promptError} from "@/utils/apis/base.ts";
import {shareIllegal, shareNotExistOrDeleted, userBanned} from "@/utils/constant.ts";

const route = useRoute()

let pwd = ref<string>(String(route.query.pwd | '')),
    pwdInput = ref('')

const props = defineProps(['shareId']),
    fileTableRef = ref<InstanceType<typeof ElTable>>(),
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
    }),
    ownerInfo = reactive({
        data:
            {
                shareStatus: 0,
                userId: -1,
                name: '',
                avatar: '',
                signature: '',
                userStatus: 0,
            }
    })

let state = '',
    selected = [],
    dialogVisible = ref(false)

const listItems = async (pwdStr: string) => {
    pwd.value = pwdStr
    let resp = await listFilesByShareId(props.shareId, pwdStr)
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
    if (state == '已过期') {
        promptError('当前分享已经过期了🥲')
        return
    }
    selected = fileTableRef.value!.getSelectionRows()
    selected.forEach(item => {
        window.open(item.url)
    })
}

async function getOwnerInfo() {
    const resp = await getOwnerInfoByShareId(props.shareId)
    if (resp.code === codeOk) {
        ownerInfo.data = resp.data
        console.log(ownerInfo.data.userStatus, ownerInfo.data.shareStatus)
    }
}

onMounted(async () => {
    await getOwnerInfo()

    if (pwd.value != '0') {
        await listItems(pwd.value)
    }
})
</script>

<style scoped>
.small-pic {
    width: 40px;
    height: 40px;
    border-radius: 5px;
}

.big-pic {
    width: 100px;
    height: 100px;
    border-radius: 50%;
}

.form-div {
    background: rgba(255, 194, 133, 0.5);
    padding: 10px;
    border: 1px solid lightgray;
    border-radius: 10px;
    display: flex;
    justify-content: center;
}

.pwd-box {
    margin: 200px 0;
}

.small-zi {
    font-weight: 700;
    font-size: 3rem;
    margin: 250px 0;
}
</style>