<template>
  <div>
    <p class="text-center text-h4">Your Personal Dashboard</p>

    <DateMenu :date="date" @update-date="moveDate" @move-date="moveDate" />

    <v-container class="grey lighten-5">
      <v-row>
        <v-col cols="6" md="4">
          <base-material-stats-card
            color="primary"
            icon="mdi-water"
            title="Total Liquids"
            :value="totalLiquids || 'N/A'"
            class="ma-3"
          />
        </v-col>
        <v-col cols="6" md="4">
          <base-material-stats-card
            color="primary"
            icon="mdi-scale-bathroom"
            title="Last registered weight"
            :value="lastMeasure || 'N/A'"
            class="ma-3"
          />
        </v-col>
        <v-col cols="6" md="4">
          <base-material-stats-card
            color="primary"
            icon="mdi-calculator"
            title="BMI"
            :value="BMI"
            class="ma-3"
          />
        </v-col>
      </v-row>

      <div class="py-3" />

      <v-row>
        <v-col cols="6" md="4">
          <base-material-stats-card
            color="primary"
            icon="mdi-silverware-fork-knife"
            title="Total Calories"
            :value="totalCalories || 'N/A'"
            class="ma-3"
          />
        </v-col>
        <v-col cols="6" md="4"> </v-col>
        <v-col cols="6" md="4">
          <v-card class="mx-auto" max-width="450">
            <v-card-title>
              <v-icon large left>
                mdi-chart-pie
                <v-card class="mx-auto" max-width="450">
                  <v-card-title>
                    <v-icon large left> mdi-chart-pie </v-icon>
                    <span class="title font-weight-light"
                      >Nutrients Distribution</span
                    >
                  </v-card-title>

                  <v-card-text>
                    <PieChart
                      :chartData="getPieChartData"
                      :options="getPieChartOptions"
                    />
                  </v-card-text>
                </v-card>
              </v-icon>
              <span class="title font-weight-light"
                >Nutrients Distribution</span
              >
            </v-card-title>

            <v-card-text>
              <PieChart
                v-if="getPieChartData"
                :chartData="getPieChartData"
                :options="getPieChartOptions"
              />
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
export default {
  middleware: 'auth',
  data: () => ({
    date: new Date().toISOString().substr(0, 10),
    getPieChartOptions: {
      responsive: true,
      maintainAspectRatio: false,
    },
  }),

  computed: {
    BMI() {
      const height = this.$store.getters['user/getHeight'] / 100 // Taking it to meters
      const weight = this.$store.getters['stats/lastMeasureTaken'](this.date)
      const bmi = weight / (height * height)
      return bmi.toFixed(1).toString() + ' kg/m²'
    },
    totalLiquids() {
      const l = this.$store.getters['stats/getFluidsByDay'](this.date)
      if (l != 'N/A') {
        return l.toString() + ' ml'
      }
    },
    lastMeasure() {
      return (
        this.$store.getters['stats/lastMeasureTaken'](this.date).toString() +
        ' kg'
      )
    },
    getPieChartData() {
      return this.$store.getters['stats/getFoodByDayPieChart'](this.date)
    },
    totalCalories() {
      var d = this.$store.getters['stats/getFoodByDay'](this.date)
      if (d != 'N/A') {
        return d.calories.toString() + ' kcal'
      }
    },
  },

  mounted() {
    this.$store.dispatch('stats/pullAll')
    this.$store.dispatch('user/pullSettings')
  },

  methods: {
    moveDate(newDate) {
      this.date = newDate
    },
  },
}
</script>
