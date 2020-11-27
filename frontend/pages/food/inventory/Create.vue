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
                      v-model="product.carbs"
                      label="Carbohydrates"
                    ></v-text-field>
                  </v-col>
                </v-row>
              </v-container>

              <v-btn depressed color="primary" @click="save">
                Create
              </v-btn>
            </v-form>
          </v-card>
        </v-col>
        <v-col align="center">
          <NutritionFactsInventory v-if="product" :product="product">
          </NutritionFactsInventory>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import { mapActions } from 'vuex'

export default {
  data: () => ({
    product: {}
  }),

  methods: {
    ...mapActions('food', ['add2Inventory']), 

    save() {
      this.add2Inventory(this.product)
    }
  }
}
</script>
