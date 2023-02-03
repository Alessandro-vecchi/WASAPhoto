<script>
import { eventBus } from "@/main.js"

export default {
    data() {
        return {
            username: '',
            errormsg: null,
            loading: false,
            isSaved: false
        }
    },
    methods: {
        async submit() {
            this.isSaved=false
            this.loading = true;
            this.errormsg = null;

            const data = JSON.stringify({
                username: this.username,
            })
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            console.log(this.$route.params.user_id)
            try {
                let response = await this.$axios.patch('/users/' + this.$route.params.user_id, data, {
                    headers: { 'Content-Type': 'application/json' }
                });
                this.username = response.data.username
                console.log(response.data, response, this.username)
                eventBus.getMyUsername = this.username
                this.$router.push({ path: "/users/", query: { username: this.username } });
            } catch (e) {
                this.errormsg = e.response.data.error.toString();
            }
            this.loading = false;
            this.isSaved=true
        },
        cancel() {
            if (this.isSaved&&!this.errormsg) {
                this.$router.push({ path: "/users/", query: { username: this.username } });
            } else {
                this.$router.push({ path: "/users/", query: { username: eventBus.getMyUsername } });
            }
        },
    },
    mounted() {
        this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
            error => { return Promise.reject(error); });
        this.$axios.get('/users/?username=').then(response => {
            this.username = response.data.username;
            eventBus.getMyUsername = response.data.username;
        });
    }
}
</script>


<template>
    <div class="edit-profile">
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <div class="form-group">
            <label for="username">Username</label>
            <input type="text" id="username" name="username" v-model="username">
        </div>
        <div class="form-group">
            <button v-if="!loading" type="submit" @click="submit">Save</button>
            <button v-if="!loading" type="go-back" @click="cancel">Cancel</button>
        </div>
    </div>

</template>

<style scoped>
.edit-profile {
    max-width: 600px;
    margin: auto;
    background-color: #fafafa;
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
</style>