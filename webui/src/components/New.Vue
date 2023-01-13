<script>
export default {
    data: function () {
        return {
            errormsg: null,
            loading: false,
            profile: {},
            file: "ciao"
        }
    },
}
</script>

<template>
    <h1>Buuuu</h1>
    <div
        class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
        <h1 class="h2">Update profile</h1>
    </div>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

    <div class="mb-3">
        <label for="description" class="form-label">Username</label>
        <input type="text" class="form-control" id="Username" v-model="profile.username"
            placeholder=this.profile.username>
    </div>
    <div class="mb-3">
        <label for="description" class="form-label">Bio</label>
        <input type="text" class="form-control" id="bio" v-model="profile.bio" placeholder=this.profile.bio>
    </div>

    <label for="description" class="form-label">profile</label>
    <div class="container">
        <form @submit.prevent="onSubmit">
            <div class="form-group">
                <input type="file" @change="onFileUpload">
            </div>
            <div class="form-group">
                <button class="btn btn-primary btn-block btn-lg">Upload File</button>
            </div>
        </form>
    </div>
</template>

<style>

</style>