<template>
  <div>
    <p class="text-center text-h4">
      Your Personal Dashboard
    </p>

    <div>
      <v-card>
        <v-card-text class="text-center">
          <v-menu
            ref="menu"
            v-model="dateMenu"
            :close-on-content-click="false"
            :return-value.sync="date"
            transition="scale-transition"
            offset-y
            min-width="290px"
          >
            <template v-slot:activator="{ on, attrs }">
              <v-btn class="ma-1" color="green" dark @click="moveDate('+1')">
                <v-icon dark center>mdi-arrow-left</v-icon>
              </v-btn>

              <v-btn class="ma-1" color="green" dark @click="moveDate('-1')">
                <v-icon dark center>mdi-arrow-right</v-icon>
              </v-btn>

              <v-btn
                color="green"
                dark
                v-bind="attrs"
                v-on="on"
                @click="dateMenu = true"
                >{{ date | formatDate }}
              </v-btn>
            </template>
            <v-date-picker v-model="date">
              <v-spacer></v-spacer>
              <v-btn text color="primary" @click="dateMenu = false"
                >Cancel</v-btn
              >
              <v-btn text color="primary" @click="$refs.menu.save(date)"
                >OK</v-btn
              >
            </v-date-picker>
          </v-menu>
        </v-card-text>
      </v-card>
    </div>

    <div class="py-3" />

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
            icon="mdi-shoe-print"
            title="Total Steps"
            value="12000"
            class="ma-3"
          />
        </v-col>
      </v-row>

      <div class="py-3" />

      <v-row>
        <v-col cols="6" md="4">
          <base-material-stats-card
            color="primary"
            icon="mdi-calculator"
            title="BMI"
            value="12000"
            class="ma-3"
          />
        </v-col>
        <v-col cols="6" md="4">
          <base-material-stats-card
            color="primary"
            icon="mdi-silverware-fork-knife"
            title="Total Calories"
            :value="totalCalories || 'N/A'"
            class="ma-3"
          />
        </v-col>
        <v-col cols="6" md="4">
          <v-card class="mx-auto" max-width="450">
            <v-card-title>
              <v-icon large left>
                mdi-chart-pie
                <v-card class="mx-auto" max-width="450">
                  <v-card-title>
                    <v-icon large left>
                      mdi-chart-pie
                    </v-icon>
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
    dateMenu: false,
    date: new Date().toISOString().substr(0, 10),
    getPieChartOptions: {
      responsive: true,
      maintainAspectRatio: false
    }
  }),

  computed: {
    totalLiquids() {
      return this.$store.getters['stats/getFluidsByDay'](this.date)
    },
    lastMeasure() {
      return this.$store.getters['stats/lastMeasureTaken'](this.date)
    },
    getPieChartData() {
      return this.$store.getters['stats/getFoodByDayPieChart'](this.date)
    },
    totalCalories() {
      var d = this.$store.getters['stats/getFoodByDay'](this.date)
      if (d != 'N/A') {
        return d.calories.toString()
      }
    }
  },

  mounted() {
    this.$store.dispatch('stats/pullAll')
  },

  methods: {
    moveDate(direction) {
      var d = new Date(this.date)
      d.setDate(d.getDate() - direction)
      this.date = d.toISOString().substr(0, 10)
    }
  }
}
</script>
