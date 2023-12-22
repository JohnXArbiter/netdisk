<template>
  <el-form :model="loginForm" label-width="120px">
    <el-form-item label="账号">
      <el-input v-model="loginForm.username"/>
    </el-form-item>
    <el-form-item label="密码">
      <el-input v-model="loginForm.password"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="login">登录</el-button>
    </el-form-item>
  </el-form>
</template>

<script lang="ts" setup>
import {reactive} from "vue";
import {loginPost, LoginReq} from "./loginForm.ts";

let loginForm = reactive<LoginReq>({
  password: "",
  username: "",
})

const login = async () => {
  console.log(loginForm.password, loginForm.username)
  const resp = await loginPost(loginForm)
  if (resp.code === 0) {
    localStorage.setItem("token", resp.data.token)
  }
}

// const ruleFormRef = ref<FormInstance>()
//
// const checkAge = (rule: any, value: any, callback: any) => {
//   if (!value) {
//     return callback(new Error('Please input the age'))
//   }
//   setTimeout(() => {
//     if (!Number.isInteger(value)) {
//       callback(new Error('Please input digits'))
//     } else {
//       if (value < 18) {
//         callback(new Error('Age must be greater than 18'))
//       } else {
//         callback()
//       }
//     }
//   }, 1000)
// }
//
// const validatePass = (rule: any, value: any, callback: any) => {
//   if (value === '') {
//     callback(new Error('Please input the password'))
//   } else {
//     if (ruleForm.checkPass !== '') {
//       if (!ruleFormRef.value) return
//       ruleFormRef.value.validateField('checkPass', () => null)
//     }
//     callback()
//   }
// }
// const validatePass2 = (rule: any, value: any, callback: any) => {
//   if (value === '') {
//     callback(new Error('Please input the password again'))
//   } else if (value !== ruleForm.pass) {
//     callback(new Error("Two inputs don't match!"))
//   } else {
//     callback()
//   }
// }
//
// const ruleForm = reactive({
//   pass: '',
//   checkPass: '',
//   age: '',
// })
//
// const rules = reactive<FormRules<typeof ruleForm>>({
//   pass: [{ validator: validatePass, trigger: 'blur' }],
//   checkPass: [{ validator: validatePass2, trigger: 'blur' }],
//   age: [{ validator: checkAge, trigger: 'blur' }],
// })
//
// const submitForm = (formEl: FormInstance | undefined) => {
//   if (!formEl) return
//   formEl.validate((valid) => {
//     if (valid) {
//       console.log('submit!')
//     } else {
//       console.log('error submit!')
//       return false
//     }
//   })
// }
//
// const resetForm = (formEl: FormInstance | undefined) => {
//   if (!formEl) return
//   formEl.resetFields()
// }

</script>

<style scoped>

</style>