<script>
export default {
    props: ['msg'],
    mounted() {
        const errorContainer = document.querySelector('.error-container');
        errorContainer.classList.add('show-error');
        setTimeout(() => {
            errorContainer.classList.remove('show-error');
        }, 15000); // hide the error message after 10 seconds
    }
}
</script>

<template>
    <div class="error-container">
        <div class="error-message">
            <p>{{ msg }}</p>
        </div>
    </div>
</template>

<style>
.error-message {
    display: none;
    /* initially hidden */
    position: fixed;
    top: 40%;
    left: 50%;
    transform: translate(-50%, -50%);
    height: 100px;
    background: radial-gradient(circle at 10% 20%, rgb(221, 49, 49) 0%, rgb(119, 0, 0) 90%);
    color: #fafafa;
    font-size: 20px;
    padding: 1rem;
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