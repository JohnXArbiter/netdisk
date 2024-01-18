<template>
    <el-row justify="center">
        <el-col :span="22">
            <div class="search-box">
                <el-input placeholder="搜索我的文件"
                          v-model="asd" clearable @clear="getUserList">
                    <template #append>
                        <el-button @click="getUserList">
                            <el-icon>
                                <search/>
                            </el-icon>
                        </el-button>
                    </template>
                </el-input>
            </div>

            <div>
                <div style="margin: 3% 0 3% 0; font-size: 0.8rem">文件详情</div>
                <el-image style="border-radius: 5px;width: 100%; height: auto" :src="url" :fit="'contain'"/>
                <div style="margin-top: 4%;padding-left: 4%;">
                    <div class="file-name">{{ exampleFile.name }}</div>
                    <div class="file-info">创建时间 {{exampleFile.created}}</div>
                    <div class="file-info">修改时间 {{exampleFile.updated}}</div>
                    <div class="file-info">文件格式：{{ exampleFile.ext }}</div>
                    <div class="file-info">文件大小：{{ exampleFile.size }}</div>
                </div>
            </div>
        </el-col>
    </el-row>
</template>

<script lang="ts" setup>
import {Search} from '@element-plus/icons-vue';
import {ref} from "vue";

const url = 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
const exampleFile = {
    name: '小猫咪.jpeg',
    created: "2023",
    updated: '2024',
    ext: '.jpeg',
    size: '12kb',
}
let asd = ref('')
</script>

<style scoped>
.search-box {
}

.file-name {
    color: #454d5a;
    font-weight: 600;
}

.file-info {
    margin-top: 3%;
    font-size: 0.8rem;
    color: #878c9c;
}
</style>