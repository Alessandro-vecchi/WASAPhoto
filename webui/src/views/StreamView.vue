<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
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
	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
  <nav>
    
    <div id="nav-home-section" class="nav-section">
      <router-link to="/stream/"><font-awesome-icon icon="fa-solid fa-house" inverse/></router-link>
      <router-link to="/stream/">Home</router-link>
    </div>
    <div id="nav-search-section" class="nav-section">
      <router-link to="#" role="link"><font-awesome-icon icon="fa-solid fa-magnifying-glass" inverse/></router-link>
    </div>
    <div id="nav-profile-section" class="nav-section">
      <!-- profile picture-->
    </div>
  </nav>
</template>

<style>
:root {
    --background-color: rgb(6, 12, 24);
    --border-color: rgba(255, 255, 255, 0.2);
  }
  
  a {
    font-family: "Rubik", sans-serif;
    font-weight: 400;
    color: white;
    text-decoration: none;
  }
  
  nav {
    background-color: var(--background-color);
    display: flex;
    flex-direction: row;
    width: 50vh;
    border: 1px solid var(--border-color);
    overflow: hidden;
    position: fixed;
    bottom: 0;
    left: 33%;
  }
  
  nav > .nav-section {
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
  
  font-awesome-icon {
    text-decoration: none;
    color: white;
  }
</style>