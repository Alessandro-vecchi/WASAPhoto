<script>
import image from "../assets/images/552480.jpeg"
import GalleryItem from "@/components/GalleryItem.vue"

export default
    {
        components: {
            GalleryItem,
        },
        data: function () {
            return {
                loading: false,
                errmsg: null,
                profile: {},
                media: [],
                image,
                logged: localStorage.getItem('Authorization'),
            }
        },
        methods: {
            async GetProfile() {
                this.loading = true;
                this.errormsg = null;
                /* The interceptor is modifying the headers of the requests being sent by adding an 'Authorization' header with a value that is stored in the browser's local storage. Just keeping the AuthToken in the header.
                If you don't use this interceptor, the 'Authorization' header with the token won't be added to the requests being sent, it can cause the requests to fail.
                */
                console.log(this.profile, localStorage.getItem('Authorization'))
                this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
                error => {return Promise.reject(error);});
                try {
                    let response = await this.$axios.get("/users/?username=")
                    this.profile = response.data
                } catch (e) {
                    this.errormsg = e.toString();
                }
                this.loading = false;
                console.log(this.profile)
            },

            async GetUserPhotos() {
                this.loading = true;
                this.errormsg = null;
                //this.$axios.interceptors.request.use(config => {config.headers['Authorization'] = localStorage.getItem('Authorization');return config;},
                //error => {return Promise.reject(error);});
                try {
                    this.$axios.get("/users/" + this.profile.user_id + "/photos/").then(response => (this.media = response.data));
                } catch (e) {
                    this.errormsg = e.toString();
                }
                this.loading = false;
            },
                
            refresh() {
			this.GetProfile();
			//this.GetUserPhotos();
		    },
            createMedia: async function () {
                this.$router.push({ path: '/users/' + this.profile.userid + "/photos/" })
            },
            edit_profile: async function () {
                this.$router.push({ path: '/users/' + this.profile.user_id + "/editProfile/",  props: { user_id: this.profile.user_id } })
            },
            changeUsername: async function () {
                this.$router.push({ path: '/users/' + this.profile.username + "/changeUsername" })
            }
        },
        mounted() {
            this.refresh();
        }
    }
</script>
<template>
    <header class="header">
        <div class="wrapper">
            <div class="profile">
                <div class="profile-image">
                    <!--<img :src="profile.profilePicUrl" alt="">-->
                    <img :src=profile.profile_picture_url alt="Mickey Mouse" />
                </div>
                <div class="profile-user-settings">
                    <h1 class="profile-user-name"> {{ profile.username }}</h1>
                    <button type="button" class="btn edit-profile-button" @click="edit_profile">Edit profile</button>
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
                    <font-awesome-icon class="upload-image-button" icon="fa-solid fa-plus" size="3x" />
                </div>
                <!--End of profile section-->
            </div>
        </div>
    </header>

    <div class="wrapper">
        <div class="gallery">
            <GalleryItem />
            <GalleryItem />
            <GalleryItem />
            <GalleryItem />
            <GalleryItem />
            <GalleryItem />
            <GalleryItem />
            <GalleryItem />
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
.profile {
    padding: 5rem 0;
    position: relative;
}

.profile::after {
    content: '';
    display: block;
    clear: both;
}

.profile-image {
    float: left;
    width: calc(33.33% - 1rem);
    display: flex;
    justify-content: center;
    align-items: center;
    margin-right: 3rem;
}

.profile-image img {
    width: 20vh;
    border-radius: 50%;
    cursor: pointer;
}


.profile-user-settings,
.profile-stats,
.profile-bio {
    float: left;
    width: calc(66.66% -2rem);
}

.profile-user-name {
    display: inline-block;
    font-size: 3rem;
}

.btn {
    display: inline-block;
    font: inherit;
    background: none;
    border: none;
    color: inherit;
    padding: 0;
    cursor: pointer;
}

.btn:hover {
    font-size: 1.5rem;
}

.edit-profile-button {
    font-size: 1.4rem;
    line-height: 1.7;
    border: 0.1rem solid #dbdbdb;
    border-radius: 0.3rem;
    padding: 0 2.4rem;
    margin-left: 2rem;
}

.profile-stats {
    margin-top: 2.3rem;
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
    margin-top: 2.3rem;
}

.profile-stat-count,
.edit-profile-button {
    font-weight: 600;
}
.upload-image {
    position: absolute;
    right:0;
    bottom:0;
}
.upload-image-button {
    width: 30px;
    height: 30px;
	border-radius: 50%;

	background: rgb(255,255,255);
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