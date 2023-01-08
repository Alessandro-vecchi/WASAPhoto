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
            <font-awesome-icon icon="fa-solid fa-ellipsis" size="3x"/>
        </button>
      </div>
    </header>

    <!-- media -->
    <div class="post-media">
      <img src="https://picsum.photos/600/400?random=1" alt="" />
    </div>

    <!-- action -->
    <div class="action-buttons section">
      <button type="button">
        <font-awesome-icon class="icon-like" icon="fa-regular fa-heart" size="2x"/>
      </button>

      <button type="button">
        <font-awesome-icon icon="fa-regular fa-comment" size="2x"/>
      </button>
    </div>

    <!-- view count -->
    <div class="view-count section">
      <CustomText tag="b"> Likes: 897 </CustomText>
    </div>

    <!-- comments list -->
    <div class="comments-list ">
      <div class="view-all section">View all 56 comments</div>
      <div class="comments section">
        <ul>
          <li>
            <CustomText tag="b">j.bezos</CustomText>
            <span class="comment-span">W Amazon!</span>
          </li>
          <li>
            <CustomText tag="b">m.zuckerberg</CustomText>
            <span class="comment-span">Wait for the metaverse...</span>
          </li>
          
        </ul>
      </div>
      <div class="time section">
        <CustomText size="xxsmall"  class="time-ago">10 MINUTES AGO</CustomText>
      </div>
      <div class="comment section">
        <Avatar :size="20" />
        <input class="text-body" type="text" placeholder="Add a comment...">
        <a href="#" type="button">Post</a>
      </div>
    </div>

    <!-- datetime -->

    <!-- comments form -->
  </div>
</template>


<style>
.post {
  border-radius: 3px;
  border: 1px solid rgba(219,219,219, 1);
  max-width: 601px !important;
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
  font-size: 15px;
}
.post .header-more {
  margin-left: auto;
}
.post .action-buttons {
  height: 44px;
  display: flex;
  align-items: center;
}
.post .action-buttons button {
  margin-right: 16px;
}
.post .action-buttons .action-save {
  margin-left: auto;
  margin-right: 0;
}
.post .view-count {
  margin-bottom: 8px;
}
.post .comments-list .view-all {
  color: rgba(142,142,142, 1);
}
.post .comments-list .comments ul li {
  display: flex;
  align-items: flex-start;
}
.post .comments-list .comments ul li b:hover {
  text-decoration: underline;
  cursor: pointer;
}
.post .comments-list .comments ul li .comment-span {
  margin-left: 4px;
  overflow: auto;
}
.post .comments-list .comments ul li .icon-like {
  height: 12px;
  width: 12px;
  margin-left: auto;
}
.post .comments-list .time-ago {
  color: rgba(142,142,142,1);
  margin-bottom: 4px;
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
  color: rgba(0,149,246,1);
}
</style>