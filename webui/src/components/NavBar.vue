<script>
export default {
    data: function () {
        return {
            Username: "",
            header: localStorage.getItem('Authorization'),
            errormsg: null,
            loading: false
        }
    },
    methods: {
        async search_user_profile() {
            this.loading = true;
            this.errormsg = null;
            /* The interceptor is modifying the headers of the requests being sent by adding an 'Authorization' header with a value that is stored in the browser's local storage. Just keeping the AuthToken in the header.
            If you don't use this interceptor, the 'Authorization' header with the token won't be added to the requests being sent, it can cause the requests to fail.
            */
            console.log("header:", localStorage.getItem('Authorization'), "\n", "username:", this.$route.query.username)
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            try {
                await this.$axios.get("/users/?username=" + this.Username)
                this.$router.push({ path: "/users/", query: { username: this.Username } })
            } catch (e) {
                this.errormsg = e.response.data.error.toString();
                this.Username = ""
            }
            this.loading = false;

        },
        get_user_profile()  {
                this.$router.push({ path: "/users/", query: { username: this.Username } })
        },
        rect() {
            const rectangle = document.querySelector('.rectangle');
            rectangle.classList.toggle('active')

        },
        goHome() {
            this.$router.push({ path: "/users/" + this.header + "/stream/" });
        },

    },
}
</script>

<template>
    <nav>
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <div class="nav-wrapper center">

            <div id="nav-home-section" class="nav-section">
                <font-awesome-icon class="icons" icon="fa-solid fa-house" size="xl" inverse @click="goHome" />
                <span class="font-style" @click="goHome">Home</span>
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
                        @click="search_user_profile" />
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
    --heightNavBar: 160px;
    --minHeightNavBar: 12vh;
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
    min-height: var(--minHeightNavBar);
    height: var(--heightNavBar);
    bottom: 0;

    z-index: 99;
}

.center {
    display: flex;
    align-items: center;
    justify-content: center;

}

.nav-wrapper>.nav-section {
    padding: 5rem 0;
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
    min-height: var(--minHeightNavBar);
    height: var(--heightNavBar);

    border-radius: 10px 10px 0 0;
    transform: translateY(0px);
    transition: transform 1.5s ease-in;
    z-index: 0;
}


.rectangle.active {
    transform: translateY(calc(-1*var(--heightNavBar)));
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