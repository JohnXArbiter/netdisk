<template>

</template>

<script lang="ts" setup>
import {reactive} from "vue/dist/vue";
import { registerPost} from "./registerForm.ts";

let registerForm = reactive<registerReq>({
  password: "",
  username: "",
})

const register = async () => {
  const resp = await registerPost(registerForm)
  if (resp.code === 0) {

  }
}


</script>

<style scoped>

</style>