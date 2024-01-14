import {createRouter, createWebHistory} from 'vue-router'

const routes = [
    {
        path: '/login',
        name: 'login',
        component: () => import('../views/Login.vue')
    },
    {
        path: '/',
        redirect: 'folder/0',
        component: () => import('../views/Main.vue'),
        children: [
            {
                path: 'folder/:folderId',
                component: () => import('../views/fileFolder/FileFolder.vue'),
                props: true
            },
            {
                path: 'file/:fileType',
                component: () => import('../views/file/File.vue'),
                props: true
            },
            {
                path: 'bin',
                component: () => import('../views/bin/Bin.vue')
            }
        ]
    },
]

const router = createRouter({history: createWebHistory(), routes})

export default router