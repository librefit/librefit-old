import axios from '~/plugins/axios'
import headers from '~/plugins/headers'

export const state = () => ({
  allMeasures: []
})

export const getters = {
  getAllMeasures: state => {
    return state.allMeasures
  },
  getByMonth: state => (start, end) => {
    var r = new Array()
    state.allMeasures.forEach(i => {
      var dt = new Date(i.date)
      if (dt.getMonth() + 1 == start.month) {
        r.push(i)
      }
    })
    return r
  },
  lastMeasureByDay: state => day => {
    var m = new Array()
    var d = new Date(day)

    state.allMeasures.forEach(i => {
      var dt = new Date(i.date)
      if (dt.getTime() <= d.getTime()) {
        m.push(i)
      }
    })

    var r = m.sort((a,b) => new Date(b.date).getTime() - new Date(a.date).getTime())[0];
    if (!r) {
      r = 'N/A'
    }

    // It won't always be kgs, this has to change
    return r.value + ' kg'
  },
  getLineChartOptions: () => {
    return {
      responsive: true,
      maintainAspectRatio: false
    }
  },
  getLineChartData: state => {
    return {
      labels: state.allMeasures.map(function(a) {
        return a.date
      }),
      datasets: [
        {
          label: 'Weight over time',
          borderColor: 'green',
          data: state.allMeasures.map(function(a) {
            return a.value
          })
        }
      ]
    }
  }
}

export const actions = {
  loadMeasures({ commit }) {
    axios.get('/measures', headers()).then(
      response => {
        commit('initMeasures', response.data)
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
  addMeasure({ commit }, payload) {
    axios.post('/measures', payload, headers()).then(
      response => {
        commit('appendMeasure', response.data)
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
  deleteMeasure: ({ commit }, payload) => {
    axios.delete('/measures/' + payload.id, headers()).then(
      () => {
        commit('removeMeasure', payload)
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
  updateMeasure: ({ commit }, payload) => {
    axios.put('/measures/' + payload.id, payload, headers()).then(
      response => {
        commit('updateMeasure', response.data)
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
  initMeasures: (state, payload) => {
    state.allMeasures = payload
  },
  appendMeasure: (state, payload) => {
    state.allMeasures.push(payload)
  },
  updateMeasure: (state, payload) => {
    const index = state.allMeasures.findIndex(x => x.id == payload.id)
    state.allMeasures.splice(index, 1, payload)
  },
  removeMeasure: (state, payload) => {
    const index = state.allMeasures.indexOf(payload)
    state.allMeasures.splice(index, 1)
  }
}
