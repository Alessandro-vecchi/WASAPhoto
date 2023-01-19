import { createRouter, createWebHashHistory } from 'vue-router'
import PostForm from '../components/PostForm.vue'
import Login from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import list from '../components/UserList.vue'
import comments from '../components/Comments.vue'
import Home from '../views/HomeView.vue'
import Edit from '../components/EditPage.vue'
import changed from '../components/ChangeUsername.vue'
import Post from '@/components/SinglePhoto.vue'

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
		// Get Single Photo
		{ path: '/post/:photo_id', component: Post},

		/* COMMENTS */
		// See comments
		{ path: '/photos/:photo_id/comments/', component: comments},

		/* LIKES */ /* FOLLOW */ /* BAN */
		// Get list of the users that added a like/follow/ban
		{ path: '/:listType(likes|followers|following|bans)/', component: list},

	], sensitive: true //, strict: true
})

export default router
