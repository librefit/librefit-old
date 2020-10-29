<template>
  <div>
    <base-material-card color="green" class="px-5 py-3">
      <template v-slot:heading>
        <div class="display-1 font-weight-light">
          Food Diary
        </div>
      </template>
      <v-card-text>
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
            <v-btn text color="primary" @click="dateMenu = false">Cancel</v-btn>
            <v-btn text color="primary" @click="$refs.menu.save(date)"
              >OK</v-btn
            >
          </v-date-picker>
        </v-menu>
      </v-card-text>

      <div class="py-3" />

      <v-row justify="center">
        <v-expansion-panels>
          <v-expansion-panel>
            <v-expansion-panel-header>Breakfast</v-expansion-panel-header>
            <v-expansion-panel-content>
              <v-data-table
                :headers="headers"
                :items="desserts"
                class="elevation-1"
                disable-pagination
                :hide-default-footer="true"
              >
                <template v-slot:item.calories="{ item }">
                  <v-chip :color="getColor(item.calories)" dark>{{
                    item.calories
                  }}</v-chip>
                </template>
              </v-data-table>
            </v-expansion-panel-content>
          </v-expansion-panel>
          <v-expansion-panel>
            <v-expansion-panel-header>Lunch</v-expansion-panel-header>
            <v-expansion-panel-content>
              Lunch
            </v-expansion-panel-content>
          </v-expansion-panel>
          <v-expansion-panel>
            <v-expansion-panel-header>Dinner</v-expansion-panel-header>
            <v-expansion-panel-content>
              Dinner
            </v-expansion-panel-content>
          </v-expansion-panel>
        </v-expansion-panels>
      </v-row>
      <div class="py-3" />

      <v-fab-transition>
        <v-btn to='/food/inventory/add' color="pink" dark absolute bottom right fab>
          <v-icon>mdi-plus</v-icon>
        </v-btn>
      </v-fab-transition>
    </base-material-card>

    <div class="py-6" />
  </div>
</template>

<script>
import baseMaterialCard from '@/components/MaterialCard'

export default {
  components: { baseMaterialCard },
  data: () => ({
    dateMenu: false,
    date: new Date().toISOString().substr(0, 10),
    headers: [
      {
        text: 'Meal (100g serving)',
        align: 'start',
        sortable: false,
        value: 'name'
      },
      { text: 'Calories', value: 'calories' },
      { text: 'Fat (g)', value: 'fat' },
      { text: 'Carbs (g)', value: 'carbs' },
      { text: 'Protein (g)', value: 'protein' },
      { text: 'Iron (%)', value: 'iron' }
    ],
    desserts: [
      {
        name: 'Frozen Yogurt',
        calories: 159,
        fat: 6.0,
        carbs: 24,
        protein: 4.0,
        iron: '1%'
      },
      {
        name: 'Ice cream sandwich',
        calories: 237,
        fat: 9.0,
        carbs: 37,
        protein: 4.3,
        iron: '1%'
      },
      {
        name: 'KitKat',
        calories: 518,
        fat: 26.0,
        carbs: 65,
        protein: 7,
        iron: '6%'
      }
    ]
  }),

  computed: {
    dateText() {
      return this.date
    }
  },

  methods: {
    moveDate(direction) {
      var d = new Date(this.date)
      d.setDate(d.getDate() - direction)
      this.date = d.toISOString().substr(0, 10)
    },
    getColor(calories) {
      if (calories > 400) return 'red'
      else if (calories > 200) return 'orange'
      else return 'green'
    }
  }
}
</script>
