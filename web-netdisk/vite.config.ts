import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import {resolve} from "path"
import * as path from "path";

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue()],
    server: {
        port: 5173,
        proxy: {
            '/api': {
                target: 'http://127.0.0.1:8888/',
                changeOrigin: true,
                rewrite: (path) => path.replace(/^\/api/, '') // 不可以省略rewrite
            }
        }
    },
    resolve: {
        alias: [
            {
               '@': path.resolve('./src'),
            },
            {
                find: '@',
                replacement: resolve(__dirname, "src")
            },
            {
                find: 'components',
                replacement: resolve(__dirname, "src/components")
            },
            {
                find: 'views',
                replacement: resolve(__dirname, "src/views")
            }
        ]
    }
})

