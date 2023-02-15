<script>
export default {
    props: ['wmsg'],
    mounted() {
        const errorContainer = document.querySelector('.success-container');
        errorContainer.classList.add('show-success');
        setTimeout(() => {
            errorContainer.classList.remove('show-success');
        }, 5000); // hide the welcome message after 10 seconds
    },
}
</script>

<template>
    <div class="success-container">
        <div class="success-message">
            <p>{{ wmsg }}</p>
        </div>
    </div>
</template>

<style scoped>
.success-message {
    display: none;
    /* initially hidden */
    background-image: url("../assets/successMessage.png");
    /* background: radial-gradient(circle at 10% 20%, rgb(221, 49, 49) 0%, rgb(119, 0, 0) 90%); */
    background-color: rgb(22, 183, 22);
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

/* show the error message when it's parent container has class "show-success" */
.show-success .success-message {
    display: block;
}
</style>