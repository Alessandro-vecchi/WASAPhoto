<script>
import { eventBus } from "@/main.js"
export default {
    data() {
        return {
            username: '',
            bio: '',
            avatar: null,
            errormsg: null,
            loading: false,
            isSaved: false,
        }
    },
    methods: {
        uploadAvatar(event) {
            this.avatar = URL.createObjectURL(event.target.files[0])
        },
        async submit() {
            this.isSaved = false
            this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            try {
                let formData = new FormData();
                formData.append('username', this.username);
                formData.append('bio', this.bio);
                formData.append('image', this.$refs.avatar.files[0]);

                eventBus.getMyUsername = this.username
                console.log(this.username, this.avatar, this.bio)
                await this.$axios.put('/users/' + this.$route.params.user_id, formData, {
                    headers: { 'Content-Type': 'multipart/form-data' }
                });
                this.$router.push({ path: "/users/", query: { username: this.username } });
            } catch (e) {
                this.errormsg = e.response.data.error.toString();
            }
            this.loading = false;
            this.isSaved = true
        },
        cancel() {
            if (this.isSaved && !this.errormsg) {
                this.$router.push({ path: "/users/", query: { username: this.username } });
            } else {
                this.$router.push({ path: "/users/", query: { username: eventBus.getMyUsername } });
            }

        },
        async deleteProfile() {
            this.loading = true;
            this.errormsg = null;
            try {
                await this.$axios.delete('/users/' + this.$route.params.user_id);
                this.$router.push({ path: "/login" });
            } catch (e) {
                this.errormsg = e.response.data.error.toString();
            }
            this.loading = false;
        }
    },
    mounted() {
        this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
            error => { return Promise.reject(error); });
        this.$axios.get('/users/?username=').then(response => {
            this.username = response.data.username;
            this.bio = response.data.bio;
            this.avatar = response.data.image;
            eventBus.getMyUsername = response.data.username;
        }).catch(e => this.errormsg = e.response.data.error.toString());
    },
    computed: {
        isDisabled() {
            /* The post button is disabled if:
            1. user didn't upload any image;
            2. the bio is empty;
            3. the username is not changed */
            return !(this.avatar)&&!(this.bio)&&(this.username===eventBus.getMyUsername);
        }
    }
}
</script>


<template>
    <div class="edit-profile">
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <div class="edit-profile-title">
            <h2>Edit your profile</h2>
        </div>
        <!--     The .prevent modifier tells Vue to call event.preventDefault() before the method, which stops the browser's
        default form submission behavior. -->
        <form @submit.prevent="submit">
            <div class="form-group">
                <label for="username">Username</label>
                <input type="text" id="username" name="username" v-model="username">
            </div>
            <div class="form-group">
                <label for="bio">Bio</label>
                <input type="text" id="bio" name="bio" v-model="bio">
            </div>
            <div class="form-group">
                <label for="avatar">Avatar</label>
                <input type="file" id="avatar" name="avatar" ref="avatar" @change="uploadAvatar">
                <img v-if="avatar" :src="avatar" alt="Avatar">
            </div>
            <div class="form-group">
                <button v-if="!loading" :disabled="isDisabled" type="submit">Save</button>
                <button v-if="!loading" type="go-back" @click="cancel">Cancel</button>
                <button v-if="!loading" type="delete" @click="deleteProfile">Delete Profile</button>
            </div>
        </form>
    </div>

</template>

<style scoped>
.edit-profile {
    max-width: 600px;
    margin: auto;
    background-color: #fafafa;
}

.edit-profile-title {
    margin-top: 20px;
    margin-bottom: 25px;
}

.edit-profile-title h2 {
    /* 
    font-style: italic; */
    font-family: Verdana, Geneva, Tahoma, sans-serif;
}

.form-group {
    margin-bottom: 20px;
}

.form-group label {
    display: block;
    font-size: 14px;
    font-weight: bold;
    margin-bottom: 8px;
}

.form-group input[type="text"] {
    width: 100%;
    padding: 12px;
    border-radius: 4px;
    border: 1px solid #ccc;
    font-size: 14px;
    box-sizing: border-box;
}

.form-group img {
    max-width: 200px;
    margin-top: 8px;
}

.form-group button {
    color: white;
    padding: 12px 20px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}

.form-group button[type="submit"] {
    background-color: #4CAF50;
}

.form-group button[type="go-back"] {
    background-color: #f44336;
}

.form-group button[type="delete"] {
    background-color: #940e0e;
    float: right;
}
</style>