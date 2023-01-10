<script>
import Post from '@/components/Post.vue'
import NavBar from '@/components/NavBar.vue'

export default {
	name: 'Home',
	components: {
		Post,
		NavBar
	}
}
</script>

<template>
	<NavBar />
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
