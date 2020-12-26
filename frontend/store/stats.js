import axios from '~/plugins/axios'
import headers from '~/plugins/headers'

export const state = () => ({
  food: [],
  fluids: [],
  measures: []
})

export const getters = {
  getFoodByMonth: state => (start, end) => {
    var r = new Array()
    state.food.forEach(i => {
      var dt = new Date(i.date)
      if (dt.getMonth() == start.getMonth()) {
        r.push(i)
      }
    })
    return r
  },
  getMeasuresByMonth: state => (start, end) => {
    var r = new Array()
    state.measures.forEach(i => {
      var dt = new Date(i.date)
      if (dt.getMonth() == start.getMonth()) {
        r.push(i)
      }
    })
    return r
  },
  getFluidsByMonth: state => (start, end) => {
    var r = new Array()

    if (state.fluids?.length) {
      state.fluids.forEach(i => {
        var dt = new Date(i.date)
        if (dt.getMonth() == start.getMonth()) {
          r.push(i)
        }
      })
    }
    return r
  },
  getFoodByDay: state => day => {
    var r
    if (state.food?.length) {
      state.food.forEach(i => {
        if (day === i.date) {
          r = i
        }
      })
    }

    if (!r) {
      r = 'N/A'
    }

    return r
  },
  getFoodByDayPieChart: state => day => {
    var r
    if (state.food?.length) {
      state.food.forEach(i => {
        if (day === i.date) {
          r = {
            datasets: [
              {
                data: [i.fat, i.carbs, i.proteins],
                backgroundColor: [
                  'red',
                  'rgb(255,253,169)',
                  'rgb(134,224,102)'
                ],
                fill: true
              }
            ],
            labels: ['Fat', 'Carbs', 'Proteins']
          }
        }
      })
    }

    if (!r) {
      r = undefined
    }

    return r
  },
  lastMeasureTaken: state => day => {
    var m = new Array()
    var d = new Date(day)

    if (state.measures?.length) {
      state.measures.forEach(i => {
        var dt = new Date(i.date)
        if (dt.getTime() <= d.getTime()) {
          m.push(i)
        }
      })
    }

    var r = m.sort(
      (a, b) => new Date(b.date).getTime() - new Date(a.date).getTime()
    )[0]
    if (!r) {
      return 'N/A'
    } else {
      // It won't always be kgs, this has to change
      return r.value + 'kg'
    }
  },
  getFluidsByDay: state => day => {
    var r
    var d = new Date(day)

    if (state.fluids?.length) {
      state.fluids.forEach(i => {
        var dt = new Date(i.date)
        if (dt.getTime() === d.getTime()) {
          r = i.value.toString()
        }
      })
    }

    if (!r) {
      return 'N/A'
    }

    return r + 'ml'
  }
}

export const actions = {
  pullAll({ commit }, payload) {
    axios.get('/stats/fluids', headers()).then(
      response => {
        commit('fluids', response.data)
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
    axios.get('/measures', headers()).then(
      response => {
        commit('measures', response.data)
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
    axios.get('/stats/food_diary', headers()).then(
      response => {
        commit('food', response.data)
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
  measures: (state, payload) => {
    state.measures = payload
  },
  fluids: (state, payload) => {
    state.fluids = payload
  },
  food: (state, payload) => {
    if (payload === undefined || payload.length == 0) {
      state['food'] = []
    } else {
      state.food = payload
    }
  }
}
