<template>
    <div class="card-list">
        <el-card>
            <el-input style="width:440px" @clear="searchUser" clearable v-model="searchForm.name"
                      placeholder="请输入用户姓名" class="input-with-select">
                <template #append>
                    <el-button icon="Search" @click="searchUser"/>
                </template>
            </el-input>
            <el-table :data="shares" border style="width: 100%; margin-top:20px">
                <el-table-column prop="id" label="ID" min-width="100"/>
                <el-table-column label="用户ID-昵称" min-width="200">
                    <template #default="scope">
                        {{ scope.row.id }} | {{ scope.row.name }}
                    </template>
                </el-table-column>
                <el-table-column prop="created" label="创建时间" min-width="150"/>
                <el-table-column label="类型" min-width="60">
                    <template #default="scope">
                        {{ typeMap[scope.row.type] }}
                    </template>
                </el-table-column>
                <el-table-column prop="downloadNum" label="下载次数" min-width="60"/>
                <el-table-column prop="clickNum" label="点击次数" min-width="60"/>
                <el-table-column prop="state" label="状态" min-width="180"/>
                <el-table-column label="操作" width="330">
                    <template #default="scope">
                        <el-button v-if="scope.row.type === typeMulti"
                                   type="primary" plain size="small"
                                   @click="openInfo(`/share/${scope.row.id}`)">进入查看
                        </el-button>
                        <el-button v-if="scope.row.type !== typeMulti"
                                   type="primary" plain size="small"
                                   @click="getUrl(scope.row.id)">下载查看
                        </el-button>
                        <el-button type="danger" size="small" @click="deleteUser(scope.row.id)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <!-- 分页 -->
            <el-pagination style="margin-top:20px" :current-page="searchForm.current" :page-size="searchForm.size"
                           :page-sizes="[10, 20, 30, 40]" layout="->,total, sizes, prev, pager, next, jumper"
                           :total="total"
                           @size-change="handleSizeChange" @current-change="handleCurrentChange"/>
        </el-card>
    </div>
</template>

<script setup>
import shareApi from "@/api/share.js";
import {onMounted, reactive, ref} from "vue";
import {ElMessage, ElMessageBox} from 'element-plus';
import {useRouter} from 'vue-router'
import {formatState} from "@/utils/util.js";
import {typeMap, typeMulti} from "../../utils/constant.js";
import {codeOk, promptError} from "@/utils/http/base.js";

const router = useRouter();
// Dom 挂载之后
onMounted(() => {
    getUserList();
})
// 用户数据
let shares = ref([])
let total = ref(0)
let urlMap = new Map()

// 搜索条件
const searchForm = reactive({
    current: 1,
    size: 10,
    name: ''
})

// 获取用户列表
const getUserList = async () => {
    const res = await shareApi.getShareList({'page': 0, 'size': 100});
    console.log(res.data);
    shares.value = res.data.data
    shares.value.forEach(share => {
        share.state = formatState(share.expired)
    })
    total.value = res.data.data.total;
}

function openInfo(link) {
    window.open(link)
}

async function getUrl(id) {
    if (urlMap.has(id)) {
        window.open(urlMap.get(id))
        return
    }
    const resp = await shareApi.getUrl(id)
    if (resp.data.code === codeOk) {
        urlMap.set(id, resp.data.data.url)
        console.log(resp.data.data)
        window.open(resp.data.data)
        return
    }
    promptError(`获取链接失败，${resp.data.msg}`)
}

const handleSizeChange = (size) => {
    searchForm.size = size;
    getUserList();
}
const handleCurrentChange = (current) => {
    searchForm.current = current;
    getUserList();
}
const searchUser = () => {
    searchForm.current = 1;
    getUserList();
}
// 删除用户
const deleteUser = (id) => {
    ElMessageBox.confirm(
        '确定要删除该用户信息吗?',
        {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
        }
    ).then(async () => {
        const res = await userApi.delUser({id: id});
        if (res.data.success) {
            ElMessage.success("删除成功")
            getUserList();
        } else {
            ElMessage.error("删除失败")
        }
    }).catch(() => {
        ElMessage({
            type: 'info',
            message: '取消删除',
        })
    })
}
</script>

<style lang="scss" scoped>
</style>