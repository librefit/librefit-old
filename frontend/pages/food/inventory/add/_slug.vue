<template>
  <div>
    <p class="text-center text-h5">
    Open Food Facts {{ this.slug }}
    </p>

    <NutritionFacts v-if="product" :product="product"> </NutritionFacts>

  </div>
</template>

<script>
import { mapActions, mapGetters, mapState } from 'vuex'

export default {
  data: () => ({
    product: null
  }),

  async asyncData({ params }) {
    const slug = params.slug 
    return { slug }
  },
  
  created() {
    this.unsubscribe = this.$store.subscribe((mutation, state) => {
      if (mutation.type === 'off/product') {
        console.log(`Mutation happening`);

        if (state.off.product) {
          this.product = state.off.product
        }
      }
    });
  },

  beforeDestroy() {
    this.unsubscribe();
  },

  mounted() {
    this.$store.dispatch('off/product', this.slug)
  },
}
</script>
