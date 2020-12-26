<template>
  <div>
    <p class="text-center text-h5" v-if="product">{{ product.product_name }}</p>

    <v-container class="grey lighten-5">
      <v-row>
        <v-col>
          <v-card class="pa-2" outlined tile>
            <v-form v-model="valid">
              <v-container>
                <v-row>
                  <v-col cols="12" md="4">
                    <v-select
                      v-model="mealSelect"
                      :items="meal"
                      item-text="text"
                      item-value="value"
                      :rules="[v => !!v || 'Item is required']"
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
                      v-model="quantity"
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
                      :rules="[v => !!v || 'Item is required']"
                      label="Units"
                      required
                    ></v-select>
                  </v-col>
                  <v-col cols="12" md="6">
                    <p v-if="total_energy" class="text-subtitle-1 mt-3">
                      Total calories: <b>{{ total_energy }}</b>
                    </p>
                  </v-col>
                </v-row>
              </v-container>
              <v-btn depressed color="primary" @click="save">
                Log to diary
              </v-btn>
              <v-btn class="ml-3" @click="saveFavourite">
                <v-icon color="pink">
                  mdi-heart
                </v-icon>
                Favorite
              </v-btn>
            </v-form>
          </v-card>
        </v-col>
        <v-col align="center">
          <NutritionFactsInventory v-if="product" :product="product">
          </NutritionFactsInventory>
          <p v-if="product" class="text-subtitle-1 mt-4">
            No the right values?
            <a
              :href="
                'https://world.openfoodfacts.org/cgi/product.pl?type=edit&code=' +
                  product.off_code
              "
              target="_blank"
              >Edit product in Open Food Facts</a
            >
            (you must Sign-in)
          </p>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import { mapActions, mapGetters, mapState } from 'vuex'

export default {
  data: () => ({
    menuCalendar: false,
    dateDiary: new Date().toISOString().substr(0, 10),
    valid: false,
    quantity: 1.0,
    mealSelect: null,
    meal: [
      { text: 'Breakfast', value: 'B' },
      { text: 'Lunch', value: 'L' },
      { text: 'Dinner', value: 'D' },
      { text: 'Snacks', value: 'S' }
    ],
    unitSelect: 2,
    units: [
      { text: 'x 100 g', value: 1 },
      { text: 'x 1 g', value: 2 },
      { text: 'x 3.5 ounce', value: 3 },
      { text: 'x 1 ounce', value: 4 }
    ],
    total_energy: null,
    favourite: false
  }),

  async asyncData({ params }) {
    const slug = params.slug
    return { slug }
  },
  
  watch: {
    amount() {
      // TODO: Transform this.amount function of unitList
      this.total_energy = this.amount * this.product.nutriments.energy_value
    }
  },
  
  mounted() {
    this.$store.dispatch('food/pullInventory')
  },

  computed: {
    product() {
      return this.$store.getters['food/getInventoryItem'](this.slug)
    }
  },

  methods: {
    ...mapActions('food', ['add2Inventory', 'add2Diary']),

    saveFavourite() {
      const payload = {
        off_code: this.slug,
        favourite: this.favourite
      }

      this.add2Inventory(payload)
    },

    save() {
      const payload = {
        off_code: this.product.off_code,
        meal_type: this.mealSelect,
        quantity: this.quantity,
        quantity_unit: this.unitSelect,
        date: this.dateDiary
      }

      this.add2Diary(payload)
    }
  }
}
</script>
