<template>
    <div class="common-layout">
        <el-container style="box-shadow: 0 2px 6px 0 #ddd;">
            <el-header style="background-color: lightsalmon;">
                <avatar/>
            </el-header>

            <el-container>
                <el-aside style="width: 10%;">
                    <navigator/>
                </el-aside>

                <el-main style="background-color: #fff2e3; padding: 15px 0 0 20px; position: relative">
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
.common-layout {
    margin: 0;
}

.header {
    top: 0;
    width: 100%;
    position: fixed;
    box-shadow: 0 2px 6px 0 #ddd;
    background-color: #fff;
    z-index: 1000;
}
</style>
