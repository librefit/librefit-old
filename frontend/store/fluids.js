import axios from '~/plugins/axios'
import headers from '~/plugins/headers'

export const state = () => ({
  byDay: [],
})

export const getters = {
  getFluids: state => {
    return state.byDay
  }
}

export const actions = {
  addFluid({ commit }, payload) {
    axios.post('/fluid', payload, headers()).then(
      response => {
        commit('appendFluid', response.data)
      },
      error => {
        commit(
          'snackbar/showMessage',
          {
            content: error + ': ' + JSON.stringify(error.response.data),
            color: 'red'
          },
          { root: true }
        )
      }
    )
  },
  getFluidsByRange({ commit }, payload) {
    axios.post('/fluids', payload, headers()).then(
      response => {
        commit('removeAllFluids')
        commit('initFluids', response.data)
      },
      error => {
        commit(
          'snackbar/showMessage',
          {
            content: error + ': ' + JSON.stringify(error.response.data),
            color: 'red'
          },
          { root: true }
        )
      }
    )
  },
  deleteFluid: ({ commit }, payload) => {
    axios.delete('/fluids/' + payload.id, headers()).then(
      () => {
        commit('removeFluid', payload)
      },
      error => {
        commit(
          'snackbar/showMessage',
          {
            content: error + ': ' + JSON.stringify(error.response.data),
            color: 'red'
          },
          { root: true }
        )
      }
    )
  },
  updateFluid: ({ commit }, payload) => {
    axios.put('/fluids/' + payload.id, payload, headers()).then(
      response => {
        commit('updateFluid', response.data)
      },
      error => {
        commit(
          'snackbar/showMessage',
          {
            content: error + ': ' + JSON.stringify(error.response.data),
            color: 'red'
          },
          { root: true }
        )
      }
    )
  }
}

export const mutations = {
  initFluids: (state, payload) => {
    state.byDay = payload
  },
  appendFluid: (state, payload) => {
    state.byDay.push(payload)
  },
  updateFluid: (state, payload) => {
    const index = state.byDay.findIndex(x => x.id == payload.id)
    state.byDay.splice(index, 1, payload)
  },
  removeFluid: (state, payload) => {
    const index = state.byDay.indexOf(payload)
    state.byDay.splice(index, 1)
  },
  removeAllFluids: state => {
    state.byDay = []
  }
}
