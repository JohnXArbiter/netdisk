import {createRouter, createWebHashHistory, createWebHistory} from 'vue-router'

const routes = [
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
    {
        path: '/netdisk',
        name: '',
        component: () => import('../views/NoAuth.vue'),
        children: [
            {
                path: '/login',
                component: () => import('../components/LoginForm.vue')
            },
            {
                path: '/register',
                component: () => import('../components/RegisterForm.vue')
            }
        ]
    },
]

const router = createRouter({history: createWebHashHistory(), routes})

export default router