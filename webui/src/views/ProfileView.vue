<script>
export default {
    data: function() {
        return {
            loading : false,
            errmsg : null,
            profile: {
                user_id: "",
                username: "",
                pictures_count:0,
                followers_count:0,
                following_count:0,
                bio:"",
                profile_picture:"",

            },
			media:[],
        }
    },
    methods: {
        async GetProfile() {
            this.loading = true;
            this.errormsg = null;
            try {
                this.$axios.get("/users/?username="+this.$route.params.username).then(response => (this.profile = response.data));
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
		async GetMedia() {
            this.loading = true;
            this.errormsg = null;
            try {
                this.$axios.get("/users/:userid="+localStorage.getItem('Authorization')+"/photos/").then(response => (this.media = response.data));
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        createMedia: async function(){
            this.$router.push({ path: '/users/'+this.profile.userid+"/photos/"})
        },
		updateProfile: async function(){
            this.$router.push({ path: '/users/'+this.profile.username+"/updateProfile"})
        },
		changeUsername: async function(){
            this.$router.push({ path: '/users/'+this.profile.username+"/changeUsername"})
        }
    },
    mounted() {
		this.GetMedia();
        this.GetProfile();
    }
}
</script>
<template>
    <div>
        <h1> Profile</h1>
            </div>
			
            <div class="card-body">
                <p class="card-text">
                    id: {{ this.profile.userid }}<br />
                    name: {{ this.profile.username }}<br />
                    bio: {{ this.profile.bio }}<br />
                    media: {{ this.media[0].photo }}<br />
					<img :src=this.media[0].photo>
                    followers: {{ this.profile.nfollowers}}<br />
                    followings: {{ this.profile.nfollowing}}

                </p>
            </div>
            <div>
            <button v-if="!loading" type="button" class="btn btn-primary" @click="createMedia">
                Create new media
            </button>
			<button v-if="!loading" type="button" class="btn btn-primary" @click="updateProfile">
                Update profile
            </button>
			<button v-if="!loading" type="button" class="btn btn-primary" @click="changeUsername">
                change username
            </button>
            <LoadingSpinner v-if="loading"></LoadingSpinner>
        </div>
</template>

<style scoped>
.card {
    margin-bottom: 20px;
}
</style>