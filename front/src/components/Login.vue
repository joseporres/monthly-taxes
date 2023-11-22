<script setup>

import {  ref, inject } from 'vue'
import { login as AuthLogin } from '../api/index'
import { useRouter } from 'vue-router';
import { toast } from 'vue3-toastify';


const email = ref('')
const password = ref('')
const state = inject('state')
const router = useRouter();

async function login() {
  let payload = {
    email: email.value,
    password: password.value
  }
  try {
    let response = await AuthLogin(payload)
    toast.success('Bienvenido')
    // console.log(response)
    const { token } = response.data;
    localStorage.setItem('token', token)
    router.push('/')
  } catch ({ response }) {
    // console.log(response)
    let errorMessage = response.data.message || 'Ocurrio un error'
    toast.error(errorMessage)
  }
  // let response = await AuthLogin(payload)      
}

</script>

<template>

  <div class="container" style="height: 100vh;">
    <h1 class="title-default">
      Iniciar Sesión
    </h1>
    <div class="form-group">
      <input type="email" v-model="email"
       placeholder="Correo"
      />
      <input type="password" 
      v-model="password" 
      placeholder="Contraseña"
      @keyup.enter="login"
      />
      <button class="btn btn-default" @click="login">Enviar</button>
      <router-link to="/register">Registrarse</router-link>
    </div>
  </div>
</template>

<style scoped>

</style>
