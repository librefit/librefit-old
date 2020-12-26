<template>
  <div>
    <p class="text-center text-h5" v-if="product">{{ product.product_name }}</p>

    <v-container class="grey lighten-5">
      <v-row>
        <v-col>
          <v-card class="pa-2" outlined tile>
            <v-form>
              <v-container>
                <v-row>
                  <v-col md="8">
                    <v-text-field
                      v-model="productName"
                      label="Product Name"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row>
                  <v-col md="8">
                    <v-text-field
                      v-model="description"
                      label="Description (optional)"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row>
                  <p class="pl-3 text-h6">Nutrition Facts</p>
                </v-row>
                <v-row>
                  <v-col md="4">
                    <v-text-field
                      v-model="calories"
                      label="Calories"
                    ></v-text-field>
                  </v-col>
                  <v-col md="4">
                    <v-text-field
                      v-model="sodium"
                      label="Sodium"
                    ></v-text-field>
                  </v-col>
                  <v-col md="4">
                    <v-text-field
                      v-model="potassium"
                      label="Potassium"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row>
                  <v-col md="4">
                    <v-text-field
                      v-model="fibers"
                      label="Dietary Fiber"
                    ></v-text-field>
                  </v-col>
                  <v-col md="4">
                    <v-text-field
                      v-model="sugars"
                      label="Sugars"
                    ></v-text-field>
                  </v-col>
                  <v-col md="4">
                    <v-text-field
                      v-model="proteins"
                      label="Proteins"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row>
                  <v-col md="4">
                    <v-text-field
                      v-model="fat_cholesterol"
                      label="Cholesterol"
                    ></v-text-field>
                  </v-col>
                  <v-col md="4">
                    <v-text-field
                      v-model="fat_saturated"
                      label="Fat Saturated"
                    ></v-text-field>
                  </v-col>
                  <v-col md="4">
                    <v-text-field
                      v-model="fat_polyunsaturated"
                      label="Fat Polyunsaturated"
                    ></v-text-field>
                  </v-col>
                </v-row>
                <v-row>
                  <v-col md="4">
                    <v-text-field
                      v-model="fat_monounsaturated"
                      label="Fat Monounsaturated"
                    ></v-text-field>
                  </v-col>
                  <v-col md="4">
                    <v-text-field
                      v-model="fat_trans"
                      label="Fat Trans"
                    ></v-text-field>
                  </v-col>
                  <v-col md="4"> 
                    <v-text-field
                      v-model="carbs"
                      label="Carbohydrates"
                    ></v-text-field>
                  </v-col>
                </v-row>
              </v-container>

              <v-btn depressed color="primary" @click="save">
                Update
              </v-btn>
              <v-btn depressed color="red" @click="remove">
                Delete
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
          <NutritionFactsInventory v-if="product" :product="product">
          </NutritionFactsInventory>
          <p v-if="offCode" class="text-subtitle-1 mt-4">
            No the right values?
            <a
              :href="
                'https://world.openfoodfacts.org/cgi/product.pl?type=edit&code=' +
                  offCode
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
  data: () => ({}),

  async asyncData({ params }) {
    const slug = params.slug
    return { slug }
  },

  mounted() {
    this.$store.dispatch('food/pullInventory')
  },

  computed: {
    product() {
      return this.$store.getters['food/getInventoryItem'](this.slug)
    },
    offCode: {
      get() {
        if (this.$store.getters['food/getInventoryItem'](this.slug)) {
          return this.product.off_code
        }
      },
    },
    productName: {
      get() {
        if (this.$store.getters['food/getInventoryItem'](this.slug)) {
          return this.product.product_name
        }
      },
      set(value) {
        this.$store.commit('food/updateInventoryItem', {
          id: this.product.id,
          value: value,
          element: 'product_name'
        })
      }
    },
    description: {
      get() {
        if (this.$store.getters['food/getInventoryItem'](this.slug)) {
          return this.product.description
        }
      },
      set(value) {
        this.$store.commit('food/updateInventoryItem', {
          id: this.product.id,
          value: value,
          element: 'description'
        })
      }
    },
    calories: {
      get() {
        if (this.$store.getters['food/getInventoryItem'](this.slug)) {
          return this.product.calories
        }
      },
      set(value) {
        this.$store.commit('food/updateInventoryItem', {
          id: this.product.id,
          value: Number(value),
          element: 'calories'
        })
      }
    },
    sodium: {
      get() {
        if (this.$store.getters['food/getInventoryItem'](this.slug)) {
          return this.product.sodium
        }
      },
      set(value) {
        this.$store.commit('food/updateInventoryItem', {
          id: this.product.id,
          value: Number(value),
          element: 'sodium'
        })
      }
    },
    potassium: {
      get() {
        if (this.$store.getters['food/getInventoryItem'](this.slug)) {
          return this.product.potassium
        }
      },
      set(value) {
        this.$store.commit('food/updateInventoryItem', {
          id: this.product.id,
          value: Number(value),
          element: 'potassium'
        })
      }
    },
    fibers: {
      get() {
        if (this.$store.getters['food/getInventoryItem'](this.slug)) {
          return this.product.fibers
        }
      },
      set(value) {
        this.$store.commit('food/updateInventoryItem', {
          id: this.product.id,
          value: Number(value),
          element: 'fibers'
        })
      }
    },
    sugars: {
      get() {
        if (this.$store.getters['food/getInventoryItem'](this.slug)) {
          return this.product.sugars
        }
      },
      set(value) {
        this.$store.commit('food/updateInventoryItem', {
          id: this.product.id,
          value: Number(value),
          element: 'sugars'
        })
      }
    },
    proteins: {
      get() {
        if (this.$store.getters['food/getInventoryItem'](this.slug)) {
          return this.product.proteins
        }
      },
      set(value) {
        this.$store.commit('food/updateInventoryItem', {
          id: this.product.id,
          value: Number(value),
          element: 'proteins'
        })
      }
    },
    fat_cholesterol: {
      get() {
        if (this.$store.getters['food/getInventoryItem'](this.slug)) {
          return this.product.fat_cholesterol
        }
      },
      set(value) {
        this.$store.commit('food/updateInventoryItem', {
          id: this.product.id,
          value: Number(value),
          element: 'fat_cholesterol'
        })
      }
    },
    fat_saturated: {
      get() {
        if (this.$store.getters['food/getInventoryItem'](this.slug)) {
          return this.product.fat_saturated
        }
      },
      set(value) {
        this.$store.commit('food/updateInventoryItem', {
          id: this.product.id,
          value: Number(value),
          element: 'fat_saturated'
        })
      }
    },
    fat_trans: {
      get() {
        if (this.$store.getters['food/getInventoryItem'](this.slug)) {
          return this.product.fat_trans
        }
      },
      set(value) {
        this.$store.commit('food/updateInventoryItem', {
          id: this.product.id,
          value: Number(value),
          element: 'fat_trans'
        })
      }
    },
    fat_polyunsaturated: {
      get() {
        if (this.$store.getters['food/getInventoryItem'](this.slug)) {
          return this.product.fat_polyunsaturated
        }
      },
      set(value) {
        this.$store.commit('food/updateInventoryItem', {
          id: this.product.id,
          value: Number(value),
          element: 'fat_polyunsaturated'
        })
      }
    },
    fat_monounsaturated: {
      get() {
        if (this.$store.getters['food/getInventoryItem'](this.slug)) {
          return this.product.fat_monounsaturated
        }
      },
      set(value) {
        this.$store.commit('food/updateInventoryItem', {
          id: this.product.id,
          value: Number(value),
          element: 'fat_monounsaturated'
        })
      }
    },
    carbs: {
      get() {
        if (this.$store.getters['food/getInventoryItem'](this.slug)) {
          return this.product.carbs
        }
      },
      set(value) {
        this.$store.commit('food/updateInventoryItem', {
          id: this.product.id,
          value: Number(value),
          element: 'carbs'
        })
      }
    }
  },

  methods: {
    ...mapActions('food', ['updateInventory', 'deleteInventory']),

    save() {
      this.updateInventory(this.product)
    },
    
    remove() {
      this.deleteInventory(this.product.id)
    }
  }
}
</script>
