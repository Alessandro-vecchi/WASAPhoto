<script>
import GalleryItem from "@/components/GalleryItem.vue"
import Avatar from "@/components/Avatar.vue"
import { eventBus} from "@/main.js"

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
            media: [],
            followers: [],
            following: [],
            bans: [],
            header: localStorage.getItem('Authorization'),
            isFollowing: null,
            isBanned: null,
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
            console.log("header:", localStorage.getItem('Authorization'))
            console.log("username:", this.$route.query.username, this.username)
            try {
                let response = await this.$axios.get("/users/?username=" + this.$route.query.username)
                this.profile = response.data
                this.username = this.profile.username
                // await this.GetBans(true)// .catch(e => (this.errormsg = e.toString()))
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
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            try {
                await this.$axios.get("/users/" + this.profile.user_id + "/photos/").then(response => (this.media = response.data));
            } catch (e) {
                // console.log(e)
                this.errormsg = e.toString();
                console.log(this.errormsg)
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
        async handleFollowClick() {
            await this.HandleClick(this.isFollowing, "/followers/")
            this.refresh()
        },

        async handleBanClick() {
            await this.HandleClick(this.isBanned, "/bans/")
            this.refresh()
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
                    this.$router.push({ path: '/' + goal + '/', });
                }
            } catch (e) {
                console.error(e.message)
                this.errormsg = e.toString();
            }
            this.loading = false;
            console.log(goal + ":", list)
            return list;
        },
        async getFollowers(isRefresh) {
            let list = await this.GetUsers("followers", isRefresh)
            this.followers = list.short_profile
            this.isFollowing = list.cond
            console.log("follow:", list.short_profile, list.cond)
            if (!isRefresh) {
                this.refresh()
            }
        },
        async getFollowing() {
            this.following = await this.GetUsers("following", false)
        },
        async GetBans(isRefresh) {
            console.log("hi")
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
                this.isBanned = response.data.cond
                if (!isRefresh) {
                    this.$router.push({ path: '/users/' + this.header + "/listUsers/" })
                }
            } catch (e) {
                console.error(e.message)
                this.errormsg = e.toString();
            }
            this.loading = false;
            console.log("banned:", this.bans, this.isBanned)
            if (!isRefresh) {
                this.refresh()
            }
        },
        async getImage() {
            this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            try {
                let response = await this.$axios.get("/images/?image_name=" + this.profile.profile_picture_url, { responseType: 'blob' })
                // Get the image data as a Blob object
                var imgBlob = response.data;

                // Create an object URL from the Blob object
                this.ppUrl = URL.createObjectURL(imgBlob);
            } catch {
                this.errormsg = error.message;

            }
            this.loading = false;
            console.log("profile_pic_URL:", this.ppUrl)
        },
        refresh() {
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });/* .then(() => this.GetBans(true)) */
            this.GetProfile().then(() => this.getImage()).then(() => this.getFollowers(true)).then(() => this.GetUserPhotos()).then(() => console.log("refresh:", this.media))
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
        username: {
            get() {
                return this.username;
            },
            set(value) {
                // prevent the username from being changed once it has a value
                if (!this.username) {
                    this.$set(this, 'username', value)
                    console.log(this.username)
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
    mounted() {
        this.refresh();
    }
}
</script>
<template>
    <div class="wrapper">
        <div class="profile">
            <ErrorMsg v-if="errormsg" :msg="errormsg" />
            <font-awesome-icon class="previous-page" icon="fa-solid fa-chevron-left" size="5x" @click="cancel" />
            <div class="profile-image">
                <Avatar :src="ppUrl" :size="180" />
            </div>
            <div class="profile-user-settings">
                <h1 class="profile-user-name"> {{ profile.username }}</h1>
                <button v-if=logged type="button" class="btn edit-profile-button" @click="edit_profile">Edit
                    profile</button>
                <button v-if=logged type="button" class="btn change-username-button" @click="change_username">Change
                    Username</button>
                <button v-if="(!loading && !logged)" type="button" class="btn follow-button" @click="handleFollowClick">
                    <font-awesome-icon v-if=isFollowing class="check" icon="fa-solid fa-check" /><font-awesome-icon
                        v-else class="check" icon="fa-solid fa-xmark" /><span class="action">Follow</span></button>
                <button v-if="(!loading && !logged)" type="button" class="btn ban-button" @click="handleBanClick">
                    <font-awesome-icon v-if=isBanned class="check" icon="fa-solid fa-ban" /><span
                        class="action">Ban</span></button>

            </div>
            <div class="profile-stats">
                <ul>
                    <li><span class="profile-stat-count">{{ profile.pictures_count }}</span> Posts</li>
                    <li><span class="profile-stat-count">{{ profile.followers_count }}</span> <span
                            @click="getFollowers(false)">Followers</span></li>
                    <li><span class="profile-stat-count">{{ profile.follows_count }}</span> <span
                            @click="getFollowing(false)">Following</span></li>
                </ul>
            </div>
            <div class="profile-bio">
                <p class="profile-bio-text">
                    {{ profile.bio }}
                </p>
            </div>

            <div v-if=logged class="upload-image">
                <font-awesome-icon class="upload-image-button" icon="fa-solid fa-plus" @click="uploadImage" />
            </div>
            <!--End of profile section-->
        </div>
        <div class="gallery">
            <GalleryItem v-if=media v-for="obj in media" :key="obj.photoId" :photo="obj" />
        </div>
    </div>
</template>

<style scoped>
/* .header {
    font-size: 10px;
    min-height: 25vh;
    padding-bottom: 1rem;
    outline: 1px red;
} */

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