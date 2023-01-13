<script>
export default {
    name: 'GalleryItem',
    data: function () {
        return {
            loading: false,
            error: false,
            
        }

    },
}
</script>

<template>
    <div class="gallery-item" tabindex="0">
        <!-- "https://picsum.photos/600/400" -->
        <img :src=photos.photo alt="" class="gallery-image">
        <div class="gallery-item-info">
            <ul>
                <li class="gallery-item-likes"><span class="visually-hidden">Likes:</span><font-awesome-icon
                        icon="fa-solid fa-heart" size="xl" /> 13</li>
                <li><span class="visually-hidden">Comments:</span><font-awesome-icon icon="fa-solid fa-comment"
                        size="xl" /> 6</li>
            </ul>
        </div>
    </div>
</template>

<style>
.gallery-item {
    position: relative;
    flex-grow: 1;
    margin: 1rem;
    color: #fafafa;
    cursor: pointer;
}

.gallery-item:hover .gallery-item-info,
.gallery-item:focus .gallery-item-info {
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    top: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.3)
}

.gallery-item-info {
    display: none;
}

.gallery-item-info li {
    display: inline-block;
    font-size: 1.3rem;
    font-weight: 600;
}

.gallery-item-likes {
    margin-right: 1.5rem;
}

.gallery-image {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.visually-hidden {
    position: absolute;
    height: 1px;
    width: 1px;
    overflow: hidden;
    clip: rect(1px, 1px, 1px, 1px);
}

@supports (display:grid) {

    .gallery-item {
        width: auto;
        margin: 0;

    }
}
</style>