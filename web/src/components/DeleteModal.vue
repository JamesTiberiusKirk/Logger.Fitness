<template>
  <div id="modal-dialog" class="modal-dialog-container">
    <div class="modal-dialog-content">
      <div class="modal-dialog-header">
        <h4>{{ deleteMessage }}</h4>
      </div>
      <div class="modal-dialog-body"></div>
      <div class="modal-dialog-footer">
        <div class="btn-wrapper">
          <button
            type="button"
            class="btn btn-primary btn-sm"
            @click="deleteRecord"
          >
            {{ continueButtonMessage }}
          </button>
          <button
            type="button"
            class="btn btn-secondary btn-sm close-btn"
            @click="closeModal"
          >
            Cancel
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "DeleteModal",
  props: {
    deleteMessage: {
      type: String,
      required: true,
    },
    continueButtonMessage: {
      type: String,
      required: true,
    },
  },
  emits: ["closeDeleteModalEvent", "deleteRecordEvent"],
  methods: {
    closeModal() {
      this.$emit("closeDeleteModalEvent");
    },
    deleteRecord() {
      this.$emit("deleteRecordEvent");
    },
  },
};
</script>

<style scoped>
.modal-dialog-container {
  position: fixed; /* Stay in place */
  z-index: 901; /* Sit on top */
  left: 0;
  top: 0;
  width: 100%; /* Full width */
  height: 100%; /* Full height */
  overflow: auto; /* Enable scroll if needed */
  background-color: rgb(0, 0, 0); /* Fallback color */
  background-color: rgba(0, 0, 0, 0.4); /* Black w/ opacity */
}
.modal-dialog-content {
  background-color: #fefefe;
  margin: 10% auto;
  margin-top: 100%;
  padding: 20px;
  border: 1px solid #888;
  border-radius: 0.3rem;
  width: 80%;
}
.modal-dialog-footer {
  text-align: center;
}
.close-btn {
  margin-left: 1em;
}
</style>