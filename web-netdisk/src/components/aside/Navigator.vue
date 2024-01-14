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
                        <location/>
                    </el-icon>
                    <span>Navigator One</span>
                </template>
                <el-menu-item-group title="Group One">
                    <el-menu-item index="1-1">item one</el-menu-item>
                    <el-menu-item index="1-2">item two</el-menu-item>
                </el-menu-item-group>
                <el-menu-item-group title="Group Two">
                    <el-menu-item index="1-3">item three</el-menu-item>
                </el-menu-item-group>
                <el-sub-menu index="1-4">
                    <template #title>item four</template>
                    <el-menu-item index="1-4-1">item one</el-menu-item>
                </el-sub-menu>
            </el-sub-menu>
            <el-menu-item index="2">
                <el-icon>
                    <icon-menu/>
                </el-icon>
                <span>Navigator Two</span>
            </el-menu-item>
            <el-menu-item index="3" disabled>
                <el-icon>
                    <document/>
                </el-icon>
                <span>Navigator Three</span>
            </el-menu-item>
            <el-menu-item index="4">
                <el-icon>
                    <setting/>
                </el-icon>
                <span>Navigator Four</span>
            </el-menu-item>
        </el-menu>
        <div class="cat-banner">
            <img src="@/assets/cat_scare.png" alt="" style="width: 30%"/>
        </div>
    </el-row>
</template>

<script lang="ts" setup>
import {Document, Location, Menu as IconMenu, Setting,} from '@element-plus/icons-vue'

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
.cat-banner{
    background-color: wheat;
    position: absolute;
    bottom: 1px;
}
.cat-banner img{
    animation-duration: 1s;
    animation-name: slideInRight-enter;
    animation-iteration-count: infinite;
}
@keyframes slideInRight-enter {
    from{
        margin-left: 0;
    }
    25%{
        width: 50%;
    }
    to{
        margin-left: 100%;
    }
}
</style>