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
  getFoodDiary: state => {
    return state.diaryDate
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
  }
}

export const mutations = {
  append2Inventory: (state, payload) => {
    state.inventory.push(payload)
  },
  append2Diary: (state, payload) => {
    state.diary.push(payload)
  },
  diaryByDate: (state, payload) => {
    // If it doesn't get any payload it means not FoodDiary entry for this date
    if (payload === undefined || payload.length == 0) {
      state["diaryDate"] = []
    } else {
      state.diaryDate = payload
    }
  }
}
