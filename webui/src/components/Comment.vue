<script>
import Avatar from "@/components/Avatar.vue"
import CustomText from "@/components/CustomText.vue"
import { eventBus } from "@/main.js"

export default {
    props: {
        commentId: String,
        author: String,
        profilePic: String,
        image: String,
        createdIn: String,
        body: String,
        modifiedIn: String,

    },
    components: {
        Avatar,
        CustomText,
    },
    data: function () {
        return {
            loading: false,
            errormsg: null,
            pp: "",
            myUsername: eventBus.getMyUsername,
            currentBody: this.body,
            modified_in: this.modifiedIn,
            editing: false,
        }
    },
    methods: {
        toggleEditing() {
            this.editing = !this.editing;
        },
        async getImage() {
            this.loading = true;
            this.errormsg = null;
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            try {
                let response = await this.$axios.get("/images/?image_name=" + this.profilePic, { responseType: 'blob' })
                // Get the image data as a Blob object
                var imgBlob = response.data;

                // Create an object URL from the Blob object
                this.pp = URL.createObjectURL(imgBlob);
            } catch (error) {
                // console.log(error);
                this.errormsg = error.message;
            }
            this.loading = false;
        },
        async uncommentPhoto() {
            this.loading = true;
            this.errormsg = null;
            try {
                await this.$axios.delete('/comments/' + this.commentId);
            } catch (error) {
                this.errormsg = error;
            }
            this.loading = false;
            this.$emit('refresh-parent');

        },
        async modifyComment() {
            this.loading = true;
            this.errormsg = null;

            const data = JSON.stringify({
                body: this.currentBody,
                author: this.myUsername,
            })
            this.$axios.interceptors.request.use(config => { config.headers['Authorization'] = localStorage.getItem('Authorization'); return config; },
                error => { return Promise.reject(error); });
            try {
                let response = await this.$axios.patch('/comments/' + this.commentId, data, {
                    headers: { 'Content-Type': 'application/json' }
                });
                this.currentBody = response.data.body
                this.modified_in = response.data.modified_in
            } catch (e) {
                this.errormsg = e;
            }
            this.loading = false;
            this.toggleEditing()
        },
        
    },
    computed: {
        timeAgo() {
           /*  var dateString
            if (this.edited) {
                dateString = this.modified_in;
            } else {
                dateString = this.createdIn;
            } */
            var dateString = this.createdIn;
            var date = new Date(dateString);
            var year = date.getFullYear();
            var month = date.getMonth(); // getMonth() returns a number between 0 and 11
            var day = date.getDate();
            var hours = date.getHours();
            var minutes = date.getMinutes();
            var seconds = date.getSeconds();
            console.log("Timestamp Year: " + year + " Month: " + month + " Day: " + day);
            console.log("Timestamp Hours: " + hours + " Minutes: " + minutes + " Seconds: " + seconds);

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
            console.log("Author:", this.author, "Username:", this.myUsername)
            return (this.author === this.myUsername)
        },
        edited() {
            console.log("created:", this.createdIn, "modified:", this.modified_in)
            return this.createdIn !== this.modified_in
        },

    },
    mounted() {
        console.log("mounted:", eventBus.getMyUsername)
        if (this.profilePic) {
            this.getImage()
        }
    },
}
</script>

<template>
    <div class="media">
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <div class="image-place">
            <Avatar :src="pp" :size="60" />
        </div>
        <div class="media-body">
            <div>
                <CustomText size="large" tag="b">{{ author }}</CustomText> 

                <span class="time">
                    <CustomText size="xsmall" class="time-ago">{{ timeAgo }}</CustomText>
                    <CustomText v-if="edited"  size="xxsmall">  (modified)</CustomText>
                </span>
                <span>
                    <font-awesome-icon icon="fa-solid fa-reply" pull="right" color="rgb(14,115,248)" size="xl" />
                </span>
            </div>

            <div class="comment-body">
                <textarea v-if="editing" v-model="currentBody"></textarea>
                <CustomText v-else size="normal">{{ currentBody }}</CustomText>
            </div>


            <div class="buttons">
                <button v-if="isMine" type="edit" @click="toggleEditing">Edit</button>
                <button v-if="isMine" type="button" @click="uncommentPhoto">Delete</button>
                <button v-if="editing" type="submit" @click="modifyComment">Save</button>
            </div>
        </div>
    </div>
</template>

                <!-- It's known that the majority have suffered alteration in some form, by injected humour, or randomised
                words. I wonder what happens if I write anothe thing. Interesting, so it keeps increasing till we reach the end of the line. what happens now? -->
<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Source+Sans+Pro:wght@300&display=swap');
@import url('https://fonts.googleapis.com/css?family=Proxima+Nova');

.media {
    font-family: 'Proxima Nova', sans-serif;
    width: inherit;
    background: #fafafa;
    border-radius: 20px;
    display: flex;
    align-items: center;
}

.media-body {
    margin-left: 15px;
}

.image-place {
    min-width: 60px;
}

.time {
    margin-left: 10px;
}

.time-ago {
    color: rgba(100, 100, 100, 1);
    text-transform: uppercase;
}

.comment-body {
    color: black;
    width: fit-content;
    margin-top: 4px;
    font-family: 'Source Sans Pro', sans-serif;
}

.buttons button {
    color: white;
    padding: 6px 10px;
    margin-top: 10px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}

.buttons button[type="edit"] {
    background-color: #31b4d5;
}

.buttons button[type="button"] {
    margin-left: 8px;
    background-color: #9d2121;
}

.buttons button[type="button"].active {
    background-color: #770707;
}
.buttons button[type="submit"] {
    background-color: #2bb148;
    float: right;
}
</style>