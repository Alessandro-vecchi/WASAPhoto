<template>
  <p>Hi</p>
  <header class="header">
    <Container class="header-inner">
      <!-- Logo -->
      <router-link to="/" class="logo">
        <div class="logo">
          <img alt="Instagram" src="../img/instagram-logo.png" />
        </div>
      </router-link>
      <!-- search -->
      <form class="search" action="#">
        <input type="text" placeholder="Search" />
      </form>
      <!-- navigation -->
      <nav class="navigation">
        <router-link to="/">
          <IconHomeFill v-if="$route.name === 'Home'" />
          <IconHome v-else />
        </router-link>
        <router-link to="/direct">
          <IconDirectFill v-if="$route.name === 'Direct'" />
          <IconDirect v-else />
        </router-link>
        <router-link to="/explore">
          <IconExploreFill v-if="$route.name === 'Explore'" />
          <IconExplore v-else />
        </router-link>
        <router-link to="/heart">
          <IconHeartFill v-if="$route.name === 'Heart'" />
          <IconHeart v-else />
        </router-link>
        <router-link to="/profile">
          <Avatar :size="20"/>
        </router-link>
        <icon />
      </nav>
    </Container>
  </header>
</template>

<script>
import Avatar from '@/components/Avatar'
export default {
  name: 'MainHeader',
  components: {
    Avatar,
  }
}
</script>

<style scoped>
.header {
z-index: 99;
position: fixed;
left: 0;
top: 0;
right: 0;
background-color: white;
border-bottom: 1px solid rgb(var(--b6a));
}

.header-inner {
display: flex;
align-items: center;
justify-content: center;
}

@media (--t) {
.header-inner {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  height: var(--header-height);
  align-items: center;
}
}

.logo {
display: flex;
align-items: center;
height: 36px;
}

.logo img {
height: 29px;
margin-top: 8px;
}

.search {
display: none;
}

@media (--t) {
.search {
  display: block;
  text-align: center;
}

.search input {
  border: 1px solid rgb(var(--b6a));
  border-radius: 4px;
  background-color: rgb(var(--b3f));
  width: 215px;
  height: 28px;
  padding-left: 20px;
  padding-right: 20px;
}

.search input::placeholder {
  display: flex;
  text-align: center;
}

.search input:focus {
  outline: none;
}
}

.navigation {
z-index: 99;
position: fixed;
left: 0;
right: 0;
bottom: 0;
background-color: white;
border-top: 1px solid rgb(var(--b6a));
display: flex;
align-items: center;
justify-content: space-around;
height: 44px;
}

.navigation a {
margin-left: auto;
margin-right: auto;
}

@media (--t) {
.navigation {
  border: none;
  position: static;
  justify-content: flex-end;
  align-items: flex-start;
  height: 22px;
}

.navigation a {
  margin-left: 22px;
  margin-right: 0;
}
}

.user-photo {
width: 22px;
height: 22px;
border-radius: 50%;
background-color: black;
}

</style>
