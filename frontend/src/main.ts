import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

import App from './App.vue'
import router from './router'
import UUID from 'vue3-uuid';

import '@arco-design/web-vue/dist/arco.css'

const app = createApp(App)

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

app.use(UUID)
app.use(pinia)
app.use(router)

app.mount('#app')
