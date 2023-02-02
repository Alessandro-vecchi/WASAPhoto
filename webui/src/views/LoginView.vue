<script>
import { eventBus } from "@/main.js"

export default {
    data: function () {
        return {
            errormsg: null,
            loading: false,
            User: {
                UserID: "",
                Username: "",
            }
        }
    },
    methods: {
        LoginUser: async function () {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios.post("/session/", {
                    username: this.User.Username,
                });
                this.User.UserID = response.data,

                    eventBus.getMyUsername = this.User.Username
                localStorage.clear();
                localStorage.setItem('Authorization', this.User.UserID),

                    this.$router.push({ path: '/users/' + this.User.UserID + '/stream/' })
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        }
    }
}
</script>

<template>
    <div class="background">
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <div>
            <img class="logo" src="../assets/logo2.png" />
        </div>
        <div class="login-container">
            <h1> Welcome </h1>
            <div class="register">
                <font-awesome-icon class="user-icon" icon="fa-solid fa-user" size="xl" />
                <input type="text" v-model="User.Username" placeholder="Enter Username">
                <button v-if="!loading" class="login-button" type="button" @click="LoginUser">LOGIN</button>
                <LoadingSpinner v-if="loading"></LoadingSpinner>
            </div>
        </div>
    </div>




</template>

<style scoped>
.background {
    background-color: rgba(18, 23, 29);
    margin: -10px;
    height: 100vh;
    display: flex;
    flex-direction: column;
    text-align: center;
}

.background div {
    margin-right: auto;
    margin-left: auto;
}

.logo {
    width: 100px;
    margin-top: 50px;
}

.login-container {
    margin-top: 50px;
    width: 300px;
    height: 450px;
    background-color: aliceblue;
    position: relative;
}

.login-container h1 {
    margin-top: 15px;
}

.register input {
    position: absolute;
    top: 66px;
    left: 45px;
    width: 70%;
    height: 40px;
    padding-left: 15px;
    border: 1px solid skyblue;
}

.user-icon {
    position: absolute;
    top: 75px;
    left: 15px;
}

.login-button {
    border-radius: 20px;
    margin-top: 80px;
    width: 70%;
    height: 40px;
    border: 1px solid skyblue;
    background: linear-gradient(109.6deg, rgb(78, 62, 255) 11.2%, rgb(164, 69, 255) 91.1%);

    font-size: 14px;
    font-family: "Rubik", sans-serif;
    font-weight: 400;
    color: white;
    letter-spacing: 4px;
    text-transform: uppercase;
    text-decoration: none;

    transition: 0.5s;
}

.login-button:hover {
    letter-spacing: 8px;
    cursor: pointer;
}
</style>
