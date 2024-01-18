<template>
  <el-row justify="center">
    <el-col :span="20">
      <el-input placeholder="请输入内容"
                v-model="asd" clearable @clear="getUserList">
        <template #append>
          <el-button @click="getUserList">
            <el-icon>
              <search/>
            </el-icon>
          </el-button>
        </template>
      </el-input>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import {Search} from '@element-plus/icons-vue';
import {ref} from "vue";

let asd = ref('')
</script>

<style scoped>

</style>