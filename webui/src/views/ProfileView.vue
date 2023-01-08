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
        <div class="profile">
            <div class="profile-image">
                <!--<img :src="profile.profilePicUrl" alt="">-->
                <img :src="image" alt="Mickey Mouse" />
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
    font-size: 10px;
    height: 33vh;
    background-color: black;
}
/*Profile section */
.profile {
    max-width: 93.5rem;
    position: relative;
}

.profile::after {
    box-sizing: border-box;
}
.profile-image {
    position: absolute;
    top: 25%;
    left: 10%;
    outline: dashed red;
}

.profile-image img {
    width: 15vh;
    border-radius: 50%;
    cursor: pointer;
}


.profile-user-settings {
    position: relative;
    width: 65%;
    top: 30px;
    left: 30%;
    outline: dashed red;
}

.profile-user-name {
    position: absolute;
    left: 10%;
    
    font-size: 3rem;
}
.edit-profile-button {
}

.edit-profile-button:focus {
}

.profile-stats{

}
.profile-bio {

}
.visually-hidden {
}

</style>