<script>
export default {
    props: ['msg'],
    mounted() {
        const errorContainer = document.querySelector('.error-container');
        errorContainer.classList.add('show-error');
        setTimeout(() => {
            errorContainer.classList.remove('show-error');
        }, 10000); // hide the error message after 10 seconds
        console.log(this.$props)
    },
    computed: {
        message() {
            if (this.msg === "404") {
                return `404 user not found.`// Check for typo in your search or try to hack our database and add it by yourself!`
            }
            return this.msg
        }
    }
}
</script>

<template>
    <div class="error-container">
        <div class="error-message">
            <p>{{ message }}</p>
        </div>
    </div>
</template>

<style scoped>
.error-message {
    display: none;
    /* initially hidden */
    background-image: url("../assets/errorMessage.png");
    /* background: radial-gradient(circle at 10% 20%, rgb(221, 49, 49) 0%, rgb(119, 0, 0) 90%); */
    background-color: rgb(130, 20, 20);
    background-position: 10px center;
    background-repeat: no-repeat;

    position: fixed;
    top: 40%;
    left: 50%;
    transform: translate(-50%, -50%);

    color: #fafafa;
    font-size: 20px;

    border: 1px solid;
    margin: 10px 0px;
    padding: 15px 10px 15px 50px;
    border-radius: 5px;
    box-shadow: 2px 2px 8px rgba(0, 0, 0, 0.2);
    
    z-index: 999;
    /* to ensure it appears on top of other elements */
}

/* show the error message when it's parent container has class "show-error" */
.show-error .error-message {
    display: block;
}
</style>