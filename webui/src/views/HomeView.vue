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
			itemsPerPage: 5,
			currentPage: 1
		}
	},
	computed: {
		offset() {
			return (this.currentPage - 1) * this.itemsPerPage
		},
	},
	methods: {
		async getInitialStream() {
			this.loading = true;
			this.errormsg = null;

			this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
				error => { return Promise.reject(error); });
			this.$axios.get("/users/" + this.header + "/stream/", { params: { offset: 0 } }).then(response => {
				this.stream = response.data, console.log(this.stream, typeof (this.stream));
			}).catch(e => this.errormsg = e.response.data.error.toString());
			this.loading = false;
		},
		async getNextStream() {
			this.errormsg = null;
			window.onscroll = () => {
				console.log(document.documentElement.scrollTop, window.innerHeight, document.documentElement.offsetHeight)
				let bottomOfWindow = document.documentElement.scrollTop + window.innerHeight === document.documentElement.offsetHeight;
				if (bottomOfWindow) {
					this.currentPage++
					this.$axios.get("/users/" + this.header + "/stream/", { params: { offset: this.offset } }).then(response => {
						if (response) {
							this.stream = [...this.stream, ...response.data]
						};
					}).catch(e => {
						if (e.response) {
							this.errormsg = e.response.data.error.toString()
						}
					});
				}
			}
		},
		async refresh() {
			await this.getNextStream();
			// this.loading = false
			console.log("stream found")
			eventBus.user_id = this.$route.params.user_id
		},
		/* async updateLikesCommentsPhoto(photoId) {
			console.log("ID", photoId);
			this.loading = true;
			this.errormsg = null;

			this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = this.header; return config; },
				error => { return Promise.reject(error); });
			try {
				let response = await this.$axios.get("/photos/" + photoId);
				this.post = response.data;
				eventBus.getPhotoId = this.post.photoId
				console.log(this.post);
			} catch (e) {
				this.errormsg = e.response.data.error.toString();
			}  
			this.loading = false;

			// Emit an event to the child component with only this.post.likeCount
				this.$emit('refresh-child', this.post.likes_count);
		}, */

	},
	beforeMount() {
		this.getInitialStream();
	},

	mounted() {
		this.getNextStream()
		console.log("stream ended")
	}
}
</script>

<template>

	<div class="Home">
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<div class="timeline">
			<Post v-on:refresh-parent="refresh" v-for="post in stream" :key="post.photoId" :photoId="post.photoId"
				:owner="post.username" :profilePictureUrl="post.profile_pic" :image="post.image"
				:timestamp="post.timestamp" :caption="post.caption" :likesCount="post.likes_count"
				:commentsCount="post.comments_count" />
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
