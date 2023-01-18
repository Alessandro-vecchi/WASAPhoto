<script>
export default {
    props: ['photo'],
    name: 'GalleryItem',
    data: function () {
        return {
            loading: false,
            errormsg: null,
            imgUrl: "",
        }

    },
    methods: {
        async getImage() {
            console.log("1", this.photo, "2", this.photo.image)
            this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            try {
                let response = await this.$axios.get("/images/?image_name=" + this.photo.image, { responseType: 'blob' })
                // Get the image data as a Blob object
                var imgBlob = response.data;

                // Create an object URL from the Blob object
                this.imgUrl = URL.createObjectURL(imgBlob);
            } catch (error) {
                // console.log(error);
                this.errormsg = error.message;

            }
            this.loading = false;
        },
    },
    mounted() {
        if (this.photo.image) {
            this.getImage()
        }
    },
}
</script>

<template>
    <div class="gallery-item" tabindex="0">
        <img v-if=!loading :src="imgUrl" alt="" class="gallery-image"> <!-- :src="'data:image/jpg;base64'+imageUrl" -->
        <div class="gallery-item-info">
            <ul>
                <li class="gallery-item-likes"><span class="visually-hidden">Likes:</span><font-awesome-icon
                        icon="fa-solid fa-heart" size="xl" /> {{ photo.likes_count }}</li>
                <li><span class="visually-hidden">Comments:</span><font-awesome-icon icon="fa-solid fa-comment"
                        size="xl" /> {{ photo.comments_count }}</li>
            </ul>
        </div>
    </div>
</template>

<style>
.gallery-item {
    /* max-width: calc((93.5rem - 2rem - 4rem)/3); */
    position: relative;
    flex-grow: 1;
    margin: 1rem;
    color: #fafafa;
    cursor: pointer;
}

.gallery-item:hover .gallery-item-info,
.gallery-item:focus .gallery-item-info {
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    top: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.3)
}

.gallery-item-info {
    display: none;
}

.gallery-item-info li {
    display: inline-block;
    font-size: 1.3rem;
    font-weight: 600;
}

.gallery-item-likes {
    margin-right: 1.5rem;
}

/* .gallery-image {
    max-width: 600px;
    max-height: 400px;
    object-fit: cover;
} */
.gallery-image {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.visually-hidden {
    position: absolute;
    height: 1px;
    width: 1px;
    overflow: hidden;
    clip: rect(1px, 1px, 1px, 1px);
}

@supports (display:grid) {

    .gallery-item {
        width: auto;
        margin: 0;

    }
}
</style>