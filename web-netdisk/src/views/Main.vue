<template>
  <div class="common-layout">
    <el-container>
      <el-header style="background-color: lightsteelblue">
        <avatar/>

      </el-header>
      <el-container>
        <el-aside width="200px" style="background-color: #16e52d">
          <navigator/>
        </el-aside>
        <el-main style="background-color: #d02e59">
          <router-view/>
        </el-main>
      </el-container>
    </el-container>
  </div>

</template>

<script lang="ts" setup>
import Avatar from "../components/header/Avatar.vue";
import Navigator from "../components/aside/Navigator.vue";

</script>

<style scoped>

</style>
