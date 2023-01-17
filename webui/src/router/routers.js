import { createRouter, createWebHashHistory } from 'vue-router'
import PostForm from '../components/PostForm.vue'
import Login from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import list from '../components/UserList.vue'
import comments from '../components/Comments.vue'
/* import search from '../components/GalleryItem.vue' */
import Home from '../views/HomeView.vue'
import prof from '../views/placeholder.vue'
import Edit from '../components/EditPage.vue'
import changed from '../components/ChangeUsername.vue'
/* import Post from '@/components/Post.vue' */

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
		{ path: '/new/', component: prof},

		/* LIKES */
		// Get list of the users that added a like
		{ path: '/:listType(listUsers|likes|followers|following)/', component: list},

		/* FOLLOW */
		// Get list of the followers
		//{ path: '/users/:user_id/listUsers/', component: List, name: 'list', alias:"/photos/:photo_id/listFollowers/" },

		// Get list of the following

		/* BAN */
		// Get list of the banned users
		{ path: '/comments', component: comments, name: 'comments' },
	], sensitive: true //, strict: true
})

export default router
