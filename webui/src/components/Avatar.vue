<template>
    <figure class="avatar">
        <img :src="resolvedSrc" :alt="alt" :width="size" :height="size" />
        <!-- <img v-if=src :src="src" :alt="alt" :width="size" :height="size" />
        <img v-else :src="src" :alt="alt" :width="size" :height="size" /> -->
    </figure>
</template>

<script>
export default {
    name: 'Avatar',
    props: {
        src: {
            type: String,
            default: `https://cdn.pixabay.com/photo/2018/11/13/22/01/avatar-3814081__480.png`,
        },
        alt: {
            type: String,
            default: 'image'
        },
        size: {
            type: Number,
            default: 56
        }
    },
    computed: {
        resolvedSrc() {
            console.log(this.src)
            if (!this.src) {
                return `https://cdn.pixabay.com/photo/2018/11/13/22/01/avatar-3814081__480.png`
            } else {
                return this.src
            }
        }
    }
}
</script>

<style>
.avatar {
    padding: 2px;
    border-radius: 50%;
    display: inline-flex;
    /*background: linear-gradient(219deg, rgba(184, 63, 140, 1) 0%, rgba(237, 158, 90, 1) 100%);*/
}

.avatar img {
    border: 2px solid white;
    border-radius: inherit;
    cursor: pointer;
}
</style>