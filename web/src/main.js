import './assets/scss/main.scss'
import 'element-plus/theme-chalk/display.css'
import 'virtual:uno.css'
import 'virtual:unocss-devtools'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'


import App from './App.vue'
import router from './router'
const app = createApp(App)
const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)
app.use(pinia)
app.use(router)


app.mount('#app')
