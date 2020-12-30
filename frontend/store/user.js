import axios from '~/plugins/axios'
import headers from '~/plugins/headers'

export const state = () => ({
  settings: [],
  username: null,
  password: null,
  loadingStatus: false,
})

export const getters = {
  getHeight: (state) => {
    return state.settings?.user_settings?.height
  },
  getSettings: (state) => {
    return state.settings
  },
  getAvatarURL: (state) => {
    return process.env.baseUrlImg + state.settings?.user_settings?.avatar
  },
  getLoadingStatus: (state) => {
    return state.loadingStatus
  },
}

export const actions = {
  pullSettings({ commit }) {
    axios.get('/user/info', headers()).then(
      (response) => {
        commit('setSettings', response.data)
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
  update({ commit, state }, payload) {
    axios.put('/user/info', payload, headers()).then(
      (response) => {
        commit(
          'snackbar/showMessage',
          { content: 'Successfully updated', color: 'green' },
          { root: true }
        )
        $nuxt._router.push('/login')
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
  async uploadAvatar({ commit, state }, payload) {
    const formData = new FormData()
    formData.append('file', payload)

    return axios.post('/upload', formData, headers()).then(
      (response) => {
        commit('setFormInput', {
          field: 'avatar',
          value: response.data.id,
        })
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
  async updatePreferences({ dispatch, commit, state }, avatar) {
    commit('loadingStatus', true)

    if (avatar.name) {
      await dispatch('uploadAvatar', avatar)
    }

    const payload = {
      full_name: state.settings.user_settings.full_name,
      email: state.settings.user_settings.email,
      birthday: state.settings.user_settings.birthday,
      use_metric: state.settings.user_settings.use_metric,
      height: state.settings.user_settings.height,
      sex: state.settings.user_settings.sex,
      avatar: state.settings.user_settings.avatar,
    }

    axios.put('/user/preferences', payload, headers()).then(
      (response) => {
        commit('loadingStatus', false)
        commit(
          'snackbar/showMessage',
          { content: 'Successfully updated preferences', color: 'green' },
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
  },
}
