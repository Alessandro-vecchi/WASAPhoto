<script>
export default {

	data: function () {
		return {
			errormsg: null,
			loading: false,
			username:"",
			stream: [],
		}
	},
	methods: {
		load() {
			return load
		},
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/stream/");
				this.stream = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async get_user_profile() {
            this.loading = true;
            this.errormsg = null;
            try {
                this.$axios.get("/users/?username="+this.$route.params.username).then(response => (this.profile = response.data));
                this.$router.push({ name: "profile" })
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
		rect() {
			/*const s = document.getElementById('nav-search-section');
			s.addEventListener('click', toggleActive);
			function toggleActive() {*/
			const rectangle = document.querySelector('.rectangle');
			rectangle.classList.toggle('active')

		},
		restart_search() {
			this.$route.params.username = ""
			onclick="document.getElementById('search').value = ''"
		}

	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<nav>
		<div class="nav-wrapper">

			<div id="nav-home-section" class="nav-section">
				<router-link to="/stream/"><font-awesome-icon class="icons" icon="fa-solid fa-house" size="xl"
						inverse /></router-link>
				<router-link to="/stream/" class="font-style">Home</router-link>
			</div>
			<div id="nav-search-section" class="nav-section" @click="rect">
				<font-awesome-icon class="icons" icon="fa-solid fa-magnifying-glass" size="xl" inverse />
			</div>
			<div id="nav-profile-section" class="nav-section">
				<!-- profile picture-->
				<router-link to="/users/:username" role="link"><font-awesome-icon class="icons"
						icon="fa-solid fa-user" size="xl" inverse /></router-link>
			</div>
		</div>
		<div class="rectangle">
			<p class="title">Search</p>
			<div class="input-wrapper">
				<font-awesome-icon class="icons" icon="fa-solid fa-magnifying-glass" inverse @click="get_user_profile"/>
				<input id="search" v-model="username" type="text" placeholder="Search" />
				<font-awesome-icon class="icons" icon="fa-solid fa-xmark" onclick="document.getElementById('search').value = ''"/>
			</div>
		</div>
	</nav>
</template>

<style>
:root {
	--background-color: rgb(6, 12, 24);
	--border-color: rgba(255, 255, 255, 0.2);
}

.font-style {
	font-size: 1.5em;
	font-family: "Rubik", sans-serif;
	font-weight: 400;
	color: white;
	text-decoration: none;
}

.nav-wrapper {
	background-color: var(--background-color);
	display: flex;
	flex-direction: row;
	width: 25vw;
	border: 1px solid var(--border-color);
	overflow: hidden;
	position: fixed;
	bottom: 0;
	left: 33%;
}

.nav-wrapper>.nav-section {
	padding: 3rem 2rem;
	display: flex;
	gap: 1rem;
	border-left: 1px solid var(--border-color);
	justify-content: center;
}

#nav-home-section,
#nav-search-section,
#nav-profile-section {
	flex-basis: calc(100% / 3);
}

.icons {
	text-decoration: none;
	color: white;
	cursor: pointer;
}

.rectangle {
	background-color: rgb(10, 12, 24);
	width: 25vw;
	height: 17vh;
	position: absolute;
	bottom: -5vh;
	left: 33%;
	border-radius: 10px 10px 0 0;
	transform: translateY(0px);
	transition: transform 1.5s ease-in;
	z-index: -1;
}


.rectangle.active {
	transform: translateY(-170px);
	transition: transform 1.5s ease-out;
}

.input-wrapper {
	display: flex;
	align-items: center;
	justify-content: space-between;
	background: rgb(100, 100, 100);
	width: 75%;
	height: 40px;
	padding: 0.5rem;
	border-radius: 0.5rem;
	color: white;
	box-shadow: 0.25rem 0.25rem 0rem rgb(189, 189, 189);
	position: absolute;
	left: 6%;
	top: 50%;

}
#search {
	margin: 0 0.5rem 0 0.5rem;
	width: 100%;
	height: 20px;
	padding: 0.5rem;
	border: none;
	outline: none;
	background: rgb(34, 34, 34);
	color: white;
	font-size: 1.5rem;
}

.title {
	font-size: 1.8em;
	font-family: "Rubik", sans-serif;
	font-weight: 400;
	color: white;
	position: absolute;
	top: -10px;
	left: 8%;

}
</style>