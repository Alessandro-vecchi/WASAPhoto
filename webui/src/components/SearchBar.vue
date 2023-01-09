<script>
export default {

    data: function () {
        return {
            errormsg: null,
            loading: false,
            username: "",
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
        restart_search() {
            this.$route.params.username = ""
            onclick = "document.getElementById('search').value = ''"
        }

    },
    mounted() {
        this.refresh()
    }
}
</script>

<template>
    <div class="rectangle">
        <p class="title font-style">Search</p>
        <div class="input-wrapper">
            <font-awesome-icon class="icons" icon="fa-solid fa-magnifying-glass" inverse @click="get_user_profile" />
            <input id="search" v-model="username" type="text" placeholder="Search" />
            <font-awesome-icon class="icons" icon="fa-solid fa-xmark"
                onclick="document.getElementById('search').value = ''" />
        </div>
    </div>
</template>

<style>
.rectangle {
    position:absolute;
    bottom: 0;
    left: 30%;
    right: 30%;
    background-color: rgb(10, 12, 24);
    height: 12vh;

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