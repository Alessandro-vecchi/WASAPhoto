<script>
export default {
    data: function() {
        return {
            errormsg: null,
            detailedmsg: null,
            loading: false,
            User: {
                UserID: null,
                Username: null,
            }
        }
    },
    methods: {
        LoginUser: async function () {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios.post("/session/", {
                    username: this.Username,
                });
				this.UserID  = response.data,
                localStorage.setItem('Authorization', this.UserID),
                this.$router.push({ name: "stream" })
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
        <div>
            <img class="logo" src="../assets/logo2.png"/>
        </div>
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <div class="login-container">
            <h1> Welcome </h1>
            <div class="register">
                <font-awesome-icon class="user-icon" icon="fa-solid fa-user" size="xl"/>
                <input type="text" v-model="Username" placeholder="Enter Username">
                <button v-if="!loading" type="button" @click="LoginUser">LOGIN</button>
                <LoadingSpinner v-if="loading"></LoadingSpinner>
            </div>
        </div>
    </div>

       
        

</template>

<style>
.background {
  background-color: rgba(18, 23, 29);
  margin: -10px;
  height: 100vh;
  display: flex;
  flex-direction: column;
}
.background div{
    margin-right: auto;
    margin-left: auto;
}
.logo {
    width: 100px;
    margin-top: 50px;
}
.register input{
    position: relative;
    width: 70%;
    height: 40px;
    padding-left: 20px; 
    border: 1px solid skyblue;
}
.register button{
    width: 70%;
    height: 40px;
    border: 1px solid skyblue;
    background: linear-gradient(109.6deg, rgb(78, 62, 255) 11.2%, rgb(164, 69, 255) 91.1%);
    color: white;
    border-radius: 20px;
    cursor: pointer;
    margin-top: 80px;
    font-family: "Rubik", sans-serif;
    font-weight: 400;
    color: white;
    text-decoration: none;
}

.login-container {
    margin-top: 50px;
    width: 300px;
    height: 450px;
    background-color: aliceblue;
}

.user-icon {
    position: relative;
    top: 2px;
    left: -10px;
}
</style>