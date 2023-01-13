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
			//username: "",
			//stream: [],
			//logged: localStorage.getItem('Authorization')
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			/* try {
				let response = await this.$axios.get("/users/" + logged + "/stream/");
				this.stream = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}  */
			this.loading = false;
		},

	},
	beforeMount(){
    //console.log(this.user_id)
},

	mounted() {
		this.refresh()
	}
}
</script>

<template>

	<NavBar :user_id="user_id" />
	<div class="Home">
		<div class="timeline">
			<Post />
			<Post />
			<Post />
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
	display: none;
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
