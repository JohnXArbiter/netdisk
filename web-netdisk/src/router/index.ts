import {createRouter, createWebHistory} from 'vue-router'

const routes = [
    {
        path: '/login',
        name: 'login',
        component: () => import('../views/Login.vue')
    },
    {
        path: '/',
        redirect: '/folder/0',
        component: () => import('../views/Main.vue'),
        children: [
            {
                path: 'folder/:parentFolderId',
                component: () => import('../views/fileFolder/FileFolder.vue'),
                props:true
            },
        ]
    },
]

const router = createRouter({history: createWebHistory(), routes})

export default router