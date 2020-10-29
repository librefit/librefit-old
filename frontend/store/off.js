import axios from '~/plugins/axios'
import headers from '~/plugins/headers'

export const state = () => ({
  offSearch: []
})

export const getters = {
  getOffSearch: state => {
    return state.offSearch
  }
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
}

export const mutations = {
  offSearch: (state, payload) => {
    state.offSearch = payload
  },
}


