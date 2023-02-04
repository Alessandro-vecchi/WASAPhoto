<!-- <template>
  <div>
    <button @click="showModal = true">Delete Account</button>
    <Modal v-if="showModal" @close="showModal = false">
      <h1>Are you sure?</h1>
      <p>This will permanently delete your profile</p>
      <div>
        <button @click="confirmDelete">Yes</button>
        <button @click="showModal = false">No</button>
      </div>
    </Modal>
  </div>
  
</template> -->

<template>
    <div class="modal-container">
        <div class="modal-overlay"></div>
        <div class="modal">
            <div class="modal-header">
                <h2>Confirm Deletion</h2>
            </div>
            <div class="modal-body">
                <p>Are you sure? This will permanently delete your profile.</p>
            </div>
            <div class="modal-footer">
                <button class="cancel-button" @click="cancel">Cancel</button>
                <button class="confirm-button" @click="confirmDelete">Confirm</button>
            </div>
        </div>
    </div>
</template>

<script>

export default {
    mounted() {

    },
    methods: {

    }
};
</script>

<style scoped>
.modal-container {
    position: fixed;
    top: 0;
    left: 0;
    height: 100%;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    /*  visibility: hidden; */
}

.modal-overlay {
    position: absolute;
    top: 0;
    left: 0;
    height: 100%;
    width: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    /* visibility: hidden; */
}

.modal {
    background-color: black;
    border-radius: 5px;
    box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.5);
    width: 400px;
    max-width: 90%;
    padding: 20px;
    text-align: center;
    /* visibility: hidden; */
}

.modal-header {
    border-bottom: 1px solid lightgray;
    margin-bottom: 20px;
}

.modal-footer {
    border-top: 1px solid lightgray;
    display: flex;
    justify-content: flex-end;
    margin-top: 20px;
    padding-top: 20px;
}

.cancel-button,
.confirm-button {
    background-color: transparent;
    border: 1px solid gray;
    border-radius: 5px;
    padding: 10px 20px;
    font-size: 14px;
    cursor: pointer;
}

.cancel-button:hover,
.confirm-button:hover {
    background-color: lightgray;
}
</style>
