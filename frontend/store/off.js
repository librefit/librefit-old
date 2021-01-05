import axios from '~/plugins/axios'
import headers from '~/plugins/headers'

export const state = () => ({
  offSearch: [],
  product: {},
  loadingStatus: false,
})

export const getters = {
  getOffSearch: (state) => {
    return state.offSearch
  },
  getProduct: (state) => {
    return {
      calories: state.product.product.nutriments['energy-kj_100g'] || null,
      sodium: state.product.product.nutriments['sodium_100g'] || null,
      potassium: state.product.product.nutriments['potassium_100g'] || null,
      fibers: state.product.product.nutriments['fiber_100g'] || null,
      sugars: state.product.product.nutriments['sugars_100g'] || null,
      proteins: state.product.product.nutriments['proteins_100g'] || null,
      fat_cholesterol:
        state.product.product.nutriments['cholesterol_100g'] || null,
      fat_saturated:
        state.product.product.nutriments['saturated-fat_100g'] || null,
      fat_polyunsaturated:
        state.product.product.nutriments['polyunsaturated-fat_100g'] || null,
      fat_monounsaturated:
        state.product.product.nutriments['monounsaturated-fat_100g'] || null,
      fat_trans: state.product.product.nutriments['trans-fat_100g'] || null,
      product_name: state.product.product.product_name,
      off_code: state.product.code,
      images: state.product.images,
    }
  },
  getProductImages: (state) => {
    return state.product.images
  },
  getLoadingStatus: (state) => {
    return state.loadingStatus
  },
}

export const actions = {
  search({ commit }, payload) {
    commit('loadingStatus', true)
    axios.get('/off/search', { params: payload }).then(
      (response) => {
        commit('loadingStatus', false)
        commit('offSearch', response.data)
      },
      (error) => {
        commit(
          'snackbar/showMessage',
          {
            content: error + ': ' + JSON.stringify(error.response.data),
            color: 'red',
          },
          { root: true }
        )
      }
    )
  },
  offProduct({ commit }, code) {
    commit('loadingStatus', true)
    axios.get('/off/product/' + code).then(
      (response) => {
        commit('loadingStatus', false)
        commit('product', response.data)
      },
      (error) => {
        commit(
          'snackbar/showMessage',
          {
            content: error + ': ' + JSON.stringify(error.response.data),
            color: 'red',
          },
          { root: true }
        )
      }
    )
  },
  reset({ commit }) {
    commit('resetState')
  },
}

export const mutations = {
  offSearch: (state, payload) => {
    state.offSearch = payload
  },
  product: (state, payload) => {
    var images = []
    state.product = payload
    if (payload.product?.image_front_url) {
      images.push(payload.product?.image_front_url)
    }
    if (payload.product?.image_ingredients_url) {
      images.push(payload.product?.image_ingredients_url)
    }
    if (payload.product?.image_nutrition_url) {
      images.push(payload.product?.image_nutrition_url)
    }
    if (payload.product?.image_url) {
      images.push(payload.product?.image_url)
    }
    state.product.images = images
  },
  resetState(state) {
    state.offSearch = []
  },
  resetProduct(state) {
    state.product = {}
  },
  loadingStatus(state, value) {
    state.loadingStatus = value
  },
}
