<script>
import GalleryItem from "@/components/GalleryItem.vue"

export default {
    components: {
        GalleryItem,
    },
    data: function () {
        return {
            loading: false,
            errormsg: null,
            profile: {},
            media: [],
            followers: [],
            following: [],
            bans: [],
            header: localStorage.getItem('Authorization'),
            //isFollowing: false,
            //isBanned: false,
        }
    },
    methods: {
        async GetProfile() {
            this.loading = true;
            this.errormsg = null;
            /* The interceptor is modifying the headers of the requests being sent by adding an 'Authorization' header with a value that is stored in the browser's local storage. Just keeping the AuthToken in the header.
            If you don't use this interceptor, the 'Authorization' header with the token won't be added to the requests being sent, it can cause the requests to fail.
            */
            console.log("header:", localStorage.getItem('Authorization'))
            console.log("username:", this.$route.query.username, this.username)
            /* if (this.$route.query.username != this.username) {

            } */
            try {
                let response = await this.$axios.get("/users/?username=" + this.$route.query.username)
                this.profile = response.data
                this.username = this.profile.username
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
            console.log("profile1:", this.profile)

        },

        async GetUserPhotos() {
            this.loading = true;
            this.errormsg = null;
            console.log('profile:', this.profile)
            try {
                await this.$axios.get("/users/" + this.profile.user_id + "/photos/").then(response => (this.media = response.data));
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
            console.log("media:", this.media)
        },

        async HandleClick(cond, func) {
            this.loading = true;
            this.errormsg = null;
            /* The interceptor is modifying the headers of the requests being sent by adding an 'Authorization' header with a value that is stored in the browser's local storage. Just keeping the AuthToken in the header.
            If you don't use this interceptor, the 'Authorization' header with the token won't be added to the requests being sent, it can cause the requests to fail.
            */
            console.log(cond)
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            if (cond) {
                await this.$axios.delete("/users/" + this.profile.user_id + func + this.header);
            } else {
                await this.$axios.put("/users/" + this.profile.user_id + func + this.header);
            }
            this.loading = false;

        },
        handleFollowClick() {
            this.HandleClick(this.isFollower(), "/followers/")
            //this.isFollowing = !this.isFollowing;
        },

        handleBanClick() {
            this.HandleClick(this.isBanned(), "/bans/")
            //this.isBanned = !this.isBanned;
        },
        async GetUsers(goal, isRefresh) {
            this.loading = true;
            this.errormsg = null;
            let list = [];
            /* The interceptor is modifying the headers of the requests being sent by adding an 'Authorization' header with a value that is stored in the browser's local storage. Just keeping the AuthToken in the header.
            If you don't use this interceptor, the 'Authorization' header with the token won't be added to the requests being sent, it can cause the requests to fail.
            */
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            try {
                let response = await this.$axios.get("/users/" + this.profile.user_id + "/" + goal + "/")
                list = response.data
                if (!isRefresh) {
                    this.$router.push({ path: '/users/' + this.profile.user_id + "/listUsers/" })
                }
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
            console.log(goal + ":", list)
            return list;
        },
        async getFollowers(isRefresh) {
            this.followers = await this.GetUsers("followers", isRefresh)
            if (!isRefresh) {
                this.refresh()
            }
        },
        async getFollowing() {
            this.following = await this.GetUsers("following", false)
        },
        async getBans(isRefresh) {
            this.bans = await this.GetUsers("bans", isRefresh)
            if (!isRefresh) {
                this.refresh()
            }
        },
        isFollower() {
            console.log("log follow:", this.logged)
            if (this.logged) {
                return false
            }
            console.log(this.username, "in", this.followers)
            return this.username in this.followers
        },
        isBanned() {
            console.log("log ban:", this.logged)
            if (this.logged) {
                return false
            }
            return this.username in this.bans
        },
        refresh() {
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
           /*  this.GetProfile().then(() => this.GetUserPhotos());
            this.getBans(true).then(() => this.getFollowers(true)) */
            this.GetProfile().then(() => this.getBans(true)).then(() => this.getFollowers(true)).then(() => this.GetUserPhotos());
        },
        uploadImage: async function () {
            this.$router.push({ path: '/users/' + this.profile.user_id + "/form/" })
        },
        edit_profile: async function () {
            this.$router.push({ path: '/users/' + this.profile.user_id + "/editProfile/" })
        },
        change_username: async function () {
            this.$router.push({ path: '/users/' + this.profile.user_id + "/changeUsername/" })
        },
        cancel() {
            this.$router.push({ path: "/users/" + this.header + "/stream/" });
        },
    },
    computed: {
        username: {
            get() {
                return this.username;
            },
            set(value) {
                // prevent the username from being changed once it has a value
                if (!this.username) {
                    this.$set(this, 'username', value)
                }
            },
        },
        logged() {
            // console.log(this.profile.user_id, localStorage.getItem('Authorization'))
            let bool = (this.header == this.profile.user_id)
            console.log("logged:", bool, this.profile.user_id)
            if (bool === null) {
                return false
            }
            return bool
        },
    },
    created() {
        this.refresh();
    }
}
</script>
<template>
    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
    <header class="header">
        <div class="wrapper">
            <div class="profile">
                <font-awesome-icon class="previous-page" icon="fa-solid fa-chevron-left" size="5x" @click="cancel" />
                <div class="profile-image">
                    <img :src=profile.profile_picture_url alt="Image" />
                </div>
                <div class="profile-user-settings">
                    <h1 class="profile-user-name"> {{ profile.username }}</h1>
                    <button v-if=logged type="button" class="btn edit-profile-button" @click="edit_profile">Edit
                        profile</button>
                    <button v-if=logged type="button" class="btn change-username-button" @click="change_username">Change
                        Username</button>
                    <button v-if="!loading && !logged" type="button" class="btn follow-button"
                        @click="handleFollowClick">
                        <font-awesome-icon v-if=isFollower() class="check" icon="fa-solid fa-check" /><font-awesome-icon
                            v-else class="check" icon="fa-solid fa-xmark" /><span class="action">Follow</span></button>
                    <button v-if="!loading && !logged" type="button" class="btn ban-button" @click="handleBanClick">
                        <font-awesome-icon v-if=isBanned() class="check" icon="fa-solid fa-ban" /><span
                            class="action">Ban</span></button>

                </div>
                <div class="profile-stats">
                    <ul>
                        <li><span class="profile-stat-count">{{ profile.pictures_count }}</span> Posts</li>
                        <li><span class="profile-stat-count">{{ profile.followers_count }}</span> <span
                                @click="getFollowers">Followers</span></li>
                        <li><span class="profile-stat-count">{{ profile.follows_count }}</span> <span
                                @click="getFollowing">Following</span></li>
                    </ul>
                </div>
                <div class="profile-bio">
                    <p class="profile-bio-text">
                        <!-- Hi! My name is John and I'm here to kill you. -->
                        {{ profile.bio }}
                    </p>
                </div>

                <div class="upload-image">
                    <font-awesome-icon class="upload-image-button" icon="fa-solid fa-plus" size="3x"
                        @click="uploadImage" />
                </div>
                <!--End of profile section-->
            </div>
        </div>
    </header>

    <div class="wrapper">
        <div class="gallery">
            <GalleryItem v-for="obj in media" :photo="obj" />
        </div>
    </div>
</template>

<style scoped>
.header {
    font-size: 10px;
    min-height: 25vh;
    padding-bottom: 1rem;
}

.wrapper {
    max-width: 93.5rem;
    margin: 0 auto;
    padding: 0 2rem;
}

.wrapper::after {
    box-sizing: border-box;
}

img {
    display: block;
}

/*Profile section */
.header .wrapper .profile {
    padding: 5rem 0;
    position: relative;
    height: 300px;
    background-color: #fafafa;
}

.header .wrapper .profile .previous-page {
    position: absolute;
    top: 25px;
    left: 25px;
    cursor: pointer;
}

.header .wrapper .profile .profile-image {
    width: calc(33.33% - 1rem);
    position: absolute;
    left: 15%;
}

.header .wrapper .profile .profile-image img {
    width: 20vh;
    border-radius: 50%;
    cursor: pointer;
}


.profile-user-settings,
.profile-stats,
.profile-bio {
    position: absolute;
    left: 40%;
    width: calc(66.66% -2rem);
}

.profile-user-settings .profile-user-name {
    display: inline-block;
    font-size: 3rem;
}

.profile-user-settings .btn {
    display: inline-block;
    font: inherit;
    font-weight: 600;
    background: none;
    border: none;
    color: inherit;
    padding: 0;
    cursor: pointer;
}

/* .profile-user-settings .btn:hover {
    font-size: 1.5rem;
} */

.profile-user-settings {
    top: 2rem;
}

.profile-user-settings .edit-profile-button {
    font-size: 1.4rem;
    line-height: 1.7;
    border: 0.1rem solid #dbdbdb;
    border-radius: 0.3rem;
    padding: 1rem 2.4rem;
    margin-left: 5rem;
}

.profile-user-settings .change-username-button {
    font-size: 1rem;
    line-height: 1;
    border: 0.1rem solid #dbdbdb;
    border-radius: 0.3rem;
    padding: 1.1rem 2.4rem;
    margin-left: 5rem;
}

.profile-user-settings .change-username-button:hover {
    font-size: 1.2rem;
    background-color: #fff;
}

.profile-user-settings .follow-button {
    background-color: #00acee;
    /* Twitter blu */
    width: 120px;
    border-radius: 0.3rem;
    padding: 0.3rem 1.4rem;
    margin-left: 5rem;
    font-size: 17px;
}

.profile-user-settings .follow-button:hover {
    background-color: #050b85;
    /* Darker blu on hover */
    color: #fafafa;
}

.profile-user-settings .ban-button {
    background-color: #ec7b7b;
    /* Twitter blu */
    width: 100px;
    border-radius: 0.3rem;
    padding: 0.3rem 1.4rem;
    margin-left: 5rem;
    font-size: 16px;
    /*  box-shadow: 0 5px #ec7b7b; */
}

/* 
.profile-user-settings .ban-button:active {
    box-shadow: 0 2px #ec7b7b;
    transform: traslateY(4px)
} */
.profile-user-settings .ban-button:hover {
    background-color: #b50707;
    /* Darker blu on hover */
    color: #fafafa;
}

.check {
    float: left;
    margin-top: 3px;
}

.action {
    float: right;
}

.profile-user-settings .follow-button .follow-check {
    margin-left: 0;
}

.profile-stats {
    top: 8rem;
}

.profile-stats li {
    display: inline-block;
    font-size: 1.6rem;
    line-height: 1.5;
    margin-right: 5rem;
    cursor: pointer;
}

.profile-stats li:last-of-type {
    margin-right: 0;
}

.profile-stats li span:last-of-type:hover {
    text-decoration: underline;
}

.profile-stats li span:first-of-type:hover {
    text-decoration: none;
}

.profile-bio {
    font-size: 1.6rem;
    font-weight: 400;
    line-height: 1.5;
    margin-top: 7rem;
}

.profile-stat-count,
.edit-profile-button {
    font-weight: 600;
}

.upload-image {
    position: absolute;
    right: 0;
    bottom: 0;
}

.upload-image-button {
    width: 30px;
    height: 30px;
    border-radius: 50%;

    background: rgb(255, 255, 255);
    transition: background 0.2s, color 0.2s;
}

.upload-image-button:hover {
    cursor: pointer;
    background: rgba(39, 55, 69, 1);
    color: #fafafa;
}

.gallery {
    display: flex;
    flex-wrap: wrap;
    padding-bottom: 3rem;
    margin: -1rem -1rem;
}


@supports (display:grid) {

    .gallery {
        display: grid;
        grid-template-columns: repeat(3, auto);
        grid-gap: 1rem;
    }

    .gallery {
        width: auto;
        margin: 0;
    }
}
</style>