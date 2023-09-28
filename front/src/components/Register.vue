<script setup>
import { ref } from 'vue'
import { register as AuthRegister } from '../api/index'
import { useRouter } from 'vue-router';
import { toast } from 'vue3-toastify';

const email = ref('')
const password = ref('')
const name = ref('')
const router = useRouter();

async function signUp() {
    let payload = {
        email: email.value,
        password: password.value,
        name: name.value
    }
    try {
        await AuthRegister(payload)
        toast.success('Usuario registrado')
        router.push('/login')
    } catch (error) {
        let errorMessage = error.response.data.message || 'Ocurrio un error'
        toast.error(errorMessage)
        // console.log(error)
    }
    // let response = await AuthLogin(payload)      
}

</script>

<template>

    <div class="container" style="height: 100vh;">
        <h1 class="title-sucess">Registrarse</h1>
        <div class="form-group">
            <input type="text" v-model="email" 
             placeholder="Correo"
            />
            <input type="text" v-model="name" 
            placeholder="Nombre"
            />
            <input type="password" v-model="password"
            placeholder="Contraseña"
            />
            <button class="btn btn-sucess" @click="signUp">Enviar</button>
            <router-link to="/login">Iniciar Sessión</router-link>
        </div>
    </div>
</template>

<style scoped>

</style>
