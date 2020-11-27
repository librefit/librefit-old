import axios from '~/plugins/axios'
import headers from '~/plugins/headers'

export const state = () => ({
  inventory: [],
  diary: [],
  diaryDate: []
})

export const getters = {
  getAllDiary: state => {
    return state.diary
  },
  getAllInventory: state => {
    return state.inventory
  },
  getInventoryItem: state => id => {
    return state.inventory.filter(x => x.id === Number(id))[0]
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
    axios.post('/food/inventory', payload, headers()).then(
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
  updateInventory({ commit }, payload) {
    const id = payload.id
    delete payload.id
    axios.put('/food/inventory/' + id, payload, headers()).then(
      response => {
        commit('updateInventory', response.data)
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
  pullInventory({ commit }) {
    axios.get('/food/inventory', headers()).then(
      response => {
        commit('setInventory', response.data)
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
  }
}

export const mutations = {
  setInventory: (state, payload) => {
    state.inventory = payload
  },
  append2Inventory: (state, payload) => {
    state.inventory.push(payload)
    $nuxt._router.push('/food/inventory')
  },
  append2Diary: (state, payload) => {
    state.diary.push(payload)
    $nuxt._router.push('/food/diary')
  },
  updateInventory: (state, payload) => {
    const index = state.inventory.findIndex(x => x.id == payload.id)
    state.inventory.splice(index, 1, payload)
    $nuxt._router.push('/food/inventory/' + payload.id)
  },
  updateInventoryItem: (state, payload) => {
    const elementPos = state.inventory
      .map(function(x) {
        return x.id
      })
      .indexOf(payload.id)
    state.inventory[elementPos][payload.element] = payload.value
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
