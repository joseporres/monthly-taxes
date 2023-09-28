<script setup>

import { defineProps, defineEmits, ref, inject } from 'vue'
import { login as AuthLogin } from '../api/index'
import { useRouter } from 'vue-router';

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
    const { token } = response.data
    localStorage.setItem('token', token)
    state.isLoggIn = true;
    router.push('/')

  } catch (error) {
    console.log(error)
  }
  // let response = await AuthLogin(payload)      
}

</script>

<template>
  <input type="text" v-model="email" />
  <input type="password" v-model="password" />
  <button @click="login">Login</button>
</template>

<style scoped>
.read-the-docs {
  color: #888;
}
</style>
