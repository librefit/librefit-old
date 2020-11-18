import axios from '~/plugins/axios'
import headers from '~/plugins/headers'

export const state = () => ({
  inventory: [],
  diary: [],
  diaryDate: [],
  statsByDate: []
})

export const getters = {
  getAllDiary: state => {
    return state.diary
  },
  getAllInventory: state => {
    return state.inventory
  },
  getPieChartDataSet: state => day => {
    var r
    if (state.statsByDate?.length) {
      state.statsByDate.forEach(i => {
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
  getStats: state => day => {
    var r
    if (state.statsByDate?.length) {
      state.statsByDate.forEach(i => {
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
  getFoodDiary: state => id => {
    var r = new Array()
    state.diaryDate.forEach(i => {
      var food = {
        product_name: i.food_inventory.product_name,
        calories: i.computed_consuption.calories,
        fat_total: i.computed_consuption.fat_total,
        carbs: i.computed_consuption.carbs,
        proteins: i.computed_consuption.proteins,
        sodium: i.computed_consuption.sodium,
        sugars: i.computed_consuption.sugars
      }
      if (i.meal_type === id) {
        r.push(food)
      }
    })

    return r
  }
}

export const actions = {
  add2Inventory({ commit }, payload) {
    axios.post('/inventory', payload, headers()).then(
      response => {
        commit('append2Inventory', response.data)
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
  add2Diary({ commit }, payload) {
    axios.post('/food/diary', payload, headers()).then(
      response => {
        commit('append2Diary', response.data)
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
  foodDiary({ commit }, payload) {
    axios
      .get('/food/diary', {
        params: { start: payload.start, end: payload.end },
        headers: { Authorization: localStorage.getItem('auth._token.local') }
      })
      .then(
        response => {
          commit('diaryByDate', response.data)
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
  foodStats({ commit }, payload) {
    axios
      .get('/stats/food_diary', {
        params: { start: payload.start, end: payload.end },
        headers: { Authorization: localStorage.getItem('auth._token.local') }
      })
      .then(
        response => {
          commit('statsByDate', response.data)
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
  append2Inventory: (state, payload) => {
    state.inventory.push(payload)
  },
  append2Diary: (state, payload) => {
    state.diary.push(payload)
    $nuxt._router.push('/food/diary')
  },
  statsByDate: (state, payload) => {
    // If it doesn't get any payload it means not FoodDiary entry for this date
    if (payload === undefined || payload.length == 0) {
      state['statsByDate'] = []
    } else {
      state.statsByDate = payload
    }
  },
  diaryByDate: (state, payload) => {
    // If it doesn't get any payload it means not FoodDiary entry for this date
    if (payload === undefined || payload.length == 0) {
      state['diaryDate'] = []
    } else {
      state.diaryDate = payload
    }
  }
}
