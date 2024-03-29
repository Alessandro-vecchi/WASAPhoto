import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router/routers.js'
import axios from './services/axios.js'
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import WelcomeMsg from './components/WelcomeMsg.vue'

import './assets/main.css'

/* import the fontawesome core */
import { library } from '@fortawesome/fontawesome-svg-core'

/* import font awesome icon component */
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

/* import specific icons */
import {faUser, faMagnifyingGlass, faHouse, faXmark, faEllipsis, faClone, faHeart, faComment, faPlusCircle, faPlus, faCheck, faBan, faImage, faChevronLeft as faCL, faReply} from '@fortawesome/free-solid-svg-icons'
import {faHeart as fa1, faComment as fa2, faImage as fai} from '@fortawesome/free-regular-svg-icons'
/* add icons to the library */
library.add(faUser, faMagnifyingGlass, faHouse, faXmark, faHeart, fa1, faComment, fa2, faEllipsis, faClone, faPlusCircle, faPlus, faCheck, faBan, faImage, fai, faCL, faReply)

const eventBus = new reactive({});
export { eventBus }

const app = createApp(App)

app.config.globalProperties.$axios = axios;
app.component('font-awesome-icon', FontAwesomeIcon)
app.component("ErrorMsg", ErrorMsg);
app.component("WelcomeMsg", WelcomeMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.use(router)
app.mount('#app')
