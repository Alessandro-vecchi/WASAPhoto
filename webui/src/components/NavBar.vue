<script>
import SearchBar from "@/components/SearchBar.vue"
export default {

    data: function () {
        return {
            errormsg: null,
            loading: false,
            username: "",
            stream: [],
            SearchBar,
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
                this.$axios.get("/users/?username=" + this.$route.params.username).then(response => (this.profile = response.data));
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
                <router-link to="/users/:username" role="link"><font-awesome-icon class="icons" icon="fa-solid fa-user"
                        size="xl" inverse /></router-link>
            </div>
        </div>
        
    </nav>
</template>

<style>
:root {
    --ba1: rgb(6, 12, 24);
    --bo1: rgba(255, 255, 255, 0.2);
}
nav{
	max-width: 604px;
	margin-left: auto;
	margin-right: auto;

}
.font-style {
    font-size: 1.5em;
    font-family: "Rubik", sans-serif;
    font-weight: 400;
    color: white;
    text-decoration: none;
}

.nav-wrapper {
    background-color: var(--ba1);
    border: 2px solid var(--bo1);
    overflow: hidden;
    position: fixed;
	width: 604px;
    height: 12vh;
    bottom: 0;

    display: flex;
    align-items: center;
    z-index: 99;
}

.nav-wrapper>.nav-section {
    padding: 4rem 0;
    display: flex;
    gap: 1rem;
    border-left: 1px solid var(--bo1);
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


</style>