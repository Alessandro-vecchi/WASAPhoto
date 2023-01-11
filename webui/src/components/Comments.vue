
<script>
import Avatar from "@/components/Avatar.vue"
import CustomText from "@/components/CustomText.vue"
export default {
    name: 'list',
    components: {
        Avatar,
        CustomText,
    }
}
</script>

<template>
    <div class="container">
        <div class="head section">
            <div class="list-title">Comments</div>
            <div class="header-more">
                <button type="button">
                    <font-awesome-icon icon="fa-solid fa-xmark" size="2x" />
                </button>
            </div>
        </div>
        <div class="comment section">
            <ul>
                <li>
                    <div class="short-profile">
                        <Avatar :size="40" class="profile-photo" />
                        <div class="author-username">
                            <CustomText tag="b">_alevecchi</CustomText>
                        </div>
                        <span class="comment-span">It's a wonderful day</span>
                <CustomText size="xxsmall" class="time-ago">10 MINUTES AGO</CustomText>
                    </div>
                </li>
                <li>
                    <div class="short-profile">
                        <Avatar :size="40" class="profile-photo" />
                        <div class="author-username">
                            <CustomText tag="b">m.zuckerberg</CustomText>
                        </div>
                        <span class="comment-span">Wait for the metaverse...</span>
                <CustomText size="xxsmall" class="time-ago">30 MINUTES AGO</CustomText>
                    </div>
                    <div class="short-profile">
                        <Avatar :size="40" class="profile-photo" />
                        <div class="author-username">
                            <CustomText tag="b">_alevecchi</CustomText>
                        </div>
                    </div>
                </li>
                <li>
                    <div class="short-profile">
                        <Avatar :size="40" class="profile-photo" />
                        <div class="author-username">
                            <CustomText tag="b">_alevecchi</CustomText>
                        </div>
                    </div>
                </li>
            </ul>
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

.container .comment {
    display: flex;
    align-items: flex-start;
    justify-content: flex-start;
    width: 300px;
    height: 300px;
    border-radius: 0 0 20px 20px;
    background: var(--background-likes);
}

.container .comment .short-profile {
    display: flex;
    flex-direction: rows;
    margin-top: 4px;
}

.container .comment .short-profile .profile-photo:hover {
    cursor: pointer;
}

.container .comment .short-profile .author-username {
    align-items: center;
    margin-left: 8px;
    font-size: 16px;
    border-bottom: 1px solid #999;
}

.container .comment .short-profile .author-username b:hover {
    text-decoration: underline;
    cursor: pointer;
}
.container .comment .short-profile .author-username .comment-span {
            margin-left: 4px;
            overflow: auto;
          }
</style>