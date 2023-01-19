<script>
export default {
    data: function () {
        return {
            logged: localStorage.getItem('Authorization'),
            Username: "",
        }
    },
    methods: {
        async get_user_profile() {
            this.$router.push({ path: "/users/", query: { username: this.Username } })
        },
        rect() {
            const rectangle = document.querySelector('.rectangle');
            rectangle.classList.toggle('active')

        },

    },
}
</script>

<template>
    <nav>
        <div class="nav-wrapper center">

            <div id="nav-home-section" class="nav-section">
                <router-link to="/"><font-awesome-icon class="icons" icon="fa-solid fa-house" size="xl"
                        inverse /></router-link>
                <router-link to="/" class="font-style">Home</router-link>
            </div>
            <div id="nav-search-section" class="nav-section" @click="rect">
                <font-awesome-icon class="icons" icon="fa-solid fa-magnifying-glass" size="xl" inverse />
            </div>
            <div id="nav-profile-section" class="nav-section">
                <!-- profile picture-->
                <font-awesome-icon class="icons" icon="fa-solid fa-user" size="xl" inverse @click="get_user_profile" />
            </div>

        </div>
        <div class="center">
            <div class="rectangle">
                <p class="title font-style">Search</p>
                <div class="input-wrapper">
                    <font-awesome-icon class="icons" icon="fa-solid fa-magnifying-glass" inverse
                        @click="get_user_profile" />
                    <input id="search" v-model="Username" type="text" placeholder="Search" />
                    <font-awesome-icon class="icons" icon="fa-solid fa-xmark"
                        onclick="document.getElementById('search').value = ''" />
                </div>
            </div>
        </div>

    </nav>
</template>

<style>
:root {
    --ba1: rgb(6, 12, 24);
    --bo1: rgba(255, 255, 255, 0.2);
}

nav {
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
    width: 601px;
    height: 12vh;
    bottom: 0;

    z-index: 99;
}

.center {
    display: flex;
    align-items: center;
    justify-content: center;

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

.rectangle {
    position: fixed;
    bottom: 0;
    width: 400px;
    background-color: rgb(10, 12, 24);
    min-height: 12vh;
    height: 12vh;

    border-radius: 10px 10px 0 0;
    transform: translateY(0px);
    transition: transform 1.5s ease-in;
    z-index: 0;
}


.rectangle.active {
    transform: translateY(-116px);
    transition: transform 1.5s ease-out;
    z-index: 98;
}

.input-wrapper {
    display: flex;
    align-items: center;
    justify-content: space-between;

    background: rgb(100, 100, 100);
    box-shadow: 0.25rem 0.25rem 0rem rgb(189, 189, 189);

    width: 75%;
    height: 40px;
    padding: 0.5rem;
    border-radius: 0.5rem;
    position: absolute;
    left: 10%;
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
    color: white;
    padding-left: 30px;
    padding-top: 10px;

}
</style>