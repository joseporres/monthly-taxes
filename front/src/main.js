import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import Home from './components/Home.vue'
import Login from './components/Login.vue'
import { globalState } from './store'
// import VueRouter from 'vue-router'
import { createRouter, createWebHistory, RouterView } from 'vue-router'

const routes = [
    {
        path: '/', component: Home,
        meta: {
            requiresAuth: true
        }

    },
    {
        path: '/login', component: Login,
        meta: {
            requiresAuth: false
        }
    },
]

const history = createWebHistory()

const router = createRouter({
    history: history,
    routes
})

router.beforeEach(async (to, form, next) => {
    if (to.matched.some(record => record.meta.requiresAuth)) {
        if (globalState.isLoggIn) {
            next()
            return
        }
        next('/login')
    } else {
        next()
    }
})


const app = createApp(App)
app.provide('state', globalState)
app.use(router);
app.mount('#app');
