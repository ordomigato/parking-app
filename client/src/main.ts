import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import ButtonComponent from './components/global/ButtonComponent.vue'
import TextInput from './components/global/TextInput.vue'
import ErrorDisplay from './components/global/ErrorDisplay.vue'

const app = createApp(App)

app.component('c-button', ButtonComponent)
app.component('text-input', TextInput)
app.component('error-display', ErrorDisplay)

app.use(createPinia())
app.use(router)

app.mount('#app')
