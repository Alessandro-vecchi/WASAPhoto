<script>
import image from "../assets/images/552480.jpeg"
export default {
    data: function () {
        return {
            loading: false,
            errmsg: null,
            profile: {
                user_id: "",
                username: "",
                pictures_count: 0,
                followers_count: 0,
                following_count: 0,
                bio: "",
                profile_picture: "",

            },
            media: [],
            image: image,
        }
    },
    methods: {
        async GetProfile() {
            this.loading = true;
            this.errormsg = null;
            try {
                this.$axios.get("/users/?username=" + this.$route.params.username).then(response => (this.profile = response.data));
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        async GetMedia() {
            this.loading = true;
            this.errormsg = null;
            try {
                this.$axios.get("/users/:userid=" + localStorage.getItem('Authorization') + "/photos/").then(response => (this.media = response.data));
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        createMedia: async function () {
            this.$router.push({ path: '/users/' + this.profile.userid + "/photos/" })
        },
        updateProfile: async function () {
            this.$router.push({ path: '/users/' + this.profile.username + "/updateProfile" })
        },
        changeUsername: async function () {
            this.$router.push({ path: '/users/' + this.profile.username + "/changeUsername" })
        }
    },
    mounted() {
        this.GetMedia();
        this.GetProfile();
    }
}
</script>
<template>
    <header>
        <div class="wrapper">
            <div class="profile">
                <div class="profile-image">
                    <!--<img :src="profile.profilePicUrl" alt="">-->
                    <img :src="image" alt="Mickey Mouse" />
                </div>
            </div>
            <div class="profile-user-settings">
                <h1 class="profile-user-name"> _alevecchi</h1>
                <button type="button" class="edit-profile-button">Edit profile</button>
            </div>
            <div class="profile-stats">
                <ul>
                    <li><span class="profile-stat-count">100</span> Posts</li>
                    <li><span class="profile-stat-count">456</span> Followers</li>
                    <li><span class="profile-stat-count">789</span> Following</li>
                </ul>
            </div>
            <div class="profile-bio">
                <!--<h2 class="profile-bio-name">
                {{ profile.name }}
            </h2>-->
                <p class="profile-bio-text">
                    {{ profile.bio }}
                </p>
            </div>
            <!--End of profile section-->
        </div>
    </header>

    <main>
        <div class="wrapper">
            <div class="gallery">
                <div class="gallery-item" tabindex="0">
                    <img src="{{ profile.photo }}" alt="" class="gallery-image">
                    <div class="gallery-item-info">
                        <ul>
                            <li class="gallery-item-likes"><span class="visually-hidden">Likes:</span></li>
                            <li class="gallery-item-comments"><span class="visually-hidden">Comments:</span></li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    </main>
</template>

<style scoped>
header {
    min-height: 100vh;
    padding-bottom: 3rem;
}
.img {
    display: block;
}

.wrapper {
    max-width: 93.5rem;
    margin: 0 auto;
    padding: 0 2rem;
}

.edit-profile-button {
    display: inline-block;
    margin-top: 1rem;
    margin-bottom: 1rem;
    font: inherit;
    padding: 0;
    cursor: pointer;
}

.edit-profile-button:focus {
    outline: 0.5rem auto #4d90fe;
}

.visually-hidden {
    position: absolute !important;
    height: 1px;
    width: 1px;
    overflow: hidden;
    clip: rect(1px, 1px, 1px, 1px);
}

/*Profile section */
.profile {
    padding: 5rem 0;
}

.profile::after {
    display: block;
    content: "";
    clear: both;
}

.profile-image {
    display: flex;
    float: left;
    justify-content: center;
    width: calc(33.33% - 1rem);
    margin-right: 3rem;
}

.profile-image img {
    width: 20vh;
    border-radius: 50%;
    cursor: pointer;
}

.profile-user-name {
    display: inline-block;
    font-size: 3rem;
}

.profile-user-setting,
.profile-bio,
.profile-stats {
    float: right;
    width: calc(66.66%-2rem);
}

.profile-user-setting {
    margin-top: 1.2rem;
}
</style>