import {createRouter, createWebHashHistory} from 'vue-router'
import Stream from '../views/StreamView.vue'
import Login from '../views/LoginView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/stream', component: Stream, name: 'stream'},
		{path: '/', component: Login, name: 'login'},
	]
})

export default router
