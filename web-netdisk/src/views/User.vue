<template>

    <div class="header">
        <div class="header-content" style="width: 1000px;">
            <router-link to="/file/folder/0" class="logo">
                <el-image class="small-pic"
                          src="/src/assets/img.png"
                          alt="/sec/assets/alt_type1.jpg"
                          :fit="'cover'"/>
            </router-link>
            <!-- 板块信息 -->
            <div class="menu-panel">
                <router-link class="menu-item home" to="/file/folder/0">首页</router-link>
            </div>

            <div class="user-info-panel">
                <avatar></avatar>
            </div>
        </div>
    </div>

    <div class="body-content">
        <router-view></router-view>
    </div>
</template>

<script lang="ts" setup>
import Avatar from "@/components/header/Avatar.vue";


</script>

<style lang="scss" scoped>
.header {
    top: 0;
    width: 100%;
    position: fixed;
    box-shadow: 0 2px 6px 0 #ddd;
    background-color: wheat;
    z-index: 1000;
}

.header-content {
    margin: 0 auto;
    height: 60px;
    display: flex;
    align-items: center;
}

.logo {
    display: block;
    text-decoration: none;
    margin-right: 5px;
}

span {
    font-size: 30px;
}

.menu-item {
    cursor: pointer;
    margin-left: 20px;
}

.menu-panel {
    flex: 1;
}

.home {
    text-decoration: none;
    color: #000;
}

.user-info-panel {
    width: 300px;
}

.body-content {
    padding-top: 60px;
    position: relative;
}

.small-pic {
    width: 50px;
    height: 50px;
    border-radius: 5px;
}
</style>