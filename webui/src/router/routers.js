import { createRouter, createWebHashHistory } from 'vue-router'
import PostForm from '../components/PostForm.vue'
import Login from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import list from '../components/UserList.vue'
import comments from '../components/Comments.vue'
import search from '../components/GalleryItem.vue'
import Home from '../views/HomeView.vue'
import Edit from '../components/EditPage.vue'
import changed from '../components/ChangeUsername.vue'
import NavBar from '../components/NavBar.vue'
import Post from '@/components/Post.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{ path: '/', component: Login, name: 'Login', alias: "/login" },
		{ path: '/users/:user_id/stream/', component: Home, name: 'Home'},
		// get user profile by name
		{ path: '/users/', component: ProfileView, name: 'Profile', query: { username: { type: String, default: '' } }},
		// change username
		{ path: '/users/:user_id/changeUsername/', component: changed, name: 'username'},
		// change profile
		{ path: '/users/:user_id/editProfile/', component: Edit, name: 'EditPage'},
		// delete profile

		/* Photo */
		// upload a photo
		{ path: '/users/:user_id/form/', component: PostForm, name: 'PostForm'},

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

		{ path: '/list', component: list, name: 'list' },
		{ path: '/comments', component: comments, name: 'comments' },
	], sensitive: true //, strict: true
})

export default router
