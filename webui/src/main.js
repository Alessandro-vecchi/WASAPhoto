import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router/routers'
import axios from './services/axios.js'

import './assets/main.css'

/* import the fontawesome core */
import { library } from '@fortawesome/fontawesome-svg-core'

/* import font awesome icon component */
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

/* import specific icons */
import { faUser, faMagnifyingGlass, faHouse, faXmark } from '@fortawesome/free-solid-svg-icons'

/* add icons to the library */
library.add(faUser, faMagnifyingGlass, faHouse, faXmark)

const app = createApp(App)

app.config.globalProperties.$axios = axios;
app.component('font-awesome-icon', FontAwesomeIcon)
//app.component("ErrorMsg", ErrorMsg);
app.use(router)
app.mount('#app')
