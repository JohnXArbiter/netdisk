<template>
    <el-row class="tac">
        <el-menu
                active-text-color="orange"
                background-color="rgba(245,222,179,1)"
                class="el-menu-vertical-demo"
                :default-openeds="['1']"
                text-color="black"
                :collapse="collapse"
                style="width: 100%;"
        >
            <el-sub-menu index="1">
                <template #title>
                    <el-icon>
                        <Files/>
                    </el-icon>
                    文件
                </template>
                <el-menu-item index="1-1">
                    <router-link class="a" to="/file/folder/0">
                        <el-icon>
                            <Folder/>
                        </el-icon>
                        全部文件
                    </router-link>
                </el-menu-item>
                <el-menu-item index="1-2">
                    <router-link class="a" to="/file/0">
                        <el-icon>
                            <Document/>
                        </el-icon>
                        文档
                    </router-link>
                </el-menu-item>
                <el-menu-item index="1-3">
                    <router-link class="a" to="/file/1">
                        <el-icon>
                            <PictureFilled/>
                        </el-icon>
                        图片
                    </router-link>
                </el-menu-item>
                <el-menu-item index="1-4">
                    <router-link class="a" to="/file/2">
                        <el-icon>
                            <VideoCameraFilled/>
                        </el-icon>
                        视频
                    </router-link>
                </el-menu-item>
                <el-menu-item index="1-5">
                    <router-link class="a" to="/file/3">
                        <el-icon>
                            <Headset/>
                        </el-icon>
                        音频
                    </router-link>
                </el-menu-item>
                <el-menu-item index="1-6">
                    <router-link class="a" to="/file/4">
                        <el-icon>
                            <More/>
                        </el-icon>
                        其他
                    </router-link>
                </el-menu-item>
            </el-sub-menu>
            <el-menu-item index="2">
                <a class="a" href="/bin">
                    <el-icon>
                        <Delete/>
                    </el-icon>
                    回收站
                </a>
            </el-menu-item>
            <el-menu-item index="3">
                <a class="a" href="/share">
                    <el-icon>
                        <Share/>
                    </el-icon>
                    分享
                </a>
            </el-menu-item>
        </el-menu>
        <div class="cat-banner">
            <img src="@/assets/cat_scare.png" alt="" style="width: 30%"/>
        </div>
    </el-row>
</template>

<script lang="ts" setup>
import {
    Files, Share,
    Folder, Document, PictureFilled, Headset, More,
    Delete, VideoCameraFilled
} from '@element-plus/icons-vue'
import {ref} from "vue";

let collapse = ref(false)

</script>
<style>
.tac {
    overflow: hidden;
    height: 94vh;
}

.cat-banner {
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

.a {
    text-decoration: none;
    color: #606060;
}

router-link:hover {
    text-decoration: none;
}
</style>