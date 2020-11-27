import axios from '~/plugins/axios'
import headers from '~/plugins/headers'

export const state = () => ({
  settings: [],
})

export const getters = {
  getSettings: state => {
    return state.settings
  },
}

export const actions = {
  pullSettings({ commit }) {
    axios.get('/user/info', headers()).then(
      response => {
        commit('setSettings', response.data)
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
}

export const mutations = {
  setSettings: (state, payload) => {
    state.settings = payload
  },
  setUsername: (state, payload) => {
    state.settings.username = payload
  },
  setFullName: (state, payload) => {
    state.settings.user_settings.full_name = payload
  },
  setEmail: (state, payload) => {
    state.settings.user_settings.email = payload
  },
  setDob: (state, payload) => {
    state.settings.user_settings.birthday = payload
  },
  setUseMetric: (state, payload) => {
    state.settings.user_settings.use_metric = payload
  },
}
