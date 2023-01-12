import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router/routers'
import axios from './services/axios.js'
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'

import './assets/main.css'

/* import the fontawesome core */
import { library } from '@fortawesome/fontawesome-svg-core'

/* import font awesome icon component */
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

/* import specific icons */
import {faUser, faMagnifyingGlass, faHouse, faXmark, faEllipsis, faClone, faHeart, faComment, faPlusCircle, faPlus, faCheck, faBan} from '@fortawesome/free-solid-svg-icons'
import {faHeart as fa1, faComment as fa2} from '@fortawesome/free-regular-svg-icons'
/* add icons to the library */
library.add(faUser, faMagnifyingGlass, faHouse, faXmark, faHeart, fa1, faComment, fa2, faEllipsis, faClone, faPlusCircle, faPlus, faCheck, faBan)

const app = createApp(App)

app.config.globalProperties.$axios = axios;
app.component('font-awesome-icon', FontAwesomeIcon)
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
//app.component("ErrorMsg", ErrorMsg);
app.use(router)
app.mount('#app')
