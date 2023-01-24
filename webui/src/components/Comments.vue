<script>
import CustomText from "@/components/CustomText.vue"
import Comment from "@/components/Comment.vue"
import { eventBus } from "@/main.js"

export default {
    components: {
        Comment,
        CustomText,
    },
    data: function () {
        return {
            loading: false,
            errormsg: null,
            comments: eventBus.getComments,
            photoId: eventBus.getPhotoId
        }
    },
    methods: {

        goBack() {
            this.$router.go(-1, { params: { photoId: this.photoId } })
        },
        refresh() {
            this.getComments()
        },
        async getComments() {
            this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            try {
                await this.$axios.get("/photos/" + this.photoId + "/comments/").then(response => (this.comments = response.data));
                eventBus.getComments = this.comments
            }
            catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
            console.log("comments:", this.comments)
        },

    },
    mounted() {
        this.refresh()
    }

}
</script>

<template>
    <div class="page">
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <div class="card">

            <div class="nested-comment">
                <CustomText size="xxlarge">Comment section</CustomText> <!-- nested -->
                <span style="float:right">
                    <button type="button">
                        <font-awesome-icon icon="fa-solid fa-xmark" size="2x" color="#666" @click="goBack" />
                    </button>
                </span>
            </div>
            <div class="section-1">
                <Comment class="comment-space" v-on:refresh-parent="refresh" v-for="comm in comments"
                    :key="comm.commentId" :commentId="comm.commentId" :author="comm.author"
                    :profilePic="comm.profile_pic" :image="comm.image" :createdIn="comm.created_in" :body="comm.body"
                    :modifiedIn="comm.modified_in" />
            </div>
            <!-- <div class="section-2">
                <Comment class="comment-space" />
                <Comment class="comment-space" />
            </div>
            <div class="section-1">
                <Comment class="comment-space" />
                <Comment class="comment-space" />
            </div>
            <div class="section-2">
                <Comment class="comment-space" />
                <Comment class="comment-space" />
                <Comment class="comment-space" />
            </div> -->

        </div>
    </div>
</template>


<style scoped>
.page {
    display: flex;
    align-items: center;
    justify-content: center;
    background: linear-gradient(109.5deg, rgb(13, 11, 136) 9.4%, rgb(86, 255, 248) 78.4%);
    /* background: linear-gradient(to right, rgb(242, 112, 156), rgb(255, 148, 114)); */
    height: 100%;
    width: 100%;
}

.card {
    position: relative;
    display: flex;
    padding: 20px;
    flex-direction: column;
    background-color: #fff;
    background-clip: border-box;
    border: 1px solid #d2d2dc;
    border-radius: 11px;
    -webkit-box-shadow: 0px 0px 5px 0px rgb(249, 249, 250);
    -moz-box-shadow: 0px 0px 5px 0px rgba(212, 182, 212, 1);
    box-shadow: 0px 0px 5px 0px rgb(161, 163, 164);
    overflow: auto;
    max-width: 700px;
    ;
}

.nested-comment {
    text-align: center;
    padding-bottom: 8px;
}

.go-back {
    margin-right: 8px;
}

.section-1 {
    padding-left: 4px;
}

.section-2 {
    padding-left: 80px;
}

.comment-space {
    margin-bottom: 8px;
}
</style>