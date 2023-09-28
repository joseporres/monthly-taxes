import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
import { globalState } from './store'
import 'vue3-toastify/dist/index.css';


const app = createApp(App)
app.provide('state', globalState)
app.use(router);
app.mount('#app');
