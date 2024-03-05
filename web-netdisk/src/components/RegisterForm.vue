<template>
    <el-form :model="form" label-width="120px">
        <el-form-item label="账号" size="large">
            <el-input v-model="form.data.username"/>
        </el-form-item>
        <el-form-item label="密码" size="large">
            <el-input v-model="form.data.password"/>
        </el-form-item>
        <el-form-item label="请再次输入" size="large">
            <el-input v-model="form.data.passwordRepeat"/>
        </el-form-item>
        <el-form-item label="邮箱" size="large">
            <el-input v-model="form.data.email"
                      placeholder="用于接收验证码"
            >
                <template #append>
                    <el-button @click="sendCode2Email(form.data.code)">发送验证码</el-button>
                </template>
            </el-input>
        </el-form-item>
        <el-form-item label="验证码" size="large">
            <el-input v-model="form.data.code"/>
        </el-form-item>

        <el-form-item>
            <el-button type="primary" size="large" @click="register">注册</el-button>
        </el-form-item>
    </el-form>
</template>

<script lang="ts" setup>
import {reactive} from "vue";
import {registerPost} from "./registerForm.ts";
import {promptSuccess} from "../utils/apis/base.ts";
import {sendCode2Email} from "@/utils/util.ts";

let form = reactive({
    data: {
        username: '',
        password: '',
        passwordRepeat: '',
        email: '',
        code: ''
    }
})

async function register() {
    const resp = await registerPost(form.data)
    if (resp.code === 0) {
        promptSuccess('注册成功！😺')
        setTimeout(() => {
            window.location.reload()
        }, 2000)
    }
}

</script>

<style scoped>

</style>