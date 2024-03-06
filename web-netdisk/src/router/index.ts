import {createRouter, createWebHistory} from 'vue-router'

const routes = [
    {
        path: '/login',
        name: 'login',
        component: () => import('@/views/Login.vue')
    },
    {
        path: '/',
        redirect: 'file/folder/0',
        component: () => import('@/views/Main.vue'),
        children: [
            {
                path: 'file',
                component: () => import('@/views/fileFolder/FileFolder.vue'),
                props: true,
                children: [
                    {
                        path: 'folder/:folderId',
                        component: () => import('@/components/files/Folder.vue'),
                        props: true,
                    },
                    {
                        path: ':fileType',
                        component: () => import('@/components/files/File.vue'),
                        props: true,
                    }
                ]
            },
            {
                path: 'bin',
                name: 'bin',
                component: () => import('@/views/bin/Bin.vue')
            },
            {
                path: 'share',
                component: () => import('@/components/files/share/Share.vue'),
                props: true,
            },
            {
                path: 'share-info',
                component: () => import('@/components/files/share/ShareInfo.vue'),
                props: true,
            }
        ]
    },
    {
        path: '/user',
        name: 'user',
        component: () => import('@/views/User.vue'),
        children: [
            {
                path: 'info',
                component: () => import('@/components/user/Info.vue'),
                props: true,
            }
        ]
    }
]

const router = createRouter({history: createWebHistory(), routes})

export default router