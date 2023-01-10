import {createRouter, createWebHashHistory} from 'vue-router'
import Stream from '../views/StreamView.vue'
import Login from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import list from '../components/UserList.vue'
import search from '../components/GalleryItem.vue'
import Home from '../views/HomeView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/stream/', component: Home, name: 'Home'},
		{path: '/', component: Login, name: 'login'},
		{path: '/search', component: search, name: 'search'},

		{path: '/list', component: list, name: 'list'},
		{path: '/users/:username', component: ProfileView, name: 'profile', props: true},
	]
})

export default router
