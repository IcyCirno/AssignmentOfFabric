import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './main.css'
import 'font-awesome/css/font-awesome.min.css'
import '@fortawesome/fontawesome-free/css/all.css'



const app = createApp(App)
app.config.globalProperties.$bus = app
app.use(router)
app.mount('#app')