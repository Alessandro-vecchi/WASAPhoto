<script>
import Post from '@/components/Post.vue'
// import NavBar from '@/components/NavBar.vue'

export default {
	name: 'Home',
	components: {
		Post,
	}	
}
</script>

<template>
	<div class="Home">
		<div class="timeline">
			<Post />
			<Post />
			<Post />
		</div>
		<div class="sidebar">
			<side-bar />
		</div>
	</div>
</template>

<style>
.Home {
	max-width: 605px;
	margin-left: auto;
	margin-right: auto;
	position: relative;
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
		top: calc(var(--header-height) + 30px + 18px);
	}
}
</style>
