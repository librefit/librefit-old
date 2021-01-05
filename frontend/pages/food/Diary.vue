<template>
  <div>
    <p class="text-center text-h4">Food Diary</p>

    <DateMenu :date="date" @update-date="moveDate" @move-date="moveDate" />

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
            <v-btn class="mr-2" text icon :to="'/food/diary/' + item.id">
              <v-icon small>mdi-pencil</v-icon>
            </v-btn>
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
            <v-btn class="mr-2" text icon :to="'/food/diary/' + item.id">
              <v-icon small>mdi-pencil</v-icon>
            </v-btn>
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
            <v-btn class="mr-2" text icon :to="'/food/diary/' + item.id">
              <v-icon small>mdi-pencil</v-icon>
            </v-btn>
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
            <v-btn class="mr-2" text icon :to="'/food/diary/' + item.id">
              <v-icon small>mdi-pencil</v-icon>
            </v-btn>
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
    date: new Date().toISOString().substr(0, 10),
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
    this.$store.dispatch('food/foodDiaryRange', {
      start: this.date,
      end: this.date,
    })
  },

  computed: {
    ...mapGetters('food', ['getFoodDiary']),
  },

  methods: {
    ...mapActions('food', ['foodDiaryRange', 'deleteDiary']),

    deleteItem(item) {
      confirm('Are you sure you want to delete this entry?') &&
        this.deleteDiary(item.id)
    },

    moveDate(newDate) {
      this.date = newDate
      this.foodDiaryRange({ start: this.date, end: this.date })
    },

    getColor(calories) {
      if (calories > 400) return 'red'
      else if (calories > 200) return 'orange'
      else return 'green'
    },
  },
}
</script>
