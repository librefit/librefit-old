import axios from '~/plugins/axios'
import headers from '~/plugins/headers'

export const state = () => ({
  allFluids: [],
  stats: []
})

export const getters = {
  getAllFluids: state => {
    return state.allFluids
  },
  getByMonth: state => (start, end) => {
    var r = new Array()
    state.stats.forEach(i => {
      var dt = new Date(i.date)
      if (dt.getMonth() + 1 == start.month) {
        r.push(i)
      }
    })
    return r
  },
  totalByDay: state => day => {
    var r
    var d = new Date(day)
    state.stats.forEach(i => {
      var dt = new Date(i.date)
      if (dt.getTime() === d.getTime()) {
        r = i.value.toString()
      }
    })

    if (!r) {
      r = 'N/A'
    }

    return r
  }
  // getChartOptions: state => {
  //   return {
  //     chart: {
  //       id: "Weight Chart"
  //     },
  //     xaxis: {
  //       categories: state.Fluids.map(function(a) {
  //         return a.date;
  //       })
  //     }
  //   };
  // },
  // getChartSeries: state => {
  //   return {
  //     name: "series-1",
  //     data: state.Fluids.map(function(a) {
  //       return a.value;
  //     })
  //   };
  // }
}

export const actions = {
  loadFluids({ commit }) {
    axios.get('/fluids', headers()).then(
      response => {
        commit('initFluids', response.data)
      },
      error => {
        commit('showMessage', { content: error, color: 'red' })
        commit(
          'snackbar/showMessage',
          { content: error, color: 'red' },
          { root: true }
        )
      }
    )
  },
  statsFluids({ commit }) {
    axios.get('/stats/fluids', headers()).then(
      response => {
        commit('statsFluids', response.data)
      },
      error => {
        commit('showMessage', { content: error, color: 'red' })
        commit(
          'snackbar/showMessage',
          { content: error, color: 'red' },
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
          { content: error, color: 'red' },
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
          { content: error, color: 'red' },
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
          { content: error, color: 'red' },
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
          { content: error, color: 'red' },
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
  statsFluids: (state, payload) => {
    state.stats = payload
  },
  removeAllFluids: state => {
    state.allFluids = []
  }
}
