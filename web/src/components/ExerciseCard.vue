<template>
  <div class="exercise col-md-12">
    <div class="card card-container">
      <div class="card-body">
        <div class="d-flex w-100 justify-content-between">
          <h5 class="mb-1">
            {{ exerciseType.name }}
          </h5>
          <small>
            <i
              v-if="exerciseType.data_type === 'sets'"
              @click="showAddSetModal(e)"
              class="fas fa-question material-icons"
            >
              add
            </i>
            <i
              v-if="exerciseType.data_type === 'single-value'"
              @click="showAddSingleValueModal(e)"
              class="header-button fas fa-question material-icons"
            >
              edit
            </i>
            <i
              @click="deleteExerciseModalShow"
              class="header-button fas fa-question material-icons"
              >delete</i
            >
          </small>
        </div>
        <span v-if="exerciseType.data_type === 'sets'">
          <table class="table">
            <thead>
              <tr>
                <th scope="col">#</th>
                <th scope="col">Resistance</th>
                <th scope="col">Reps</th>
                <th scope="col"></th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(s, i) in exercise.sets" :key="i">
                <th scope="row">{{ calculateSetNumber(i) }}</th>
                <td>{{ s.resistance }}{{ exerciseType.measurement_type }}</td>
                <td>{{ s.reps }}</td>

                <i
                  @click="deleteSetModal(i)"
                  class="set-remove-button fas fa-question material-icons"
                  >remove_circle_outline</i
                >
              </tr>
            </tbody>
          </table>
        </span>
        <span v-if="exerciseType.data_type === 'single-value'">
          <span v-if="exercise.single_value">
            Single value:
            {{ exercise.single_value.value }}
            {{ exerciseType.measurement_type }}
          </span>
          <span v-if="!exercise.single_value">
            No data yet
          </span>
        </span>
      </div>
    </div>
  </div>

  <SetsModal
    v-if="setModal.show"
    :exercise="exercise"
    :exerciseType="exerciseType"
    @closeModal="closeAddSetModal"
  />

  <SingleValueModal
    v-if="singleValueModal.show"
    :exercise="exercise"
    :exerciseType="exerciseType"
    @closeModal="closeAddSingleValueModal"
  />

  <DeleteModal
    class="custom-modal"
    v-if="deleteModal.show"
    :deleteMessage="deleteModal.message"
    :continueButtonMessage="deleteModal.continueButtonMessage"
    @closeDeleteModalEvent="deleteExerciseModalClose"
    @deleteRecordEvent="deleteExercise"
  />

  <Toast
    v-if="toast.show"
    :title="toast.title"
    :time="toast.time"
    @closeToast="toastHide"
  />
</template>

<script>
import DeleteModal from "./DeleteModal.vue";
import SetsModal from "./SetsModal.vue";
import SingleValueModal from "./SingleValueModal.vue";
import Toast from "./Toast.vue";

export default {
  name: "ExerciseCard",
  components: {
    DeleteModal,
    SingleValueModal,
    SetsModal,
    Toast,
  },
  props: {
    exercise: {
      type: Object,
      required: true,
    },
    exerciseType: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      setModal: {
        show: false,
      },
      singleValueModal: {
        show: false,
      },
      deleteModal: {
        show: false,
        message: "Are you sure you want to delete this exercise?",
        continueButtonMessage: "Yes",
      },
      toast: {
        show: false,
        title: "Not yet implemented",
        time: 2500,
      },
    };
  },
  mounted() {
    console.log(this.exercise);
  },
  methods: {
    deleteExercise() {
      this.$store
        .dispatch("exercises/deleteOne", this.exercise.exercise_id)
        .then(
          (res) => {
            console.log(res);

            // TODO: BUG: encountered this before, need to find a way to update the displayed data when state changes so forcing a page refresh instead.
            this.$router.go();
          },
          (err) => {
            console.error(err);
          }
        );
    },
    deleteExerciseModalClose() {
      this.deleteModal.show = false;
    },
    deleteExerciseModalShow() {
      this.deleteModal.show = true;
    },
    showAddSetModal() {
      this.setModal.show = true;
    },
    closeAddSetModal() {
      this.setModal.show = false;
    },
    showAddSingleValueModal() {
      this.singleValueModal.show = true;
    },
    closeAddSingleValueModal() {
      this.singleValueModal.show = false;
    },
    deleteSetModal() {
      this.toast.show = true;
    },
    toastHide() {
      this.toast.show = false;
    },
    calculateSetNumber(arrI) {
      if (this.exercise.sets[arrI].is_drop_set) return "";

      let offset = 1;
      this.exercise.sets.forEach((set, index) => {
        if (arrI <= index) return;
        if (set.is_drop_set) {
          console.log(set);
          offset = offset - 1;
          console.log(offset);
        }
      });
      return arrI + offset;
    },
  },
};
</script>

<style scoped>
.exercise {
  margin-bottom: 5px;
  margin-right: 0px;
  padding: 0px;
  width: 100%;
}
.set-remove-button {
  margin-top: 10px;
  /* padding-top: 5px; */
}
.header-button{
  margin-left:15px;
}
</style>