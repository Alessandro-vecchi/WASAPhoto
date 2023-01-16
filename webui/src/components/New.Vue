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
            username: "",
            isFollowing: false,
            isBanned: false,
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
            console.log(this.$route.query.username, this.username)
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

      /*   async handleClick(cond, func) {
            this.loading = true;
            this.errormsg = null;
            /* The interceptor is modifying the headers of the requests being sent by adding an 'Authorization' header with a value that is stored in the browser's local storage. Just keeping the AuthToken in the header.
            If you don't use this interceptor, the 'Authorization' header with the token won't be added to the requests being sent, it can cause the requests to fail.
            
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            if (cond) {
                await this.$axios.delete("/photos/" + this.media.photoid + func + this.media.authorname);
            } else {
                await this.$axios.put("/photos/" + this.media.photoid + func + this.media.authorname);
            }
            this.loading = false;

        },
        handleFollowClick() {
            this.handleClick(this.isFollowing, "/followers/")
            this.isFollowing = !this.isFollowing;
        },

        handleBanClick() {
            this.handleClick(this.isBanned, "/bans/")
            this.isBanned = !this.isBanned;
        }, */

        refresh() {
            this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
            error => {return Promise.reject(error);});
            this.GetProfile().then(() => this.GetUserPhotos());
        },
        uploadImage: async function () {
            this.$router.push({ path: '/users/' + this.profile.user_id + "/form/"})
        },
        edit_profile: async function () {
            this.$router.push({ path: '/users/' + this.profile.user_id + "/editProfile/"})
        },
        change_username: async function () {
            this.$router.push({ path: '/users/' + this.profile.user_id + "/changeUsername" })
        }
    },
    computed: {
        logged() {
            // console.log(this.profile.user_id, localStorage.getItem('Authorization'))
            console.log("logged:", this.logged, this.profile.user_id)
            return localStorage.getItem('Authorization')==this.profile.user_id
        }
    },
    mounted() {
        this.refresh();
    }
}
</script>
<template><!-- 
    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg> -->
    <header class="header">
        <div class="wrapper">
            <div class="profile">
                <div class="profile-image">
                    <!--<img :src="profile.profilePicUrl" alt="">-->
                    <img :src=profile.profile_picture_url alt="Image" />
                </div>
                <div class="profile-user-settings">
                    <h1 class="profile-user-name"> {{ profile.username }}</h1>
                    <button v-if=logged type="button" class="btn edit-profile-button" @click="edit_profile">Edit profile</button>
                    <button v-if=logged type="button" class="btn change-username-button" @click="change_username">Change
                        Username</button>
                    <button v-if=!loading&!logged type="button" class="btn follow-button" @click="handleFollowClick">
                        <font-awesome-icon v-if=isFollowing class="follow-check"
                            icon="fa-solid fa-check" /><span>Follow</span></button>
                    <button v-if=!loading&!logged type="button" class="btn ban-button" @click="handleBanClick">
                        <font-awesome-icon v-if=isBanned class="ban-check"
                            icon="fa-solid fa-ban" /><span>Ban</span></button>

                </div>
                <div class="profile-stats">
                    <ul>
                        <li><span class="profile-stat-count">{{ profile.pictures_count }}</span> Posts</li>
                        <li><span class="profile-stat-count">456</span> Followers</li>
                        <li><span class="profile-stat-count">789</span> Following</li>
                    </ul>
                </div>
                <div class="profile-bio">
                    <p class="profile-bio-text">
                        <!-- Hi! My name is John and I'm here to kill you. -->
                        {{ profile.bio }}
                    </p>
                </div>

                <div class="upload-image">
                    <font-awesome-icon class="upload-image-button" icon="fa-solid fa-plus" size="3x" @click="uploadImage" />
                </div>
                <!--End of profile section-->
            </div>
        </div>
    </header>

    <div class="wrapper">
        <div class="gallery">
            <GalleryItem v-for="obj in media" :photo="obj"/>
           <!--  <GalleryItem />
            <GalleryItem />
            <GalleryItem />
            <GalleryItem />
            <GalleryItem />
            <GalleryItem />
            <GalleryItem /> -->
        </div>
    </div>
</template>

<style scoped>
.header {
    font-size: 10px;
    min-height: 25vh;
    background-color: #fafafa;
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
    background: none;
    border: none;
    color: inherit;
    padding: 0;
    cursor: pointer;
}

.profile-user-settings .btn:hover {
    font-size: 1.5rem;
}

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
    border: none;
    border-radius: 0.3rem;
    padding: 0 1.5rem;
    margin-left: 5rem;
    font-size: 16px;
}

.profile-user-settings .follow-button:hover {
    background-color: #050b85;
    /* Darker blu on hover */
    color: #fafafa;
}

.profile-user-settings .ban-button {
    background-color: #ec7b7b;
    /* Twitter blu */
    border: none;
    border-radius: 0.3rem;
    padding: 0 1.5rem;
    margin-left: 5rem;
    font-size: 16px;
}

.profile-user-settings .ban-button:hover {
    background-color: #b50707;
    /* Darker blu on hover */
    color: #fafafa;
}

.profile-user-settings .follow-button .follow-check {
    margin-left: 0;
}

.profile-stats {
    top: 7.5rem;
}

.profile-stats li {
    display: inline-block;
    font-size: 1.6rem;
    line-height: 1.5;
    margin-right: 4rem;
    cursor: pointer;
}

.profile-stats li:last-of-type {
    margin-right: 0;
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