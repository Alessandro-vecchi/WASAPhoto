<script>
import Avatar from "@/components/Avatar.vue"
import CustomText from "@/components/CustomText.vue"
export default {
    name: 'Post',
    components: {
        Avatar,
        CustomText,
    }
}
</script>
    CustomText

<template>
    <div class="post">
        <!-- header -->
        <header class="header section">
            <div class="header-author">
                <Avatar :size="40" />
                <div class="header-author-info">
                    <CustomText tag="b">_alevecchi</CustomText>
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
            <img src="https://picsum.photos/600/400?random=1" alt="" />
        </div>

        <div class="two-col section">
            <!-- action & count-->
            <div class="action-buttons">
                <ul>
                    <li>
                        <button type="button">
                            <font-awesome-icon class="icon-like" icon="fa-regular fa-heart" size="2x" />
                            <span class="num"> 897 </span>
                        </button>
                    </li>
                    <li>
                        <button type="button">
                            <font-awesome-icon class="icon-comment" icon="fa-regular fa-comment" size="2x" />
                            <span class="num"> 33 </span>
                        </button>
                    </li>
                </ul>
            </div>

            <div class="caption">
                <li>
                    <CustomText tag="b">_alevecchi</CustomText>
                    <span class="caption-span">It's a wonderful day</span>
                </li>
            </div>
        </div>

        <div class="comments-list">
            <!-- datetime-->
            <div class="time section">
                <CustomText size="xxsmall" class="time-ago">10 MINUTES AGO</CustomText>
            </div>

            <!-- comments form -->
            <div class="comment section">
                <Avatar :size="20" />
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

.post .two-col  {
    display: flex;
    align-items: center;
   /*outline: solid;*/
}

.post .action-buttons {
    margin-top: 0px;
    height: 60px;
    display: flex;
    align-items: center;
}

.post .action-buttons button {
    padding-top: 5px;
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
    font-variant-numeric: oldstyle-nums;
    color: rgb(04,40,40);
}

.post .action-buttons .icon-like {
    height: 25px;
    width: 25px;
}

.post .action-buttons .icon-comment {
    height: 25px;
    width: 25px;
}
.post .caption{
    display: flex;
    flex-wrap: wrap;
    align-items: flex-start;
   /*outline: solid;*/
    margin-bottom: 0%;
    margin-left: 10%;
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