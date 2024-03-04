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
            </el-input>
        </el-form-item>
        <el-form-item label="验证码" size="large">
            <el-input v-model="form.data.code">
                <template #append>
                    <el-button @click="sendCode2Email()">发送</el-button>
                </template>
            </el-input>
        </el-form-item>

        <el-form-item>
            <el-button type="primary" size="large" @click="register">注册</el-button>
        </el-form-item>
    </el-form>
</template>

<script lang="ts" setup>
import {reactive} from "vue";
import {registerPost, RegisterReq, sendCode} from "./registerForm.ts";
import {promptError, promptSuccess} from "../utils/apis/base.ts";
import router from "../router";

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

async function sendCode2Email() {
    const resp = await sendCode(form.data.email)
    if (resp.code === 0) {
        promptSuccess('验证码已发送至邮件😊')
    }
}

</script>

<style scoped>

</style>