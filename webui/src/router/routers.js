import {createRouter, createWebHashHistory} from 'vue-router'
import Stream from '../views/StreamView.vue'
import Login from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import nav from '../components/NavBar.vue'
import search from '../components/SearchBar.vue'
import Home from '../views/HomeView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/stream/', component: nav, name: 'stream'},
		{path: '/', component: Login, name: 'login'},
		{path: '/search', component: search, name: 'search'},

		{path: '/post', component: Home, name: 'home'},
		{path: '/users/:username', component: ProfileView, name: 'profile', props: true},
	]
})

export default router
