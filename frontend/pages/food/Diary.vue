<template>
  <div>
    <p class="text-center text-h4">
      Food Diary
    </p>

    <v-card class="mt-4 mb-4">
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
            <v-btn text color="primary" @click="dateMenu = false">Cancel</v-btn>
            <v-btn text color="primary" @click="$refs.menu.save(date)"
              >OK</v-btn
            >
          </v-date-picker>
        </v-menu>
      </v-card-text>
    </v-card>

    <v-card>
      <v-card-text>
        <p>Breakfast</p>

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

        <div class="py-3" />

        <p>Lunch</p>
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

        <div class="py-3" />

        <p>Dinner</p>
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

        <div class="py-3" />

        <p>Snacks</p>
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

        <div class="py-3" />

        <v-fab-transition>
          <v-btn
            to="/food/inventory/search"
            color="pink"
            dark
            absolute
            bottom
            right
            fab
          >
            <v-icon>mdi-plus</v-icon>
          </v-btn>
        </v-fab-transition>
      </v-card-text>
    </v-card>
  </div>
</template>

<script>
export default {
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
      { text: 'Calories (kcal)', value: 'calories' },
      { text: 'Carbs (g)', value: 'carbs' },
      { text: 'Fat (g)', value: 'fat' },
      { text: 'Protein (g)', value: 'protein' },
      { text: 'Sodium (g)', value: 'sodium' },
      { text: 'Sugar', value: 'sugar' }
    ],
    desserts: [
      {
        name: 'Frozen Yogurt',
        calories: 159,
        fat: 6.0,
        carbs: 24,
        protein: 4.0,
        sodium: '12',
        sugar: '2'
      },
      {
        name: 'Ice cream sandwich',
        calories: 237,
        fat: 9.0,
        carbs: 37,
        protein: 4.3,
        sodium: '12',
        sugar: '2'
      },
      {
        name: 'KitKat',
        calories: 518,
        fat: 26.0,
        carbs: 65,
        protein: 7,
        sodium: '12',
        sugar: '2'
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
