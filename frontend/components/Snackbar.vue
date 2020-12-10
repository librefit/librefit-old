<template>
  <div class="text-center">
    <v-snackbar
      v-model="show"
      top
      :color="color"
      timeout="2000"
      :multi-line="multiLine"
      elevation="24"
    >
      {{ message }}

      <template v-slot:action="{ attrs }">
        <v-btn text v-bind="attrs" @click="snackbar = false">
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </div>
</template>

<script>
export default {
  data() {
    return {
      show: false,
      multiLine: true,
      message: "",
      color: ""
    };
  },

  created() {
    this.$store.subscribe((mutation, state) => {
      if (mutation.type === "snackbar/showMessage") {
        this.message = state.snackbar.content;
        this.color = state.snackbar.color;
        this.show = true;
      }
    });
  }
};
</script>
