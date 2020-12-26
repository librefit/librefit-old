import axios from '~/plugins/axios'
import headers from '~/plugins/headers'

export const state = () => ({
  allFluids: [],
})

export const getters = {
  getAllFluids: state => {
    return state.allFluids
  }
}

export const actions = {
  loadFluids({ commit }) {
    axios.get('/fluids', headers()).then(
      response => {
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
    state.allFluids = payload
  },
  appendFluid: (state, payload) => {
    state.allFluids.push(payload)
  },
  updateFluid: (state, payload) => {
    const index = state.allFluids.findIndex(x => x.id == payload.id)
    state.allFluids.splice(index, 1, payload)
  },
  removeFluid: (state, payload) => {
    const index = state.allFluids.indexOf(payload)
    state.allFluids.splice(index, 1)
  },
  removeAllFluids: state => {
    state.allFluids = []
  }
}
