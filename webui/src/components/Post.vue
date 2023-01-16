<script>
import Avatar from "@/components/Avatar.vue"
import CustomText from "@/components/CustomText.vue"
export default {
    props: ['post'],
    name: 'Post',
    components: {
        Avatar,
        CustomText,
    },
    data: function () {
        return {
        }
    },
    methods: {
        async get_user_profile() {
            this.$router.push({ path: "/users/", query: { username: this.post.username } })
        },
    },
    computed: {
        timeAgo() {
            var dateString = this.post.timestamp;
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

    },
    mounted() {
        this.timeAgo
        console.log(this.post.profile_pic)
    }
}
</script>

<template>
    <div class="post">
        <!-- header -->
        <header class="header section">
            <div class="header-author">
                <Avatar :src="post.profile_pic" :size="45" @click="get_user_profile()"/> <!-- :src= "post.profile_pic" -->
                <div class="header-author-info">
                    <CustomText tag="b" >{{ post.username }}</CustomText> <!-- _alevecchi -->
                </div>
            </div>
            <div class="header-more">
                <button type="button">
                    <font-awesome-icon icon="fa-solid fa-ellipsis" size="3x" />
                </button>
            </div>
        </header>

        <!-- media -->
        <div class="post-media">
            <img :src="post.image" alt="" class="post-image" /> <!-- src="https://picsum.photos/600/400?random=1" -->
        </div>

        <div class="two-col section">
            <!-- action & count-->
            <div class="action-buttons">
                <ul>
                    <li>
                        <button type="button">
                            <font-awesome-icon class="icon" id="like" icon="fa-regular fa-heart" size="2x" />
                            <span class="num"> {{ post.likes_count }} </span> <!-- {{ post.likes_count }} -->
                        </button>
                    </li>
                    <li>
                        <button type="button">
                            <font-awesome-icon class="icon" id="comment" icon="fa-regular fa-comment" size="2x" />
                            <span class="num"> {{ post.comments_count }} </span> <!-- {{ post.comments_count }} -->
                        </button>
                    </li>
                </ul>
            </div>

            <div class="caption">
                <li>
                    <CustomText tag="b" @click="get_user_profile()">{{ post.username }}</CustomText> <!-- {{ post.username }} -->
                    <span class="caption-span">{{ post.caption }}</span> <!-- {{ post.caption }} -->
                </li>
            </div>
        </div>

        <div class="comments-list">
            <!-- datetime-->
            <div class="time section">
                <CustomText size="xxsmall" class="time-ago">{{ timeAgo }}</CustomText> <!-- {{ post.timestamp }} -->
            </div>

            <!-- comments form -->
            <div class="comment section">
                <Avatar :src="post.profile_pic" :size="30" @click="get_user_profile()"/>
                <input class="text-body" type="text" placeholder="Add a comment...">
                <a href="#" type="button">Post</a>
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

.post .comments-list .comment a {
    margin-left: 16px;
    font-size: 16px;
    color: rgba(0, 160, 230, 1);
}

.post .comments-list .comment a:hover {
    text-decoration: underline;
    cursor: pointer;
}
</style>