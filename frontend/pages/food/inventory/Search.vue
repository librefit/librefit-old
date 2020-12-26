<template>
  <v-card class="fill-height" fluid>
    <v-form>
      <v-container>
        <v-text-field
          v-model="food"
          label="Please enter food name or brand"
          required
        ></v-text-field>
        <v-btn @click="searchOff">
          Search
        </v-btn>
      </v-container>
    </v-form>

    <v-overlay v-if="getLoadingStatus" class="text-center">
      <v-progress-circular
        :size="70"
        indeterminate
      ></v-progress-circular>
    </v-overlay>

    <v-list three-line>
      <template v-for="(item, index) in getOffSearch.products">
        <v-divider></v-divider>

        <v-list-item>
          <v-list-item-avatar>
            <v-img :src="item.image_front_url"></v-img>
          </v-list-item-avatar>

          <v-list-item-content>
            <v-list-item-title v-html="item.product_name"></v-list-item-title>
            <v-list-item-subtitle>
              <div v-if="item.nutriments['energy-kcal']">
                {{ item.nutriments['energy-kcal'] }} cals
              </div>
            </v-list-item-subtitle>
            <v-spacer></v-spacer>
          </v-list-item-content>

          <v-btn class="ma-2" color="info" :to="'add/' + item.code">
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
                <v-list-item-title>Check in OpenFoodFacts</v-list-item-title>
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
</template>

<script>
import { mapActions, mapGetters } from 'vuex'

export default {
  data: () => ({
    food: '',
    page: 1
  }),

  computed: {
    ...mapGetters('off', ['getOffSearch', 'getLoadingStatus'])
  },

  mounted() {
    this.$store.commit('off/resetState')
  },

  methods: {
    ...mapActions('off', ['search']),

    searchOff() {
      this.search({
        search_terms: encodeURIComponent(this.food),
        page: this.page
      })
    },

    next(page) {
      this.page = page
      this.search({ search_terms: encodeURIComponent(this.food), page: page })
    }
  }
}
</script>
