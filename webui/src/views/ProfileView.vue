<script>
import GalleryItem from "@/components/GalleryItem.vue"
import Avatar from "@/components/Avatar.vue"
import { eventBus } from "@/main.js"

export default {
    components: {
        GalleryItem,
        Avatar,
    },
    data: function () {
        return {
            loading: false,
            errormsg: null,
            profile: {},
            username: "",
            media: [],
            followers: [],
            following: [],
            bans: [],
            header: localStorage.getItem('Authorization'),
            isFollowing: false,
            iAmBanned: false,
            isBanned: false,
            ppUrl: "",
        }
    },
    methods: {
        async GetProfile() {
            this.loading = true;
            this.errormsg = null;
            /* The interceptor is modifying the headers of the requests being sent by adding an 'Authorization' header with a value that is stored in the browser's local storage. Just keeping the AuthToken in the header.
            If you don't use this interceptor, the 'Authorization' header with the token won't be added to the requests being sent, it can cause the requests to fail.
            */
            console.log("header:", localStorage.getItem('Authorization'), "\n", "username:", this.$route.query.username, this.username)
            try {
                let response = await this.$axios.get("/users/", { params: { username: this.$route.query.username } })
                this.profile = response.data
                this.username = this.profile.username
                console.log("GetProfile")
            } catch (e) {
                this.errormsg = e.response.data.error.toString();
                if (this.errormsg === "User not found") {
                    this.$router.push("/error/", { props: { msg: this.errormsg } });
                }
            }
            this.loading = false;
            console.log("profile1:", this.profile)

        },

        async GetUserPhotos() {
            this.loading = true;
            this.errormsg = null;
            console.log('profile:', this.profile)
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            try {
                await this.$axios.get("/users/" + this.profile.user_id + "/photos/").then(response => (this.media = response.data));
            } catch (e) {
                this.errormsg = e.response.data.error.toString();
            }
            this.loading = false;
            console.log("media:", this.media)
        },
        async FollowClick() {
            this.loading = true;
            this.errormsg = null;
            console.log(this.isFollowing)
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            if (this.isFollowing) {
                await this.$axios.delete("/users/" + this.profile.user_id + "/followers/" + this.header).then(() => (this.profile.followers_count--, this.isFollowing = false)).catch(e => this.errormsg = e.response.data.error.toString());
            } else {
                await this.$axios.put("/users/" + this.profile.user_id + "/followers/" + this.header).then(() => (this.profile.followers_count++, this.isFollowing = true)).catch(e => this.errormsg = e.response.data.error.toString());
            }
            this.loading = false;

        },

        async BanClick() {
            this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            if (this.isBanned) {
                await this.$axios.delete("/users/" + this.profile.user_id + "/bans/" + this.header).then(() => this.isBanned = false).catch(e => this.errormsg = e.response.data.error.toString());
            } else {
                await this.$axios.put("/users/" + this.profile.user_id + "/bans/" + this.header).then(() => this.isBanned = true).catch(e => this.errormsg = e.response.data.error.toString());
                if (this.logged) {
                    // adding a ban remove my follow from the profile
                    if (this.isFollowing) {
                        await this.$axios.delete("/users/" + this.profile.user_id + "/followers/" + this.header).catch(e => this.errormsg = e.response.data.error.toString())
                        this.isFollowing = false
                        this.profile.followers_count--
                    }
                    // adding a ban remove his follow from my profile
                    await this.getFollowing(true).catch(e => this.errormsg = e.response.data.error.toString())
                    if (this.isFollowed) {
                        await this.$axios.delete("/users/" + this.header + "/followers/" + this.profile.user_id).catch(e => this.errormsg = e.response.data.error.toString());
                        this.isFollowed = false
                        this.profile.follows_count--
                    }
                }
            }
            this.loading = false;
        },
        async GetUsers(goal, isRefresh) {
            this.loading = true;
            this.errormsg = null;
            let list = {};
            /* The interceptor is modifying the headers of the requests being sent by adding an 'Authorization' header with a value that is stored in the browser's local storage. Just keeping the AuthToken in the header.
            If you don't use this interceptor, the 'Authorization' header with the token won't be added to the requests being sent, it can cause the requests to fail.
            */
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            try {
                let response = await this.$axios.get("/users/" + this.profile.user_id + "/" + goal + "/")
                list = response.data
                if (!isRefresh) {
                    eventBus.getShortProfiles = list.short_profile
                    eventBus.getTitle = goal
                    eventBus.getUsername = this.profile.username
                    this.$router.push({ path: '/' + goal + '/', });
                }
            } catch (e) {
                this.errormsg = e.response.data.error.toString();
            }
            this.loading = false;
            console.log(goal + ":", list)
            return list;
        },
        async getFollowers(isRefresh) {
            console.log("GetFollowers")
            let list = await this.GetUsers("followers", isRefresh).catch(e => this.errormsg = e.response.data.error.toString())

            this.followers = list.short_profile
            this.isFollowing = list.cond
            console.log("follow:", list.short_profile, list.cond)
            if (!isRefresh) {
                this.refresh()
            }
        },
        async getFollowing(isBug = false) {
            let list = await this.GetUsers("following", isBug).catch(e => this.errormsg = e.response.data.error.toString())
            this.following = list.short_profile
            this.isFollowed = list.cond
        },
        async getMyBans(isRefresh) {
            this.loading = true;
            this.errormsg = null;
            /* The interceptor is modifying the headers of the requests being sent by adding an 'Authorization' header with a value that is stored in the browser's local storage. Just keeping the AuthToken in the header.
            If you don't use this interceptor, the 'Authorization' header with the token won't be added to the requests being sent, it can cause the requests to fail.
            */
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            try {
                let response = await this.$axios.get("/users/" + this.profile.user_id + "/mybans/")
                this.bans = response.data.short_profile
                this.isBanned = response.data.cond
                if (!isRefresh) {
                    eventBus.getShortProfiles = this.bans
                    eventBus.getTitle = "bans"
                    eventBus.getUsername = this.profile.username
                    this.$router.push({ path: '/bans/' })
                }
            } catch (e) {
                this.errormsg = e.response.data.error.toString();
            }
            this.loading = false;
            console.log("banned:", this.bans, this.isBanned)
            if (!isRefresh) {
                this.refresh()
            }
        },
        async getTheirBans() {
            this.loading = true;
            this.errormsg = null;
            /* The interceptor is modifying the headers of the requests being sent by adding an 'Authorization' header with a value that is stored in the browser's local storage. Just keeping the AuthToken in the header.
            If you don't use this interceptor, the 'Authorization' header with the token won't be added to the requests being sent, it can cause the requests to fail.
            */
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            try {
                let response = await this.$axios.get("/users/" + this.profile.user_id + "/bans/")
                this.bans = response.data.short_profile
                this.iAmBanned = response.data.cond
            } catch (e) {
                this.errormsg = e.response.data.error.toString();
            }
            this.loading = false;
            console.log("banned:", this.bans, this.iAmBanned)
        },
        async getImage() {
            this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            console.log("GetImage")
            try {
                let response = await this.$axios.get("/images/?image_name=" + this.profile.profile_picture_url, { responseType: 'blob' })
                // Get the image data as a Blob object
                var imgBlob = response.data;

                // Create an object URL from the Blob object
                this.ppUrl = URL.createObjectURL(imgBlob);
                console.log("GetImage")
            } catch (e) {
                this.errormsg = e.response.data.error.toString();

            }
            this.loading = false;
            console.log("profile_pic_URL:", this.ppUrl)
        },
        async refresh() {
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            this.loading=true
            await this.GetProfile()
            if (this.profile.profile_picture_url) { await this.getImage() }
            this.getTheirBans().then(() => this.getMyBans(true))
            if (!this.iAmBanned) { this.getFollowers(true).then(() => this.GetUserPhotos()).then(() => console.log("refresh:", this.media)) }
            this.loading=false
        },

        uploadImage: function () {
            this.$router.push({ path: '/users/' + this.profile.user_id + "/form/" })
        },
        edit_profile: function () {
            this.$router.push({ path: '/users/' + this.profile.user_id + "/editProfile/" })
        },
        change_username: function () {
            this.$router.push({ path: '/users/' + this.profile.user_id + "/changeUsername/" })
        },
        cancel() {
            this.$router.push({ path: "/users/" + this.header + "/stream/" });
        },
    },
    computed: {
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
    mounted() {
        this.$axios.get("/users/", { params: { username: this.$route.query.username } }).then(() => this.refresh()).catch(e => {
            if (e.response.data.error.toString() === "User not found") {
                this.$router.push("/error/" + "404");
            }
        })
    },
}
</script>
<template>
    <div v-if="!loading" class="wrapper">
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <div class="profile">
            <font-awesome-icon class="previous-page" icon="fa-solid fa-chevron-left" size="5x" @click="cancel" />
            <div v-if=!this.iAmBanned class="profile-image">
                <Avatar :src="ppUrl" :size="180" />
            </div>
            <div v-else class="profile-image">
                <Avatar :size="180" />
            </div>
            <div class="profile-user-settings">
                <h1 class="profile-user-name"> {{ profile.username }}</h1>
                <button v-if=logged type="button" class="btn edit-profile-button" @click="edit_profile">Edit
                    profile</button>
                <button v-if=logged type="button" class="btn change-username-button" @click="change_username">Change
                    Username</button>
                <button v-if="(!logged && !iAmBanned && !isBanned)" type="button" class="btn follow-button"
                    @click="FollowClick">
                    <font-awesome-icon v-if=isFollowing class="check" icon="fa-solid fa-check" /><font-awesome-icon
                        v-else class="check" icon="fa-solid fa-xmark" /><span class="action">Follow</span></button>
                <button v-if="(!logged && !this.iAmBanned)" type="button" class="btn ban-button"
                    @click="BanClick">
                    <font-awesome-icon v-if=isBanned class="check" icon="fa-solid fa-ban" /><span
                        class="action">Ban</span></button>
                <button v-else-if="(!logged && !this.isBanned)" type="button" class="btn ban-button"
                    @click="BanClick"><font-awesome-icon v-if=isBanned class="check" icon="fa-solid fa-ban" /><span
                        class="action">Ban</span></button>
                <button v-else-if="(!logged)" type="button" class="btn ban-button"
                    @click="BanClick"><font-awesome-icon v-if=isBanned class="check" icon="fa-solid fa-ban" /><span
                        class="action">Ban</span></button>


            </div>
            <div class="profile-stats">
                <ul>
                    <li v-if=!this.iAmBanned><span class="profile-stat-count">{{ profile.pictures_count }}</span> Posts
                    </li>
                    <li v-if=!this.iAmBanned><span class="profile-stat-count">{{ profile.followers_count }}</span> <span
                            @click="getFollowers(false)">Followers</span></li>
                    <li v-if=!this.iAmBanned><span class="profile-stat-count">{{ profile.follows_count }}</span> <span
                            @click="getFollowing(false)">Following</span></li>
                    <li v-if="(logged)"><button type="button" class="btn see-bans-button"
                            @click="getMyBans(false)">View banned users</button></li>
                </ul>
            </div>
            <div v-if=!this.iAmBanned class="profile-bio">
                <p class="profile-bio-text">
                    {{ profile.bio }}
                </p>
            </div>

            <div v-if=logged class="upload-image">
                <font-awesome-icon class="upload-image-button" icon="fa-solid fa-plus" @click="uploadImage" />
            </div>
            <!--End of profile section-->
        </div>
        <div v-if=!this.iAmBanned class="gallery">
            <GalleryItem v-for="obj in media" :key="obj.photoId" :photo="obj" />
        </div>
    </div>
</template>

<style scoped>
.wrapper {
    min-height: 25vh;
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
.wrapper .profile {
    padding: 5rem 0;
    position: relative;
    height: 300px;
    background-color: #fafafa;
}

.wrapper .profile .previous-page {
    position: absolute;
    top: 25px;
    left: 25px;
    cursor: pointer;
}

.wrapper .profile .profile-image {
    width: calc(33.33% - 1rem);
    position: absolute;
    top: 8%;
    left: 14%;
}

/* 
.header .wrapper .profile .profile-image img {
    width: 20vh;
    border-radius: 50%;
    cursor: pointer;
} */


.profile-user-settings,
.profile-stats,
.profile-bio {
    position: absolute;
    left: 40%;
    width: calc(66.66% - 2rem);
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
    transition: transform 0.3s;
}

.profile-user-settings .edit-profile-button:hover {
    background-color: #fff;
    transform: scale(1.05);
}

.profile-user-settings .change-username-button {
    font-size: 1rem;
    line-height: 1;
    border: 0.1rem solid #dbdbdb;
    border-radius: 0.3rem;
    padding: 1.1rem 2.4rem;
    margin-left: 5rem;
    transition: transform 0.2s;
}

.profile-user-settings .change-username-button:hover {
    background-color: #fff;
    transform: scale(1.1);
}

.see-bans-button {
    font-size: 0.6rem;
    line-height: 0.8;
    border: 0.1rem solid #dbdbdb;
    border-radius: 0.3rem;
    padding: 0.8rem 1.4rem;
    margin-left: -3rem;
    transition: transform 0.4s;
}

.see-bans-button:hover {
    transform: scale(1.3);
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
    top: 8.8rem;
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
    margin-top: 7.5rem;
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
    width: 40px;
    height: 40px;
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
        grid-template-columns: repeat(3, calc((93.5rem - 2rem - 4rem)/3));
        grid-gap: 1rem;
    }

    .gallery {
        width: auto;
        margin: 0;
    }
}
</style>