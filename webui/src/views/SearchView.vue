<script>


</script>

<template>
    <div class="input-wrapper">
        <font-awesome-icon icon="fa-solid fa-magnifying-glass" inverse />
        <input class="search" type="text" placeholder="Search" />
        <font-awesome-icon icon="fa-solid fa-xmark" inverse />
    </div>
</template>

<style>

.input-wrapper {
    display: flex;
    align-items: center;
    justify-content: space-between;
    background: #e26d5c;
    padding: 0.5rem;
    border-radius: 0.5rem;
    color: white;
    box-shadow: 0.25rem 0.25rem 0rem #f0b2a8;
}

.search {
    margin: 0 0.5rem 0 0.5rem;
    width: 100%;
    border: none;
    outline: none;
    background: #e26d5c;
    color: white;
}

</style>