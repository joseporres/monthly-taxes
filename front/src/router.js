import { createRouter, createWebHistory } from 'vue-router';
import Home from './components/Home.vue'
import Login from './components/Login.vue'
import Register from './components/Register.vue'

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
    {
        path: '/register',
        component: Register,
        meta: {
            requiresAuth: false
        }
    }
]

const history = createWebHistory()

const router = createRouter({
    history: history,
    routes
})

router.beforeEach(async (to, form, next) => {
    if (to.matched.some(record => record.meta.requiresAuth)) {
        if (localStorage.getItem('token')) {
            next()
            return
        }
        next('/login')
    } else {
        next()
    }
})


export default router;