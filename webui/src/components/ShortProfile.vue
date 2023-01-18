<template>
    <div class="short-profile">
        <!-- <Avatar :src="s_p.profilePictureUrl" :size="40" class="profile-photo" /> -->
        <Avatar v-if="sp" :src="sp" :size="40" class="profile-photo" />
        <Avatar v-else :size="40" class="profile-photo" />
        <div class="author-username">
            <CustomText tag="b">{{ shortProfile.username }}</CustomText>
        </div>
    </div>
</template>

<script>
import Avatar from "@/components/Avatar.vue"
import CustomText from "@/components/CustomText.vue"

export default {
    props: ['shortProfile'],
    components: {
        Avatar,
        CustomText,
    },
    data: function () {
        return {
            loading: false,
            errormsg: null,
            sp: "",
        }

    },
    methods: {
        async getImage() {
            console.log("1", this.shortProfile, "2", this.shortProfile.profilePictureUrl)
            this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            try {
                let response = await this.$axios.get("/images/?image_name=" + this.shortProfile.profilePictureUrl, { responseType: 'blob' })
                // Get the image data as a Blob object
                var imgBlob = response.data;

                // Create an object URL from the Blob object
                this.sp = URL.createObjectURL(imgBlob);
            } catch (error) {
                // console.log(error);
                this.errormsg = error.message;

            }
            this.loading = false;
        },
    },
    mounted() {
        if (this.shortProfile.profilePictureUrl) {
            this.getImage()
        }
    },
}
</script>

<style scoped>
.short-profile {
    display: flex;
    flex-direction: rows;
    margin-top: 4px;
}

.short-profile .profile-photo:hover {
    cursor: pointer;
}

.short-profile .author-username {
    align-items: center;
    margin-left: 8px;
    font-size: 16px;
    border-bottom: 1px solid #999;
}

.short-profile .author-username b:hover {
    text-decoration: underline;
    cursor: pointer;
}
</style>