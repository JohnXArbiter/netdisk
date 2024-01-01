import {createApp} from 'vue'
import ElementPlus from 'element-plus'
import router from "./router"
import 'element-plus/dist/index.css'
import {createPinia} from "pinia";
import App from './App.vue'


const app = createApp(App);
app.use(createPinia()).use(ElementPlus).use(router).mount('#app');
