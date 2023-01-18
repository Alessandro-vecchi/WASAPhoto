<script>
import Avatar from "@/components/Avatar.vue"
import CustomText from "@/components/CustomText.vue"
import ShortProfile from "@/components/ShortProfile.vue"
import { eventBus } from "@/main.js"

export default {
    components: {
        Avatar,
        CustomText,
        ShortProfile,
    },
    data: function () {
        return {
            short_profiles: eventBus.getShortProfiles,
            title: eventBus.getTitle,
            username: eventBus.getUsername,
            header: localStorage.getItem('Authorization'),
            ppUrl: "",
        }
    },
    methods: {

        goBack(isLikeList) {
            if (isLikeList=="LIKES") {
                this.$router.push({ path: "/users/" + this.header + "/stream/" });
            } else { this.$router.push({ path: "/users/", query: { username: this.username } }) }
        },
    },
    mounted() {
        
    }


}
</script>

<template>
    <div class="container">
        <div class="head section">
            <div class="list-title">{{ title }}</div>
            <div class="header-more">
                <button type="button">
                    <font-awesome-icon icon="fa-solid fa-xmark" size="2x" @click="goBack(title)" />
                </button>
            </div>
        </div>
        <div class="content section">
            <ShortProfile v-for="s_p in short_profiles" :key="s_p.username" :shortProfile="s_p" />
        </div>
    </div>
</template>

<style>
:root {
    --background-likes: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
    --background-header-likes: linear-gradient(112.1deg, rgb(32, 38, 57) 11.4%, rgb(63, 76, 119) 70.2%);
}

.container {
    max-width: 300px;
    position: absolute;
    top: 40%;
    left: 50%;
    transform: translate(-50%, -50%);
}

.container .section {
    padding-left: 16px;
    padding-right: 16px;
}

.container .head {
    display: flex;
    align-items: center;
    width: 300px;
    height: 50px;
    background: var(--background-header-likes);
}

.container .head .list-title {
    margin-left: 40%;
    font-size: 16px;
    font-weight: 600;
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
    text-transform: uppercase;
    color: #f5f7fa;
}

.container .head .header-more {
    margin-left: auto;
}

.container .content {
    width: 300px;
    height: 300px;
    border-radius: 0 0 20px 20px;
    background: var(--background-likes);
    overflow: auto;
}
</style>