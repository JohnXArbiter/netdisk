<template>
    <el-row class="tac">
        <el-menu
                active-text-color="orange"
                background-color="wheat"
                class="el-menu-vertical-demo"
                default-active="2"
                text-color="black"
                @open="handleOpen"
                @close="handleClose"
                style="width: 100%"
        >
            <el-sub-menu index="1">
                <template #title>
                    <el-icon>
                        <Folder/>
                    </el-icon>

                    <span>我的文件</span>
                </template>
                <el-menu-item-group>
                    <el-menu-item index="1-1">
                        <router-link to="/file/0/0">
                            <el-icon>
                                <Document/>
                            </el-icon>
                            文档
                        </router-link>
                    </el-menu-item>
                    <el-menu-item index="1-2">
                        <router-link to="/file/1/0">
                            <el-icon>
                                <PictureFilled/>
                            </el-icon>
                            图片
                        </router-link>
                    </el-menu-item>
                    <el-menu-item index="1-3">
                        <router-link to="/file/2/0">
                            <el-icon>
                                <VideoCameraFilled/>
                            </el-icon>
                            视频
                        </router-link>
                    </el-menu-item>
                    <el-menu-item index="1-4">
                        <router-link to="/file/3/0">
                            <el-icon>
                                <Headset/>
                            </el-icon>
                            音频
                        </router-link>
                    </el-menu-item>
                    <el-menu-item index="1-5">
                        <router-link to="/file/4/0">
                            <el-icon>
                                <More/>
                            </el-icon>
                            其他
                        </router-link>
                    </el-menu-item>
                </el-menu-item-group>
            </el-sub-menu>
            <el-menu-item index="2">
                <el-icon>
                    <Delete/>
                </el-icon>
                <span>回收站</span>
            </el-menu-item>
            <el-menu-item index="3" disabled>
                <el-icon>
                    <document/>
                </el-icon>
                <span>Navigator Three</span>
            </el-menu-item>
        </el-menu>
        <div class="cat-banner">
            <img src="@/assets/cat_scare.png" alt="" style="width: 30%"/>
        </div>
    </el-row>
</template>

<script lang="ts" setup>
import {
    Folder, Document, PictureFilled, Headset, More,
    Delete, VideoCameraFilled
} from '@element-plus/icons-vue'

const handleOpen = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}
</script>
<style>
.tac {
    overflow: hidden;
    height: 94vh;
}

.cat-banner {
    background-color: wheat;
    position: absolute;
    bottom: 1px;
}

.cat-banner img {
    animation-duration: 1s;
    animation-name: slideInRight-enter;
    animation-iteration-count: infinite;
}

@keyframes slideInRight-enter {
    from {
        margin-left: 0;
    }
    25% {
        width: 50%;
    }
    to {
        margin-left: 100%;
    }
}
</style>