import axios from '~/plugins/axios'
import headers from '~/plugins/headers'

export const state = () => ({
  food: []
})

export const getters = {}

export const actions = {
  food({ commit }, payload) {
    axios
      .get('/stats/food_diary', {
        params: { start: payload.start, end: payload.end },
        headers: { Authorization: localStorage.getItem('auth._token.local') }
      })
      .then(
        response => {
          commit('food', response.data)
        },
        error => {
          commit(
            'snackbar/showMessage',
            { content: error, color: 'red' },
            { root: true }
          )
        }
      )
  }
}

export const mutations = {
  food: (state, payload) => {
    if (payload === undefined || payload.length == 0) {
      state['food'] = []
    } else {
      state.food = payload
    }
  }
}
