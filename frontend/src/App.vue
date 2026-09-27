<script setup lang="ts">
import { computed, ref } from 'vue'
import { useI18n } from 'vue-i18n'
const { locale } = useI18n()
const languages = [{id:'fa',name:'فارسی'},{id:'en',name:'English'},{id:'zh',name:'中文'},{id:'ru',name:'Русский'},{id:'ar',name:'العربية'}]
function setLang(id:string){locale.value=id;localStorage.setItem('lang',id);document.documentElement.lang=id;document.documentElement.dir=['fa','ar'].includes(id)?'rtl':'ltr'}
const isRTL=computed(()=>['fa','ar'].includes(locale.value))
</script>
<template><div class="app-shell" :dir="isRTL?'rtl':'ltr'"><div class="ambient one"></div><div class="ambient two"></div><header class="topbar"><a class="brand" href="/"><span class="bat">🦇</span><span>DARK <b>PANEL</b></span></a><select aria-label="Language" :value="locale" @change="setLang(($event.target as HTMLSelectElement).value)"><option v-for="l in languages" :key="l.id" :value="l.id">{{l.name}}</option></select></header><main><router-view /></main><footer> DARK PANEL <span>•</span> VPN management </footer></div></template>