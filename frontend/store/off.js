import axios from '~/plugins/axios'
import headers from '~/plugins/headers'

export const state = () => ({
  offSearch: []
})

export const getters = {
  getOffSearch: state => {
    return state.offSearch
  },
  getProduct: state => {
    return state.product
  },
}

export const actions = {
  search({ commit }, payload) {
    axios.get('/off/search', { params: payload }).then(
      response => {
        commit('offSearch', response.data)
      },
      error => {
        commit(
          'snackbar/showMessage',
          { content: error, color: 'red' },
          { root: true }
        )
      }
    )
  },
  product({ commit }, code) {
    axios.get('/off/product/' + code).then(
      response => {
        commit('product', response.data)
      },
      error => {
        commit(
          'snackbar/showMessage',
          { content: error, color: 'red' },
          { root: true }
        )
      }
    )
  },
  reset({ commit }) {
    commit('resetState')
  }
}

export const mutations = {
  offSearch: (state, payload) => {
    state.offSearch = payload
  },
  product: (state, payload) => {
    state.product = payload
  },
  resetState(state) {
    Object.assign(state.offSearch, {})
  }
}
