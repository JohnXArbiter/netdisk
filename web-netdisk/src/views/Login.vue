<template>
    <div class="bg">
        <div class="form-box">
            <div v-show="flag" class="text">
                <span class="big-text">登录</span>
                <button v-show="flag"
                        class="change"
                        @click="change(1)">注册
                </button>
            </div>
            <div v-show="!flag" class="text">
                <span class="big-text">注册</span>
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
    </div>
</template>

<script lang="ts" setup>
import LoginForm from '../components/LoginForm.vue'
import RegisterForm from "../components/RegisterForm.vue"
import {ref} from "vue";

let flag = ref<boolean>(true)

const change = (num: number) => {
    flag.value = num === 0
}

</script>

<style>
.bg {
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: #f7f7ff;
    width: 100vw;
    height: 100vh;
    overflow: hidden;
}

.form-box {
    background-color: #e3f1ff;
    width: 20%;
    border-radius: 20px;
    box-shadow: 5px 3px 3px #e5e5e5;
}

.text {
    margin: 2rem 3rem;
}

.big-text {
    font-size: 5rem;
    margin-right: 3rem;
    color: #787eef;
}


</style>