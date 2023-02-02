<script>
import Post from '@/components/Post.vue'
import NavBar from '@/components/NavBar.vue'
import { eventBus } from "@/main.js"

export default {
	components: {
		Post,
		NavBar
	},
	data: function () {
		return {
			errormsg: null,
			loading: false,
            header: localStorage.getItem('Authorization'),
			post: "",
            photoId: eventBus.getPhotoId,
		}
	},
	methods: {
		async GetSinglePhoto() {
            console.log("ID", this.photoId);
			this.loading = true;
			this.errormsg = null;

            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = this.header; return config; },
                error => { return Promise.reject(error); });
			try {
				let response = await this.$axios.get("/photos/" + this.photoId);
				this.post = response.data;
				eventBus.getPhotoId = this.photoId
				console.log(this.post);
			} catch (e) {
				this.errormsg = e.toString();
			}  
			this.loading = false;
		},
		async refresh() {
			await this.GetSinglePhoto();
		}

	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
    <div class="photo">
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <Post v-on:refresh-parent="refresh" v-if="post"
            :photoId="post.photoId" :owner="post.username" :profilePictureUrl="post.profile_pic" :image="post.image"
            :timestamp="post.timestamp" :caption="post.caption" :likesCount="post.likes_count"
            :commentsCount="post.comments_count" />
        <div class="navbar">
            <NavBar />
        </div>
    </div>
</template>

<style scoped>
.photo {
	max-width: 601px;
	margin: auto;
	/* padding-bottom: 10vh; */
}

.navbar {
	display: contents;
}
</style>