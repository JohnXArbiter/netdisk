<template>
    <el-form :model="registerForm" label-width="120px">
        <el-form-item label="账号" size="large">
            <el-input v-model="registerForm.username"/>
        </el-form-item>
        <el-form-item label="密码" size="large">
            <el-input v-model="registerForm.password"/>
        </el-form-item>
        <el-form-item label="请再次输入" size="large">
            <el-input v-model="registerForm.passwordRepeat"/>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" size="large" @click="register">登录</el-button>
        </el-form-item>
    </el-form>
</template>

<script lang="ts" setup>
import {reactive} from "vue";
import {registerPost, RegisterReq} from "./registerForm.ts";

let registerForm = reactive<RegisterReq>({
    username: '',
    password: '',
    passwordRepeat: ''
})

const register = async () => {
    const resp = await registerPost(registerForm)
    if (resp.code === 0) {

    }
}


</script>

<style scoped>

</style>