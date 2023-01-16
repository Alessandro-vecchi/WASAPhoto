<script>
import Post from '@/components/Post.vue'
import NavBar from '@/components/NavBar.vue'

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
			user_id: this.$route.params.user_id,
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
				this.errormsg = e.toString();
			}  
			this.loading = false;
		},
		refresh() {
			this.GetStream();
		}

	},


	mounted() {
		this.refresh()
	}
}
</script>

<template>

	<div class="Home">
		<div class="timeline">
			<Post v-on:refresh-parent="refresh" v-for="obj in stream" :key="obj.photo_id" :post="obj"/>
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
	padding-bottom: 10vh;
}
@media (--t) {
	.Home {
		max-width: none;
		display: grid;
		grid-template-columns: 1fr 295px;
		grid-gap: 28px;
	}
}
.sidebar {
	display: contents;
}
@media (--t) {
	.sidebar {
		display: block;
		margin-top: 16px;
	}
	.sidebar p {
		position: sticky;
		top: calc(53px + 30px + 18px);
	}
}
</style>
