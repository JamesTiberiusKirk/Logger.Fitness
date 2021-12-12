<template>
  <div class="fab-container">
    <div class="fab fab-icon-holder">
      <i @click="fabClick" class="fas fa-question material-icons"> add </i>
    </div>
  </div>
</template>

<script>
export default {
  name: "Fab",
  props: {},
  emits: ["fabClickEvent"],
  methods: {
    fabClick() {
      this.$emit("fabClickEvent");
    },
  },
};
</script>

<style scoped>
.fab-container {
  position: fixed;
  bottom: 80px;
  right: 20px;
  z-index: 900;
  cursor: pointer;
}

.fab-icon-holder {
  width: 45px;
  height: 45px;
  border-radius: 100%;
  background: black;
  box-shadow: 0 6px 20px rgba(0, 0, 0, 0.9);
}

.fab-icon-holder:hover {
  opacity: 0.8;
}

.fab-icon-holder i {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  font-size: 25px;
  color: #ffffff;
}

.fab {
  width: 60px;
  height: 60px;
  background: var(--primary);
}

.fab-options {
  list-style-type: none;
  margin: 0;
  position: absolute;
  bottom: 70px;
  right: 0px;
  opacity: 0;
  transition: all 0.2s ease;
  transform: scale(0);
  transform-origin: 70% bottom;
}

.fab:hover + .fab-options,
.fab-options:hover {
  opacity: 1;
  transform: scale(1);
}

.fab-options li {
  display: flex;
  justify-content: flex-end;
  padding: 5px;
}
</style>