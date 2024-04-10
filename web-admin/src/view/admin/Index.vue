<template>
    <div class="card-list">
        <el-card>
            <el-input style="width:440px" @clear="searchUser" clearable v-model="searchForm.name"
                      placeholder="请输入用户姓名" class="input-with-select">
                <template #append>
                    <el-button icon="Search" @click="searchUser"/>
                </template>
            </el-input>
            <el-button type="primary"
                       @click="dialogVisible.option[2] = true"
                       style="margin-left: 3rem">
                <el-icon>
                    <Plus/>
                </el-icon>
                &nbsp;新增管理员
            </el-button>
            <el-table :data="admins" border style="width: 100%; margin-top:20px">
                <el-table-column prop="id" label="ID" min-width="100"/>
                <el-table-column prop="name" label="名称" min-width="200"/>
                <!--                <el-table-column prop="password" label="密码" min-width="200"/>-->
                <el-table-column prop="created" label="创建时间" min-width="150"/>
                <el-table-column prop="updated" label="更新时间" min-width="150"/>
                <el-table-column label="类型" min-width="60">
                    <template #default="scope">
                        <div v-if="scope.row.status === 1">超管</div>
                        <div v-else>管理员</div>
                    </template>
                </el-table-column>
                <el-table-column prop="state" label="状态" min-width="180"/>
                <el-table-column label="操作" width="330">
                    <template #default="scope">
                        <el-button v-if="scope.row.status === adminSuper"
                                   type="danger" size="small" disabled
                                   @click="buttonClick(1, scope.row.id, adminBanned, scope.row.type)">停用
                        </el-button>
                        <el-button v-else
                                   type="danger" size="small"
                                   @click="buttonClick(1, scope.row.id, adminBanned, scope.row.type)">停用
                        </el-button>

                        <el-button v-if="scope.row.status === adminSuper"
                                   type="primary" size="small" disabled
                                   @click="buttonClick(0, scope.row.id, adminNormal, scope.row.type)">恢复
                        </el-button>
                        <el-button v-else
                                   type="primary" size="small"
                                   @click="buttonClick(0, scope.row.id, adminNormal, scope.row.type)">恢复
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

    <el-dialog v-model="dialogVisible.option[0]" title="启用账号">
        <h3>
            <el-icon>
                <Warning/>
            </el-icon>
            需要启用这个账号吗？😶
        </h3>
        <template #footer>
              <span class="dialog-footer">
                <el-button @click="dialogVisible.option[0]=false">取消</el-button>
                <el-button type="primary" @click="setStatus(0)">
                  确定
                </el-button>
              </span>
        </template>
    </el-dialog>

    <el-dialog v-model="dialogVisible.option[1]" title="停用账号">
        <h3>
            <el-icon>
                <Warning/>
            </el-icon>
            确定停用这个账号吗😶
        </h3>
        <template #footer>
              <span class="dialog-footer">
                <el-button @click="dialogVisible.option[1]=false">取消</el-button>
                <el-button type="primary" @click="setStatus(1)">
                  确定
                </el-button>
              </span>
        </template>
    </el-dialog>

    <el-dialog v-model="dialogVisible.option[2]" title="新增管理员">
        <el-form label-position="top">
            <el-form-item label="账号">
                <el-input v-model="addForm.username"></el-input>
            </el-form-item>
            <el-form-item label="密码">
                <el-input v-model="addForm.password"></el-input>
            </el-form-item>
            <el-form-item label="确认密码">
                <el-input v-model="addForm.password2"></el-input>
            </el-form-item>
            <el-form-item label="名称">
                <el-input v-model="addForm.name"></el-input>
            </el-form-item>
        </el-form>

        <template #footer>
              <span class="dialog-footer">
                <el-button @click="dialogVisible.option[2]=false">取消</el-button>
                <el-button type="primary" @click="addAdmin()">
                  确定
                </el-button>
              </span>
        </template>
    </el-dialog>
</template>

<script lang="js" setup>
import {Plus} from "@element-plus/icons-vue";
import adminApi from "@/api/admin.js";
import {onMounted, reactive, ref} from "vue";
import {ElMessage, ElMessageBox} from 'element-plus';
import {useRouter} from 'vue-router'
import {codeOk, promptError, promptSuccess} from "@/utils/http/base.js";
import {adminBanned, adminNormal, adminSuper} from "@/utils/constant.js";

const router = useRouter();

onMounted(() => {
    listAdmins()
})

let admins = ref([]),
    total = ref(0)

const searchForm = reactive({
        current: 1,
        size: 10,
        name: ''
    }),
    dialogVisible = reactive({option: [false, false, false]}),
    setStatusObj = {id: 0, status: 0},
    addForm = reactive({})

async function listAdmins() {
    const res = await adminApi.getAdminList({'page': 0, 'size': 100});
    console.log(res.data);
    admins.value = res.data.data
    admins.value.forEach(admin => {
        switch (admin.status) {
            case adminSuper:
                admin.state = '当前管理员账号'
                return
            case adminNormal:
                admin.state = '可用'
                return
            case adminBanned:
                admin.state = '停用'
                return
        }
    })
    total.value = res.data.data.total;
}

function buttonClick(option, id, status) {
    setStatusObj.id = id
    setStatusObj.status = status
    dialogVisible.option[option] = true
}

async function setStatus(option) {
    console.log(setStatusObj)
    const resp = await adminApi.setStatus(setStatusObj)
    if (resp.data.code === codeOk) {
        await listAdmins()
        promptSuccess('操作成功')
        dialogVisible.option[option] = false
        return
    }
    promptError(`操作失败，${resp.data.msg}`)
}

async function addAdmin() {
    const resp = await adminApi.add(addForm)
    if (resp.data.code === codeOk) {
        await listAdmins()
        promptSuccess('操作成功')
        dialogVisible.option[2] = false
        return
    }
    promptError(`操作失败，${resp.data.msg}`)
}

const handleSizeChange = (size) => {
    searchForm.size = size;
    listAdmins();
}

const handleCurrentChange = (current) => {
    searchForm.current = current;
    listAdmins();
}

const searchUser = () => {
    searchForm.current = 1;
    listAdmins();
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