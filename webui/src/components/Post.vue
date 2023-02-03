<script>
import Avatar from "@/components/Avatar.vue"
import CustomText from "@/components/CustomText.vue"
import { eventBus } from "@/main.js"

export default {
    props: {
        photoId: String,
        owner: String,
        profilePictureUrl: String,
        image: String,
        timestamp: String,
        caption: String,
        likesCount: Number,
        commentsCount: Number,

    },
    components: {
        Avatar,
        CustomText,
    },
    data: function () {
        return {
            header: localStorage.getItem('Authorization'),
            loading: false,
            errormsg: null,
            username: "",
            textComment: "",
            isLiked: null,
            myProfilePic: "",
            likes: [],
            comments: [],
            ppUrl: "",
            imgUrl: "",
        }
    },
    methods: {
        async get_user_profile(name) {
            this.$router.push({ path: "/users/", query: { username: name } })
        },
        async Get_my_profile() {
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            console.log("GetMyProfile")
            return this.$axios.get("/users/?username=" + this.username).then(response => (this.username = response.data.username, this.myProfilePic = response.data.profile_picture_url, eventBus.getMyUsername = response.data.username)).catch(e => (this.errormsg = e.response.data.error.toString() ));
        },
        async LikeClick() {
            if (this.isMine) {
                return
            }
            this.loading = true;
            this.errormsg = null;
            /* The interceptor is modifying the headers of the requests being sent by adding an 'Authorization' header with a value that is stored in the browser's local storage. Just keeping the AuthToken in the header.
            If you don't use this interceptor, the 'Authorization' header with the token won't be added to the requests being sent, it can cause the requests to fail.
            */
            console.log(this.isLiked)
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            if (this.isLiked) {
                await this.$axios.delete("/photos/" + this.photoId + "/likes/" + this.header);
            } else {
                await this.$axios.put("/photos/" + this.photoId + "/likes/" + this.header)//.then(response => (this.username = response.data.name));
            }
            this.loading = false;
            this.refresh()
            this.$emit('refresh-parent');
        },
        async GetLikes(isRefresh) {
            this.loading = true;
            this.errormsg = null;
            /* The interceptor is modifying the headers of the requests being sent by adding an 'Authorization' header with a value that is stored in the browser's local storage. Just keeping the AuthToken in the header.
            If you don't use this interceptor, the 'Authorization' header with the token won't be added to the requests being sent, it can cause the requests to fail.
            */
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            try {
                let response = await this.$axios.get("/photos/" + this.photoId + "/likes/")
                this.likes = response.data.short_profile
                this.isLiked = response.data.cond
                if (!isRefresh) {
                    eventBus.getShortProfiles = this.likes
                    eventBus.getTitle = "LIKES"
                    this.$router.push({ path: '/likes/' })

                }
            } catch (e) {
                this.errormsg = e.response.data.error.toString();
            }
            this.loading = false;
            console.log("likes:", this.likes, this.isLiked)
            if (!isRefresh) {
                this.refresh()
            }
        },

        async GetComments(isRefresh) {
            this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            try {
                await this.$axios.get("/photos/" + this.photoId + "/comments/").then(response => (this.comments = response.data));
                if (!isRefresh) {
                    eventBus.getComments = this.comments
                    eventBus.getPhotoId = this.photoId
                    this.$router.push({ path: '/photos/' + this.photoId + "/comments/" })
                }
            } catch (e) {
                this.errormsg = e.response.data.error.toString();
            }
            this.loading = false;
            console.log("comments:", this.comments)
            if (!isRefresh) {
                this.refresh()
            }
        },
        async submitComment() {
            this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });

            try {
                let response = await this.$axios.post('/photos/' + this.photoId + '/comments/', {
                    body: this.textComment, isReplyComment: false, author: this.username,
                });
                console.log(response.data)
                this.GetComments(false)
                // this.$router.push({ path: '/photos/' + this.photoId + "/comments/" });
            } catch (e) {
                this.errormsg = e
            }
            this.loading = false;
        },
        async GetImage(url, filter) {
            console.log("url:", url, "filter:", filter)
            this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });

            await this.$axios.get("/images/?image_name=" + url, { responseType: 'blob' }).then(response => {
                // Get the image data as a Blob object
                var imgBlob = response.data;

                // Create an object URL from the Blob object
                var uri = URL.createObjectURL(imgBlob);
                if (filter === "my") {
                    this.myProfilePic = uri
                } else if (filter === "img") {
                    this.imgUrl = uri
                } else if (filter === "pp") {
                    this.ppUrl = uri
                }
                console.log("uri:", uri)
            })
                .catch(e => {
                this.errormsg = e.response.data.error.toString();
                });
            this.loading = false;
        },
        async getImages() {
            
            console.log("GetImages")
            console.log("MyPP")
            if (this.myProfilePic) {
                await this.GetImage(this.myProfilePic, "my")
            }
            console.log("Image")
            if (this.image) {
                await this.GetImage(this.image, "img")
            }
            console.log("commentPP")
            if (this.profilePictureUrl) {
                await this.GetImage(this.profilePictureUrl, "pp")
            }
        },
        async refresh() {
           await this.GetLikes(true).then(() => this.GetComments(true))
        },

        async deletePhoto() {
            this.loading = true;
            this.errormsg = null;
            try {
                await this.$axios.delete('/photos/' + this.photoId);
                this.$router.push({ path: "/users/", query: { username: this.username } })
            } catch (e) {
                this.errormsg = e.response.data.error.toString();
            }
            this.loading = false;

        }
    },
    computed: {
        timeAgo() {
            var dateString = this.timestamp;
            var date = new Date(dateString);
            var year = date.getFullYear();
            var month = date.getMonth(); // getMonth() returns a number between 0 and 11
            var day = date.getDate();
            var hours = date.getHours();
            var minutes = date.getMinutes();
            var seconds = date.getSeconds();
            /* console.log("Timestamp Year: " + year + " Month: " + month + " Day: " + day);
            console.log("Timestamp Hours: " + hours + " Minutes: " + minutes + " Seconds: " + seconds); */

            var currentDate = new Date();
            var c_year = currentDate.getFullYear();// current year
            var c_month = currentDate.getMonth();
            var c_day = currentDate.getDate();
            var c_hours = currentDate.getHours();
            var c_minutes = currentDate.getMinutes();
            var c_seconds = currentDate.getSeconds();
            /* console.log("Current Year: " + c_year + " Month: " + c_month + " Day: " + c_day)
            console.log("Current Hours:" + c_hours + " Minutes: " + c_minutes + " Seconds: " + c_seconds); */

            var timeAgo = "";
            var diffYear = c_year - year;
            var diffMonth = c_month - month;
            var diffDay = c_day - day;
            var diffHour = c_hours - hours;
            var diffMinutes = c_minutes - minutes;
            var diffSeconds = c_seconds - seconds;

            if (diffYear !== 0) {
                timeAgo = diffYear + " years ago";
            } else if (diffMonth !== 0) {
                timeAgo = diffMonth + " months ago";
            } else if (diffDay !== 0) {
                timeAgo = diffDay + " days ago";
            } else if (diffHour !== 0) {
                timeAgo = diffHour + " hours ago";
            } else if (diffMinutes !== 0) {
                timeAgo = diffMinutes + " minutes ago";
            } else if (diffSeconds !== 0) {
                timeAgo = diffSeconds + " seconds ago";
            } else {
                timeAgo = "Just now";
            }
            /* console.log(timeAgo); */
            return timeAgo;
        },
        isMine() {
            console.log("Owner:", this.owner, "Username:", this.username)
            return (this.owner === this.username)
        }

    },
    mounted() {
        this.Get_my_profile().then(() => this.getImages()).then(() => this.refresh())
    }
}
</script>

<template>
    <div class="post">
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <!-- header -->
        <header class="header section">
            <div class="header-author">
                <Avatar :src="ppUrl" :size="45" @click="get_user_profile(owner)" />
                <div class="header-author-info">
                    <CustomText tag="b">{{ owner }}</CustomText> <!-- _alevecchi -->
                </div>
            </div>
            <div class="header-more">
                <button v-if=!isMine type="button">
                    <font-awesome-icon icon="fa-solid fa-ellipsis" size="3x" />
                </button>
                <button v-else type="delete" @click="deletePhoto">Delete Photo</button>
            </div>
        </header>

        <!-- media -->
        <div class="post-media">
            <img :src="imgUrl" alt="" class="post-image" />
            <!-- src="https://picsum.photos/600/400?random=1" -->
        </div>

        <div class="two-col section">
            <!-- action & count-->
            <div class="action-buttons">
                <ul>
                    <li>
                        <button v-if=!loading type="button">
                            <font-awesome-icon v-if=!isLiked class="icon" id="like" icon="fa-regular fa-heart"
                                @click="LikeClick" />
                            <font-awesome-icon v-else class="icon" id="like" icon="fa-solid fa-heart"
                                color="rgb(232, 62, 79)" @click="LikeClick" />
                            <span class="num" @click="GetLikes(false)"> {{ likesCount }} </span>
                        </button>
                    </li>
                    <li>
                        <button v-if=!loading type="button" @click="GetComments(false)">
                            <font-awesome-icon class="icon" id="comment" icon="fa-regular fa-comment" />
                            <span class="num"> {{ commentsCount }} </span>
                        </button>
                    </li>
                </ul>
            </div>

            <div class="caption">
                <li>
                    <CustomText tag="b" @click="get_user_profile(owner)">{{ owner }}</CustomText>
                    <span class="caption-span">{{ caption }}</span>
                </li>
            </div>
        </div>

        <div class="comments-list">
            <!-- datetime-->
            <div class="time section">
                <CustomText size="xxsmall" class="time-ago">{{ timeAgo }}</CustomText>
            </div>

            <!-- comments form -->
            <div class="comment section">
                <Avatar :src="myProfilePic" :size="30" @click="get_user_profile(username)" />
                <input class="text-body" type="text" placeholder="Add a comment..." v-model="textComment">
                <button v-if="!loading" type="submit" @click="submitComment">Post</button>
            </div>
        </div>
    </div>
</template>


<style scoped>
.post {
    border-radius: 3px;
    border: 1px solid rgba(219, 219, 219, 1);
    max-width: 604px !important;
    margin-bottom: 60px;
}

.post .section {
    padding-left: 16px;
    padding-right: 16px;
}

.post .header {
    display: flex;
    align-items: center;
    height: 60px;
}

.post .header-author {
    display: flex;
    align-items: center;
}

.post .header-author-info {
    margin-left: 8px;
    font-size: 16px;
}

.post .header-more {
    margin-left: auto;
}

.post .header-more button[type="delete"] {
    color: white;
    padding: 6px 10px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    background-color: #911b1b;
}

.post .post-media {
    width: 600px;
    height: 400px;
}

.post .post-media .post-image {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.post .two-col {
    display: flex;
    /*outline: solid;*/
}

.post .action-buttons {
    margin-top: 0px;
    height: 60px;
}

.post .action-buttons button {
    padding-top: 6px;
    display: flex;
    align-items: center;
    /*outline: solid;*/
}

.post .action-buttons .num {
    display: inline;
    align-items: center;
    padding-left: 8px;
    font-size: 15px;
    font-weight: 600;
    font-family: Georgia, 'Times New Roman', Times, serif;
    color: #333;
}

.post .action-buttons .icon {
    height: 25px;
    width: 25px;
}

#like:hover {
    color: #555;
    /* background-color: rgb(232, 62, 79) */
}


#comment:hover {
    color: #555
}

.post .caption {
    flex-wrap: wrap;
    margin-top: 9px;
    margin-left: 10rem;
}

.post .caption li b:hover {
    text-decoration: underline;
    cursor: pointer;
}

.post .caption li .caption-span {
    margin-left: 4px;
    overflow: auto;
}

.post .comments-list .time {
    margin-top: 8px;
}

.post .comments-list .time-ago {
    color: rgba(142, 142, 142, 1);
    text-transform: uppercase;
}

.post .comments-list .comment {
    max-width: inherit;
    border-top: 1px solid #efefef;
    margin-top: 4px;
    display: flex;
    height: 55px;
    align-items: center;
}

.post .comments-list .comment .text-body {
    margin-left: 16px;
}

.post .comments-list .comment input {
    flex: 1;
}

.post .comments-list .comment input:focus {
    outline: none;
}

.post .comments-list .comment input::placeholder {
    background-color: -internal-light-dark(rgb(255, 255, 255), rgb(59, 59, 59));
}

.post .comments-list .comment button[type="submit"] {
    background-color: #fafafa;
    margin-left: 16px;
    font-size: 16px;
    color: rgba(0, 160, 230, 1);
}

.post .comments-list .comment button[type="submit"]:hover {
    text-decoration: underline;
    cursor: pointer;
}
</style>