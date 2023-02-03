<script>
import Post from '@/components/Post.vue'
import NavBar from '@/components/NavBar.vue'
import { eventBus } from "@/main.js"

export default {
	name: 'Home',
	components: {
		Post,
		NavBar
	},
	data: function () {
		return {
			errormsg: null,
			loading: false,
            header: localStorage.getItem('Authorization'),
			stream: [],
		}
	},
	methods: {
		async GetStream() {
			this.loading = true;
			this.errormsg = null;

            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
			try {
				let response = await this.$axios.get("/users/" + this.header + "/stream/");
				this.stream = response.data;
				console.log(this.stream);
			} catch (e) {
                this.errormsg = e.response.data.error.toString();
			}  
			this.loading = false;
		},
		async refresh() {
			await this.GetStream();
			// this.loading = false
			console.log("stream found")
			eventBus.user_id = this.$route.params.user_id
		}

	},

	mounted() {
		this.refresh()
		console.log("stream ended")
	}
}
</script>

<template>

	<div class="Home">
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<div v-if="stream" class="timeline">
			<Post v-on:refresh-parent="refresh" v-for="post in stream" :key="post.photoId" 
			:photoId="post.photoId" :owner="post.username" :profilePictureUrl="post.profile_pic" :image="post.image"
			:timestamp="post.timestamp" :caption="post.caption" :likesCount="post.likes_count" :commentsCount="post.comments_count"/>
		</div>
		<div class="sidebar">
			<NavBar />
		</div>
	</div>
</template>

<style>
.Home {
	max-width: 601px;
	margin-left: auto;
	margin-right: auto;
	padding-bottom: 140px;
}

.sidebar {
	display: contents;
}
</style>
