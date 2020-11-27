<template>
  <div>
    <p class="text-center text-h4">
      Food Inventory
    </p>
    <v-bottom-navigation>
      <v-btn to="inventory/create">
        <span>Create</span>
        <v-icon>mdi-plus-circle</v-icon>
      </v-btn>

      <v-btn>
        <span>Favorites</span>
        <v-icon>mdi-heart</v-icon>
      </v-btn>
    </v-bottom-navigation>

    <v-card class="fill-height" fluid>
      <v-list three-line>
        <template v-for="(item, index) in getAllInventory">
          <v-divider></v-divider>

          <v-list-item>
            <v-list-item-avatar>
              <v-img :src="item.image_front_url"></v-img>
            </v-list-item-avatar>

            <v-list-item-content>
              <v-list-item-title v-html="item.product_name"></v-list-item-title>
              <v-list-item-subtitle>
                <div>
                  {{ item }}
                </div>
              </v-list-item-subtitle>
              <v-spacer></v-spacer>
            </v-list-item-content>

            <v-btn class="ma-2" :to="'inventory/' + item.id">
              Edit
            </v-btn>
          </v-list-item>
        </template>
      </v-list>
      <!-- <div class="text-center"> -->
      <!--   <v-pagination -->
      <!--     v-model="page" -->
      <!--     v-if="getAllInventory" -->
      <!--     :length="Math.ceil(getAllInventory.count / getAllInventory.page_size)" -->
      <!--     :total-visible="7" -->
      <!--     class="my-4" -->
      <!--     @input="next" -->
      <!--     circle -->
      <!--   ></v-pagination> -->
      <!-- </div> -->
      <div class="py-3" />
    </v-card>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'

export default {
  data: () => ({
    page: 1
  }),

  computed: {
    ...mapGetters('food', ['getAllInventory'])
  },

  mounted() {
    this.$store.dispatch('food/pullInventory')
  },

  methods: {
    // ...mapActions('food', ['getAllInventory']),
    // searchOff() {
    //   this.search({ search_terms: this.food, page: this.page })
    // },
    // next(page) {
    //   this.page = page
    //   this.search({ search_terms: this.food, page: page })
    // }
  }
}
</script>
