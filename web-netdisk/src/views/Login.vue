<template>
  <div>
    <div v-show="flag">
      <button >登录</button>
      <button v-show="flag"
              class="change"
              @click="change(1)">注册
      </button>
    </div>
    <div v-show="!flag">
      <button >注册</button>
      <button v-show="!flag"
              class="change"
              @click="change(0)">登录
      </button>
    </div>
    <div v-show="flag">
      <login-form></login-form>
    </div>
    <div v-show="!flag">
      <register-form></register-form>
    </div>
  </div>

</template>

<script lang="ts" setup>
import LoginForm from '../components/LoginForm.vue'
import RegisterForm from "../components/RegisterForm.vue"
import {ref} from "vue";

let flag = ref<boolean>(true)

const change = (num: number) => {
  flag.value = num === 0
  console.log(flag.value)
}

</script>

<style>

</style>