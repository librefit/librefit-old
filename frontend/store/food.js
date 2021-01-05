import axios from '~/plugins/axios'
import headers from '~/plugins/headers'

export const state = () => ({
  inventory: [],
  diaryDate: [],
  diaryItem: [],
  UploadImages: [],
  loadingStatus: false,
})

export const getters = {
  getAllInventory: (state) => {
    return state.inventory
  },
  getInventoryItem: (state) => (id) => {
    return state.inventory.filter((x) => x.id === Number(id))[0]
  },
  getInventoryItemImages: (state) => (id) => {
    return state.inventory.filter((x) => x.id === Number(id))[0]?.images
  },
  getInventoryByOffCode: (state) => (offCode) => {
    return state.inventory.filter((x) => x.off_code === offCode)[0]?.id
  },
  getFoodDiary: (state) => (id) => {
    var r = new Array()
    state.diaryDate.forEach((i) => {
      var food = {
        id: i.id,
        product_name: i.food_inventory.product_name,
        calories: i.computed_consuption.calories,
        fat_total: i.computed_consuption.fat_total,
        carbs: i.computed_consuption.carbs,
        proteins: i.computed_consuption.proteins,
        sodium: i.computed_consuption.sodium,
        sugars: i.computed_consuption.sugars,
      }
      if (i.meal_type === id) {
        r.push(food)
      }
    })

    return r
  },
  getDiaryItem: (state) => {
    return state.diaryItem
  },
  getLoadingStatus: (state) => {
    return state.loadingStatus
  },
}

export const actions = {
  add2Inventory({ commit }, payload) {
    commit('loadingStatus', true)
    axios.post('/food/inventory', payload, headers()).then(
      (response) => {
        commit('loadingStatus', false)
        commit('append2Inventory', response.data)
      },
      (error) => {
        commit(
          'snackbar/showMessage',
          {
            content: error + ': ' + JSON.stringify(error.response.data),
            color: 'red',
          },
          { root: true }
        )
      }
    )
  },
  updateInventory({ commit }, payload) {
    const id = payload.id
    delete payload.id
    axios.put('/food/inventory/' + id, payload, headers()).then(
      (response) => {
        commit('updateInventory', response.data)
        commit(
          'snackbar/showMessage',
          { content: 'Successfully updated product', color: 'green' },
          { root: true }
        )
      },
      (error) => {
        commit(
          'snackbar/showMessage',
          {
            content: error + ': ' + JSON.stringify(error.response.data),
            color: 'red',
          },
          { root: true }
        )
      }
    )
  },
  updateDiary({ commit, state }) {
    const payload = {
      meal_type: state.diaryItem.meal_type,
      quantity: state.diaryItem.quantity,
      date: new Date(state.diaryItem.date).toISOString().substr(0, 10),
      quantity_unit: state.diaryItem.quantity_unit,
    }

    axios.put('/food/diary/' + state.diaryItem.id, payload, headers()).then(
      (response) => {
        commit(
          'snackbar/showMessage',
          { content: 'Successfully updated diary', color: 'green' },
          { root: true }
        )
      },
      (error) => {
        commit(
          'snackbar/showMessage',
          {
            content: error + ': ' + JSON.stringify(error.response.data),
            color: 'red',
          },
          { root: true }
        )
      }
    )
  },
  deleteInventory({ commit }, payload) {
    axios.delete('/food/inventory/' + payload, headers()).then(
      (response) => {
        commit('deleteInventory', payload)
      },
      (error) => {
        commit(
          'snackbar/showMessage',
          {
            content: error + ': ' + JSON.stringify(error.response.data),
            color: 'red',
          },
          { root: true }
        )
      }
    )
  },
  deleteDiary({ commit }, payload) {
    axios.delete('/food/diary/' + payload, headers()).then(
      (response) => {
        commit('deleteDiary', payload)
      },
      (error) => {
        commit(
          'snackbar/showMessage',
          {
            content: error + ': ' + JSON.stringify(error.response.data),
            color: 'red',
          },
          { root: true }
        )
      }
    )
  },
  pullInventory({ commit }) {
    axios.get('/food/inventory', headers()).then(
      (response) => {
        commit('setInventory', response.data)
      },
      (error) => {
        commit(
          'snackbar/showMessage',
          {
            content: error + ': ' + JSON.stringify(error.response.data),
            color: 'red',
          },
          { root: true }
        )
      }
    )
  },
  uploadImage({ commit, state }, payload) {
    const formData = new FormData()
    formData.append('file', payload)

    axios.post('/upload', formData, headers()).then(
      (response) => {
        console.log(response.data)
        commit('append2Images', response.data.id)
      },
      (error) => {
        commit(
          'snackbar/showMessage',
          {
            content: error + ': ' + JSON.stringify(error.response.data),
            color: 'red',
          },
          { root: true }
        )
      }
    )
  },
  add2Diary({ commit }, payload) {
    console.log(payload)
    commit('loadingStatus', true)
    axios.post('/food/diary', payload, headers()).then(
      (response) => {
        commit('loadingStatus', false)
        $nuxt._router.push('/food/diary')
      },
      (error) => {
        commit(
          'snackbar/showMessage',
          {
            content: error + ': ' + JSON.stringify(error.response.data),
            color: 'red',
          },
          { root: true }
        )
      }
    )
  },
  foodDiary({ commit }, payload) {
    axios.get('/food/diary/' + payload, headers()).then(
      (response) => {
        commit('diaryItem', response.data)
      },
      (error) => {
        commit(
          'snackbar/showMessage',
          {
            content: error + ': ' + JSON.stringify(error.response.data),
            color: 'red',
          },
          { root: true }
        )
      }
    )
  },
  foodDiaryRange({ commit }, payload) {
    axios
      .get('/food/diary', {
        params: { start: payload.start, end: payload.end },
        headers: { Authorization: localStorage.getItem('auth._token.local') },
      })
      .then(
        (response) => {
          commit('diaryByDate', response.data)
        },
        (error) => {
          commit(
            'snackbar/showMessage',
            {
              content: error + ': ' + JSON.stringify(error.response.data),
              color: 'red',
            },
            { root: true }
          )
        }
      )
  },
}

export const mutations = {
  setInventory: (state, payload) => {
    state.inventory = payload
  },
  append2Images: (state, payload) => {
    state.UploadImages.push(payload)
  },
  append2Inventory: (state, payload) => {
    state.inventory.push(payload)
    $nuxt._router.push('/food/inventory')
  },
  updateDiaryForm: (state, payload) => {
    state.diaryItem[payload.element] = payload.value
  },
  updateInventory: (state, payload) => {
    const index = state.inventory.findIndex((x) => x.id == payload.id)
    state.inventory.splice(index, 1, payload)
    $nuxt._router.push('/food/inventory/' + payload.id)
  },
  deleteInventory: (state, payload) => {
    const index = state.inventory.findIndex((x) => x.id == payload.id)
    if (index > -1) {
      state.inventory.splice(index, 1)
    }
    $nuxt._router.push('/food/inventory/')
  },
  updateInventoryItem: (state, payload) => {
    const elementPos = state.inventory
      .map(function (x) {
        return x.id
      })
      .indexOf(payload.id)
    state.inventory[elementPos][payload.element] = payload.value
  },
  deleteDiary: (state, payload) => {
    const index = state.diaryDate.findIndex((x) => x.id == payload)
    if (index > -1) {
      state.diaryDate.splice(index, 1)
    }
  },
  diaryByDate: (state, payload) => {
    // If it doesn't get any payload it means not FoodDiary entry for this date
    if (payload === undefined || payload.length == 0) {
      state['diaryDate'] = []
    } else {
      state.diaryDate = payload
    }
  },
  diaryItem: (state, payload) => {
    state.diaryItem = payload
  },
  loadingStatus(state, value) {
    state.loadingStatus = value
  },
}
