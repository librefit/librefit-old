<template>
  <div>
    <p class="text-center text-h5" v-if="getDiaryItem.food_inventory">{{ getDiaryItem.food_inventory.product_name }}</p>

    <v-overlay v-if="getLoadingStatus" class="text-center">
      <v-progress-circular :size="70" indeterminate></v-progress-circular>
    </v-overlay>

    <v-container class="grey lighten-5">
      <v-row>
        <v-col>
          <v-card class="pa-2" outlined tile>
            <v-form ref="form" lazy-validation v-model="valid">
              <v-container>
                <v-row>
                  <v-col cols="12" md="4">
                    <v-select
                      v-model="mealSelect"
                      :items="meal"
                      item-text="text"
                      item-value="value"
                      :rules="[(v) => !!v || 'Item is required']"
                      label="To which meal?"
                      required
                    ></v-select>
                  </v-col>
                  <v-col cols="12" md="4">
                    <v-menu
                      v-model="menuCalendar"
                      :nudge-right="40"
                      transition="scale-transition"
                      offset-y
                      min-width="290px"
                    >
                      <template v-slot:activator="{ on, attrs }">
                        <v-text-field
                          v-model="dateDiary"
                          label="Date"
                          prepend-icon="mdi-calendar"
                          v-bind="attrs"
                          v-on="on"
                        ></v-text-field>
                      </template>
                      <v-date-picker
                        v-model="dateDiary"
                        @input="menuCalendar = false"
                      ></v-date-picker>
                    </v-menu>
                  </v-col>
                </v-row>
                <v-row>
                  <v-col cols="12" md="4">
                    <v-text-field
                      v-model.number="quantity"
                      label="Serving"
                      required
                    ></v-text-field>
                  </v-col>
                  <v-col cols="12" md="4">
                    <v-select
                      v-model="unitSelect"
                      :items="units"
                      item-text="text"
                      item-value="value"
                      :rules="[(v) => !!v || 'Item is required']"
                      label="Units"
                      required
                    ></v-select>
                  </v-col>
                  <v-col cols="12" md="6">
                    <!-- <p v-if="total_energy" class="text-subtitle-1 mt-3"> -->
                    <!--   Total calories: <b>{{ total_energy }}</b> -->
                    <!-- </p> -->
                  </v-col>
                </v-row>
              </v-container>
              <v-btn depressed color="primary" @click="save">
                Log to diary
              </v-btn>
            </v-form>
          </v-card>
        </v-col>
        <v-col align="center">
          <NutritionFactsInventory v-if="getDiaryItem.food_inventory" :product="getDiaryItem.food_inventory" />
          <v-carousel
            hide-delimiter-background
            delimiter-icon="mdi-minus"
            v-if="getDiaryItem.food_inventory && getDiaryItem.food_inventory.images"
            height="400"
            class="mt-4"
            cycle
            interval="4000"
          >
            <v-carousel-item
              v-for="(item, i) in getDiaryItem.food_inventory.images"
              v-if="item"
              :key="i"
            >
              <img
                :src="baseUrlImg + item"
                style="width: 200px; height: auto"
                alt="image-from-openfoodfacts"
              />
            </v-carousel-item>
          </v-carousel>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import { mapActions, mapGetters, mapState } from 'vuex'

export default {
  data: () => ({
    baseUrlImg: process.env.baseUrlImg,
    menuCalendar: false,
    dateDiary: new Date().toISOString().substr(0, 10),
    valid: false,
    meal: [
      { text: 'Breakfast', value: 'B' },
      { text: 'Lunch', value: 'L' },
      { text: 'Dinner', value: 'D' },
      { text: 'Snacks', value: 'S' },
    ],
    units: [
      { text: 'x 100 g', value: 1 },
      { text: 'x 1 g', value: 2 },
      { text: 'x 3.5 ounce', value: 3 },
      { text: 'x 1 ounce', value: 4 },
    ],
    total_energy: null,
    favourite: false,
  }),

  async asyncData({ params }) {
    const slug = params.slug
    return { slug }
  },

  watch: {
    //quantity() {
    //  // TODO: Transform this.amount function of unitList
    //  this.total_energy = this.quantity * this.product.calories
    //},
  },

  mounted() {
    this.$store.dispatch('food/foodDiary', this.slug)
  },

  computed: {
    ...mapGetters('food', ['getLoadingStatus', 'getDiaryItem']),

    getProductImages() {
      return this.$store.getters['food/getInventoryItemImages'](this.slug)
    },
    
    mealSelect: {
      get() {
        return this.getDiaryItem?.meal_type
      },
      set(val) {
        this.$store.commit('food/updateDiaryForm', {
          value: val,
          element: 'meal_type',
        })
      },
    },
    
    unitSelect: {
      get() {
        return this.getDiaryItem?.quantity_unit
      },
      set(val) {
        this.$store.commit('food/updateDiaryForm', {
          value: val,
          element: 'quantity_unit',
        })
      },
    },
    
    quantity: {
      get() {
        return this.getDiaryItem?.quantity
      },
      set(val) {
        this.$store.commit('food/updateDiaryForm', {
          value: val,
          element: 'quantity',
        })
      },
    },
  },

  methods: {
    ...mapActions('food', ['updateDiary']),

    save() {
      if (!this.$refs.form.validate()) {
        return
      } else {
        this.updateDiary()
      }
    },
  },
}
</script>
