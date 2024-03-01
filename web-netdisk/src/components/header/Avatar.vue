<template>
    <el-row class="el-row">
        <el-col :span="23">
            <div class="grid-content ep-bg-purple"></div>
        </el-col>
        <el-col :span="1">
            <div class="grid-content ep-bg-purple-light">
                <div class="demo-basic--circle">
                    <el-popover
                            placement="top-start"
                            title="Title"
                            :width="200"
                            trigger="hover"
                            content="this is content, this is content, this is content"
                    >
                        <template #reference>
                            <div class="block head-pic">
                                <el-avatar :size="50" :src="url"/>
                            </div>
                        </template>
                    </el-popover>

                </div>
            </div>
        </el-col>
    </el-row>
</template>

<script lang="ts" setup>
import {onMounted} from "vue";
import {useBaseStore, UserInfo} from "../../store";

const url = "https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"

let user: { data: UserInfo }

async function showUserInfo() {
    if (!user) {
        user = await useBaseStore().getUserInfo()
    }
}

onMounted(() => {
    showUserInfo()
})

</script>

<style scoped>
.head-pic {
    margin-top: 5px;
}

.el-row {
    margin-bottom: 20px;
}

.el-row:last-child {
    margin-bottom: 0;
}

.grid-content {
    min-height: 36px;
}
</style>
