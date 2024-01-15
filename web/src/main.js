import './assets/scss/main.scss'
import 'element-plus/dist/index.css'
import 'virtual:uno.css'
// import 'virtual:unocss-devtools'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

import App from './App.vue'

const app = createApp(App)
const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)
app.use(pinia)

import router from './router'
app.use(router)

app.mount('#app')
