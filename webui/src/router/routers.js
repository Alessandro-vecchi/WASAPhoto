import {createRouter, createWebHashHistory} from 'vue-router'
import Stream from '../views/StreamView.vue'
import Login from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import list from '../components/UserList.vue'
import comments from '../components/Comments.vue'
import search from '../components/GalleryItem.vue'
import Home from '../views/HomeView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: Login, name: 'login'},
		{path: '/users/:user_id/stream/', component: Home, name: 'Home', props: true},
		// get user profile by name
		{path: '/users/', component: ProfileView, name: 'profile', query: {username: {type: String, default:''}}},
		// change username

		// change profile

		// delete profile

		/* Photo */
		// upload a photo

		// delete a photo

		/* COMMENTS */
		// See comments
		
		// Upload a comment
		
		// Modify a comment

		// Delete a comment

		/* LIKES */
		// Get list of the users that added a like

		// Like a photo

		// Unlike a photo

		/* FOLLOW */
		// Get list of the followers

		// Get list of the following

		// Follow a user

		// Unfollow a user

		/* BAN */
		// Get list of the banned users

		// Ban user

		// Unban user
		{path: '/search', component: search, name: 'search'},

		{path: '/list', component: list, name: 'list'},
		{path: '/comments', component: comments, name: 'comments'},
	]
})

export default router
