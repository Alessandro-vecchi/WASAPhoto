<script>
export default {
    data() {
        return {
            postCaption: '',
            postImage: null,
            media: {},
            previewUrl: null,
            error: null,
            loading: false,
        }
    },
    methods: {
        handleImageUpload(event) {
            this.previewImage(event)
            this.uploadImage(event)
        },
        previewImage(event) {
            let input = event.target;
            if (input.files && input.files[0]) {
                let reader = new FileReader();
                reader.onload = (e) => {
                    this.previewUrl = e.target.result;
                }
                reader.readAsDataURL(input.files[0]);
            }
        },
        uploadImage(event) {
            this.postImage = event.target.files[0];
        },
        async submitPost() {
            this.loading = true;
            this.error = null;
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });

            let formData = new FormData();
            formData.append('caption', this.postCaption);
            formData.append('image', this.postImage);
            console.log(this.$route.params.user_id, this.postCaption)

            try {
                let response = await this.$axios.post('/users/' + this.$route.params.user_id + '/photos/', formData, {
                    headers: {
                        'Content-Type': 'multipart/form-data'
                    }
                });
                this.media = response
                console.log(this.media, response)
                //this.postCaption = response.caption
                this.$router.push({ path: "/users/", query: { username: this.media.username } });
            } catch (e) {
                this.error = e
            }
            this.loading = false;
        },

        goBack() {
            this.$router.push({ path: "/users/", query: { username: this.username } });
        },
    }
}

</script>

<template>
    <div class="upload-post">
        <form>
            <div class="form-group">
                <input type="file" id="post-image" @change="handleImageUpload" accept="image/*">
                <label for="post-image" class="upload-label">
                    <span>Add a photo</span>
                    <font-awesome-icon icon="fa-regular fa-image" size="xl" />
                </label>
                <img id="preview-image" v-if="previewUrl" :src="previewUrl">
            </div>
            <div class="form-group">
                <textarea placeholder="What's on your mind?" id="post-caption" v-model="postCaption"></textarea>
            </div>
            <div class="form-group">
                <button v-if="!loading" type="submit" @click="submitPost">Post</button>
                <button v-if="!loading" type="button" @click="goBack">Cancel</button>
            </div>
        </form>
    </div>

</template>

<style scoped>
.upload-post {
    max-width: 600px;
    margin: auto;
    background-color: #fafafa;
    padding: 20px;
}

.form-group {
    margin-bottom: 10px;
}

#post-caption {
    width: 100%;
    height: 100px;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
    resize: none;
}

#post-image {
    display: none;
    /* Don't wanna see the SELEZIONA FILE button */
}

#preview-image {
    max-width: 400px;
    border: 1px solid #ccc;
    padding: 2px;
    margin: 10px 0;
    margin-left: calc((600px - 400px - 40px - 4px - 2px) / 2); /* max-width parent - max width child - padding parent - padding child - border */
}

.upload-label {
    display: inline-block;
    width: 100%;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
    text-align: center;
    cursor: pointer;
}

.upload-label span {
    margin-right: 10px;
}

button[type='submit'] {
    background-color: #4caf50;
    color: #fff;
    padding: 10px 20px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}

button[type='button'] {
    background-color: #f44336;
    color: #fff;
    padding: 10px 20px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    float: right;
}
</style>