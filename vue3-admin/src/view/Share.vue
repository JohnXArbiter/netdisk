<template>
    <div class="card-list">
        <el-card>
            <el-input style="width:440px" @clear="searchUser" clearable v-model="searchForm.name"
                      placeholder="请输入用户姓名" class="input-with-select">
                <template #append>
                    <el-button icon="Search" @click="searchUser"/>
                </template>
            </el-input>
            <el-table :data="shares" border style="width: 100%;margin-top:20px">
                <el-table-column prop="id" label="ID" width="180"/>
                <el-table-column prop="pwd" label="密码" width="180"/>
                <el-table-column prop="name" label="名字" width="180"/>
                <el-table-column prop="userId" label="用户ID" width="180"/>
                <el-table-column prop="created" label="创建时间" width="180"/>
                <el-table-column prop="expired" label="过期时间" width="180"/>
                <el-table-column prop="downloadNum" label="下载次数" width="180"/>
                <el-table-column prop="clickNum" label="点击次数" width="180"/>
                <el-table-column prop="status" label="状态" width="180"/>
                <el-table-column label="操作" width="330">
                    <template #default="scope">
                        <el-button type="danger" size="small" @click="deleteUser(scope.row.id)">删除</el-button>
                        <el-button size="small"
                                   @click="() => router.push({ path: '/user/detail', query: { id: scope.row.id } })">详情
                        </el-button>
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

const router = useRouter();
// Dom 挂载之后
onMounted(() => {
    getUserList();
})
// 用户数据
let shares = ref([]);
let total = ref(0);
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
    total.value = res.data.data.total;
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