import axios from '~/plugins/axios'
import headers from '~/plugins/headers'

export const state = () => ({
  inventory: [],
  diary: []
})

export const getters = {
  getAllDiary: state => {
    return state.diary
  },
  getAllInventory: state => {
    return state.inventory
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
  }
}

export const mutations = {
  append2Inventory: (state, payload) => {
    state.inventory.push(payload)
  },
  append2Diary: (state, payload) => {
    state.diary.push(payload)
  }
}
