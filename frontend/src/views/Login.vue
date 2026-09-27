<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
const router=useRouter(),{t}=useI18n()
const username=ref(''),password=ref(''),busy=ref(false),error=ref('')
async function login(){busy.value=true;error.value='';try{const r=await fetch('/api/auth/login',{method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify({username:username.value,password:password.value})});if(!r.ok)throw new Error(t('login.failed'));const d=await r.json();if(!d.token)throw new Error(t('login.failed'));localStorage.setItem('token',d.token);router.push('/dashboard')}catch(e){error.value=e instanceof Error?e.message:t('login.failed')}finally{busy.value=false}}
</script>
<template><section class="login-card glass"><div class="logo">🦇</div><p class="eyebrow">SECURE ACCESS</p><h1>{{t('login.title')}}</h1><p class="muted">{{t('login.subtitle')}}</p><form @submit.prevent="login"><label>{{t('login.username')}}<input v-model.trim="username" autocomplete="username" required maxlength="64"></label><label>{{t('login.password')}}<input v-model="password" type="password" autocomplete="current-password" required></label><p v-if="error" class="error">{{error}}</p><button class="primary" :disabled="busy">{{busy?t('login.loading'):t('login.submit')}} <span>↗</span></button></form><p class="hint">{{t('login.hint')}}</p></section></template>