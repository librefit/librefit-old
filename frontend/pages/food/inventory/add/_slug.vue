<template>
  <div>
    <p class="text-center text-h5">Open Food Facts</p>

    <v-container class="grey lighten-5">
      <v-row>
        <v-col>
          <p class="text-center text-h5" v-if="product">
            {{ product.product_name }}
          </p>

          <v-card class="pa-2" outlined tile>
            <v-form v-model="valid">
              <v-container>
                <v-row>
                  <v-col cols="12" md="4">
                    <v-text-field
                      v-model="amount"
                      label="Serving"
                      required
                    ></v-text-field>
                  </v-col>
                  <v-col cols="12" md="4">
                    <v-select
                      v-model="unit"
                      :items="unitList"
                      :rules="[v => !!v || 'Item is required']"
                      label="Units"
                      required
                    ></v-select>
                  </v-col>
                </v-row>
                <v-row>
                  <v-col cols="12" md="4">
                    <v-select
                      v-model="meal"
                      :items="mealList"
                      :rules="[v => !!v || 'Item is required']"
                      label="To which meal?"
                      required
                    ></v-select>
                  </v-col>
                  <v-col cols="12" md="6">
                    <p v-if="total_energy" class="text-subtitle-1 mt-3">Total calories: <b>{{ total_energy }}</b></p>
                  </v-col>
                </v-row>
              </v-container>
              <v-btn depressed color="primary">
                Log to diary
              </v-btn>
              <v-btn class="ml-3">
                <v-icon color="pink">
                  mdi-heart
                </v-icon>
                Favorite
              </v-btn>
            </v-form>
          </v-card>
        </v-col>
        <v-col align="center">
          <NutritionFacts v-if="product" :product="product"> </NutritionFacts>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import { mapActions, mapGetters, mapState } from 'vuex'

export default {
  data: () => ({
    product: null,
    valid: false,
    amount: 1.0,
    meal: '',
    mealList: ['Breakfast', 'Lunch', 'Dinner', 'Snacks'],
    unit: '100 gr',
    unitList: ['100 gr', '1g', '3.5 ounce', '1 ounce'],
    total_energy: null 
  }),

  async asyncData({ params }) {
    const slug = params.slug
    return { slug }
  },

  created() {
    this.unsubscribe = this.$store.subscribe((mutation, state) => {
      if (mutation.type === 'off/product') {
        if (state.off.product) {
          this.product = state.off.product.product
        }
      }
    })
  },

  beforeDestroy() {
    this.unsubscribe()
  },
  
  watch: {
    amount() {
      // TODO: Transform this.amount function of unitList
      this.total_energy = this.amount * this.product.nutriments.energy_value
    }
  },

  mounted() {
    this.$store.dispatch('off/product', this.slug)
  }
}
</script>
