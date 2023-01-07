import {createRouter, createWebHashHistory} from 'vue-router'
import Stream from '../views/StreamView.vue'
import Login from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/stream/', component: Stream, name: 'stream'},
		{path: '/', component: Login, name: 'login'},
		{path: '/users/:username', component: ProfileView, name: 'profile', props: true},
	]
})

export default router
