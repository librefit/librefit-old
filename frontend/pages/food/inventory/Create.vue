<template>
  <div>
    <v-tabs v-model="tab" background-color="transparent" color="basil" grow>
      <v-tab> Search </v-tab>
      <v-tab> Product </v-tab>
    </v-tabs>

    <v-tabs-items v-model="tab">
      <v-tab-item>
        <v-card color="basil" flat>
          <v-form @submit="searchOff" onSubmit="return false;">
            <v-container>
              <v-text-field
                v-model="searchText"
                label="Enter food name or brand to search in OpenFoodFacts"
                required
              ></v-text-field>
              <v-btn @click="searchOff"> Search </v-btn>
            </v-container>
          </v-form>

          <v-overlay v-if="getLoadingStatus" class="text-center">
            <v-progress-circular :size="70" indeterminate></v-progress-circular>
          </v-overlay>

          <v-list three-line>
            <template v-for="(item, index) in getOffSearch.products">
              <v-divider></v-divider>

              <v-list-item>
                <v-list-item-avatar>
                  <v-img :src="item.image_front_url"></v-img>
                </v-list-item-avatar>

                <v-list-item-content>
                  <v-list-item-title
                    v-html="item.product_name"
                  ></v-list-item-title>
                  <v-list-item-subtitle>
                    <div v-if="item.nutriments['energy-kcal']">
                      {{ item.nutriments['energy-kcal'] }} cals
                    </div>
                  </v-list-item-subtitle>
                  <v-spacer></v-spacer>
                </v-list-item-content>

                <v-btn
                  class="ma-2"
                  color="info"
                  @click="pullProduct(item.code)"
                >
                  Add
                </v-btn>

                <v-menu offset-y rounded>
                  <template v-slot:activator="{ on, attrs }">
                    <v-btn icon v-bind="attrs" v-on="on">
                      <v-icon>mdi-dots-vertical</v-icon>
                    </v-btn>
                  </template>
                  <v-list>
                    <v-list-item :href="item.url" target="_blank">
                      <v-list-item-title
                        >Check in OpenFoodFacts</v-list-item-title
                      >
                    </v-list-item>
                    <v-list-item>
                      <v-list-item-title>Add to favorites</v-list-item-title>
                    </v-list-item>
                  </v-list>
                </v-menu>
              </v-list-item>
            </template>
          </v-list>
          <div class="text-center">
            <v-pagination
              v-model="page"
              v-if="getOffSearch.products"
              :length="Math.ceil(getOffSearch.count / getOffSearch.page_size)"
              :total-visible="7"
              class="my-4"
              @input="next"
              circle
            ></v-pagination>
          </div>
          <div class="py-3" />
        </v-card>
      </v-tab-item>
      <v-tab-item>
        <v-container class="grey lighten-5">
          <p class="text-center text-h5" v-if="product">
            {{ product.product_name }}
          </p>

          <v-dialog
            transition="dialog-bottom-transition"
            max-width="600"
            v-model="dialog"
          >
            <template v-slot:default="dialog">
              <v-card>
                <v-toolbar color="primary" dark
                  >Product barcode is already in your inventory</v-toolbar
                >
                <v-card-text>
                  <div class="text-h6 pa-12">
                    This product barcode is already registered in your inventory
                  </div>
                </v-card-text>
                <v-card-actions class="justify-end">
                  <v-btn text :href="'/food/inventory/' + existingInInventory">Take me to the product</v-btn>
                  <v-btn text @click="dialog.value = false">Ignore</v-btn>
                </v-card-actions>
              </v-card>
            </template>
          </v-dialog>

          <v-row>
            <v-col>
              <v-card class="pa-2" outlined tile>
                <v-form>
                  <v-container>
                    <v-row>
                      <v-col md="8">
                        <v-text-field
                          v-model="product.product_name"
                          label="Product Name"
                        ></v-text-field>
                      </v-col>
                    </v-row>
                    <v-row>
                      <v-col md="8">
                        <v-text-field
                          v-model="product.description"
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
                          v-model.number="product.calories"
                          label="Calories"
                        ></v-text-field>
                      </v-col>
                      <v-col md="4">
                        <v-text-field
                          v-model.number="product.sodium"
                          label="Sodium"
                        ></v-text-field>
                      </v-col>
                      <v-col md="4">
                        <v-text-field
                          v-model.number="product.potassium"
                          label="Potassium"
                        ></v-text-field>
                      </v-col>
                    </v-row>
                    <v-row>
                      <v-col md="4">
                        <v-text-field
                          v-model.number="product.fibers"
                          label="Dietary Fiber"
                        ></v-text-field>
                      </v-col>
                      <v-col md="4">
                        <v-text-field
                          v-model.number="product.sugars"
                          label="Sugars"
                        ></v-text-field>
                      </v-col>
                      <v-col md="4">
                        <v-text-field
                          v-model.number="product.proteins"
                          label="Proteins"
                        ></v-text-field>
                      </v-col>
                    </v-row>
                    <v-row>
                      <v-col md="4">
                        <v-text-field
                          v-model.number="product.fat_cholesterol"
                          label="Cholesterol"
                        ></v-text-field>
                      </v-col>
                      <v-col md="4">
                        <v-text-field
                          v-model.number="product.fat_saturated"
                          label="Fat Saturated"
                        ></v-text-field>
                      </v-col>
                      <v-col md="4">
                        <v-text-field
                          v-model.number="product.fat_polyunsaturated"
                          label="Fat Polyunsaturated"
                        ></v-text-field>
                      </v-col>
                    </v-row>
                    <v-row>
                      <v-col md="4">
                        <v-text-field
                          v-model.number="product.fat_monounsaturated"
                          label="Fat Monounsaturated"
                        ></v-text-field>
                      </v-col>
                      <v-col md="4">
                        <v-text-field
                          v-model.number="product.fat_trans"
                          label="Fat Trans"
                        ></v-text-field>
                      </v-col>
                      <v-col md="4">
                        <v-text-field
                          v-model.number="product.carbs"
                          label="Carbohydrates"
                        ></v-text-field>
                      </v-col>
                    </v-row>
                  </v-container>

                  <v-btn depressed color="primary" @click="save">
                    Create
                  </v-btn>
                  <v-btn depressed color="primary" @click="save">
                    Create & Log To Diary
                  </v-btn>
                </v-form>
              </v-card>
            </v-col>
            <v-col align="center">
              <NutritionFactsInventory v-if="product" :product="product" />
              <v-carousel
                hide-delimiter-background
                delimiter-icon="mdi-minus"
                v-if="getProductImages"
                height="400"
                class="mt-4"
                cycle
                interval="4000"
              >
                <v-carousel-item
                  v-for="(item, i) in getProductImages"
                  v-if="item"
                  :key="i"
                >
                  <img
                    :src="item"
                    style="width: 200px; height: auto"
                    alt="image-from-openfoodfacts"
                  />
                </v-carousel-item>
              </v-carousel>
            </v-col>
          </v-row>
        </v-container>
      </v-tab-item>
    </v-tabs-items>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'

export default {
  data: () => ({
    page: 1,
    searchText: '',
    product: {},
    tab: null,
    images: [],
    dialog: false,
    existingInInventory: '',
  }),

  mounted() {
    this.$store.commit('off/resetState')
    this.$store.commit('off/resetProduct')
  },

  computed: {
    ...mapGetters('off', [
      'getOffSearch',
      'getLoadingStatus',
      'getProductImages',
    ]),
  },

  methods: {
    ...mapActions('off', ['search', 'offProduct']),

    save() {
      this.$store.dispatch('food/add2Inventory', this.product)
    },

    uploadImage() {
      this.$store.dispatch('food/uploadImage', this.image)
    },

    pullProduct(item) {
      if (this.$store.getters['food/getInventoryByOffCode'](item)) {
        this.existingInInventory = this.$store.getters['food/getInventoryByOffCode'](item)
        this.dialog = true
      }
      let unsubscribe = null
      this.$store.commit('off/resetProduct')
      this.offProduct(item)
      unsubscribe = this.$store.subscribe((mutation, state) => {
        if (mutation.type === 'off/product') {
          this.product = this.$store.getters['off/getProduct']
          unsubscribe()
          this.tab = 1
        }
      })
    },

    searchOff() {
      this.search({
        search_terms: encodeURIComponent(this.searchText),
        page: this.page,
      })
    },

    next(page) {
      this.page = page
      this.search({
        search_terms: encodeURIComponent(this.searchText),
        page: page,
      })
    },
  },
}
</script>
