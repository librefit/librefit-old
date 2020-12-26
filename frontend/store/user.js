import axios from '~/plugins/axios'
import headers from '~/plugins/headers'

export const state = () => ({
  settings: [],
  username: null,
  password: null,
  loadingStatus: false
})

export const getters = {
  getSettings: state => {
    return state.settings
  },
  getLoadingStatus: state => {
    return state.loadingStatus
  }
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
          {
            content: error + ': ' + JSON.stringify(error.response.data),
            color: 'red'
          },
          { root: true }
        )
      }
    )
  },
  update({ commit, state }, payload) {
    axios.put('/user/info', payload, headers()).then(
      response => {
        commit(
          'snackbar/showMessage',
          { content: 'Successfully updated', color: 'green' },
          { root: true }
        )
        $nuxt._router.push('/login')
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
  updatePreferences({ commit, state }, payload) {
    axios
      .put('/user/preferences', payload, headers())
      .then(
        response => {
          commit(
            'snackbar/showMessage',
            { content: 'Successfully updated preferences', color: 'green' },
            { root: true }
          )
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
  setSettings: (state, payload) => {
    state.settings = payload
  },
  setFormInput: (state, payload) => {
    state.settings.user_settings[payload.field] = payload.value
  },
  setUsername: (state, payload) => {
    state.settings.username = payload
  },
  update: (state, payload) => {
    state[payload.element] = payload.value
  },
  loadingStatus(state, value) {
    state.loadingStatus = value
  }
}
