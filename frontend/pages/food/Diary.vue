<template>
  <div>
    <p class="text-center text-h4">Food Diary</p>

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
          <v-date-picker v-model="date" @input="$refs.menu.save(date)">
          </v-date-picker>
        </v-menu>
      </v-card-text>
    </v-card>

    <v-card>
      <v-card-text>
        <p>Breakfast</p>

        <v-data-table
          :headers="headers"
          :items="getFoodDiary('B')"
          v-if="getFoodDiary('B').length"
          class="elevation-1"
          disable-pagination
          :hide-default-footer="true"
        >
          <template v-slot:item.calories="{ item }">
            <v-chip :color="getColor(item.calories)" dark>{{
              item.calories
            }}</v-chip>
          </template>
          <template v-slot:item.actions="{ item }">
            <v-icon small class="mr-2"> mdi-pencil </v-icon>
            <v-icon small @click="deleteItem(item)"> mdi-delete </v-icon>
          </template>
        </v-data-table>
        <p v-else class="text-center font-italic">
          You don't have anything logged for Breakfast
        </p>

        <div class="py-3" />

        <p>Lunch</p>

        <v-data-table
          :headers="headers"
          v-if="getFoodDiary('L').length"
          :items="getFoodDiary('L')"
          class="elevation-1"
          disable-pagination
          :hide-default-footer="true"
        >
          <template v-slot:item.calories="{ item }">
            <v-chip :color="getColor(item.calories)" dark>{{
              item.calories
            }}</v-chip>
          </template>
          <template v-slot:item.actions="{ item }">
            <v-icon small class="mr-2"> mdi-pencil </v-icon>
            <v-icon small @click="deleteItem(item)"> mdi-delete </v-icon>
          </template>
        </v-data-table>
        <p v-else class="text-center font-italic">
          You don't have anything logged for Lunch
        </p>

        <div class="py-3" />

        <p>Dinner</p>
        <v-data-table
          :headers="headers"
          :items="getFoodDiary('D')"
          v-if="getFoodDiary('D').length"
          class="elevation-1"
          disable-pagination
          :hide-default-footer="true"
        >
          <template v-slot:item.calories="{ item }">
            <v-chip :color="getColor(item.calories)" dark>{{
              item.calories
            }}</v-chip>
          </template>
          <template v-slot:item.actions="{ item }">
            <v-icon small class="mr-2" @click="editItem(item)">
              mdi-pencil
            </v-icon>
            <v-icon small @click="deleteItem(item)"> mdi-delete </v-icon>
          </template>
        </v-data-table>
        <p v-else class="text-center font-italic">
          You don't have anything logged for Dinner
        </p>

        <div class="py-3" />

        <p>Snacks</p>
        <v-data-table
          :headers="headers"
          v-if="getFoodDiary('S').length"
          :items="getFoodDiary('S')"
          class="elevation-1"
          disable-pagination
          :hide-default-footer="true"
        >
          <template v-slot:item.calories="{ item }">
            <v-chip :color="getColor(item.calories)" dark>{{
              item.calories
            }}</v-chip>
          </template>

          <template v-slot:item.actions="{ item }">
            <v-icon small class="mr-2" @click="editItem(item)">
              mdi-pencil
            </v-icon>
            <v-icon small @click="deleteItem(item)"> mdi-delete </v-icon>
          </template>
        </v-data-table>
        <p v-else class="text-center font-italic">
          You don't have anything logged for Snacks
        </p>

        <div class="py-3" />

        <v-fab-transition>
          <v-btn
            to="/food/inventory"
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
import { mapActions, mapGetters } from 'vuex'

export default {
  data: () => ({
    dateMenu: false,
    date: new Date().toISOString().substr(0, 10),
    foodDiaryArr: [],
    headers: [
      {
        text: 'Meal (100g serving)',
        align: 'start',
        sortable: false,
        value: 'product_name',
      },
      { text: 'Calories (kcal)', value: 'calories' },
      { text: 'Carbs (g)', value: 'carbs' },
      { text: 'Fat (g)', value: 'fat_total' },
      { text: 'Protein (g)', value: 'proteins' },
      { text: 'Sodium (g)', value: 'sodium' },
      { text: 'Sugars', value: 'sugars' },
      { text: 'Actions', value: 'actions', sortable: false },
    ],
  }),

  mounted() {
    this.$store.dispatch('food/foodDiary', { start: this.date, end: this.date })
  },

  computed: {
    ...mapGetters('food', ['getFoodDiary']),

    dateText() {
      return this.date
    },
  },

  watch: {
    date: function (val) {
      const d = new Date(val).toISOString().slice(0, 10)
      this.foodDiary({ start: d, end: d })
    },
  },

  methods: {
    ...mapActions('food', ['foodDiary', 'deleteDiary']),

    deleteItem(item) {
      confirm('Are you sure you want to delete this entry?') &&
        this.deleteDiary(item.id)
    },

    moveDate(direction) {
      var d = new Date(this.date)
      d.setDate(d.getDate() - direction)
      this.date = d.toISOString().substr(0, 10)
      this.foodDiary({ start: this.date, end: this.date })
    },
    getColor(calories) {
      if (calories > 400) return 'red'
      else if (calories > 200) return 'orange'
      else return 'green'
    },
  },
}
</script>
