import {createRouter, createWebHashHistory} from 'vue-router'

const routes = [
    {
        path: '/login',
        name: 'login',
        component: () => import('../views/Login.vue'),
    },
    {
        path: '/',
        redirect: '/folder/0',
        component: () => import('../views/Main.vue'),
        children: [
            {
                path: 'folder/:id',
                component: () => import('../views/fileFolder/FileFolder.vue'),
            },
        ]
    },
]

const router = createRouter({history: createWebHashHistory(), routes})

export default router