<template>
  <div>
    <v-card>
      <v-tabs
        v-model="tab"
        align-with-title
        center-active
        class="text-center justify-center"
        grow
      >
        <v-tab>Credentials</v-tab>
        <v-tab>Settings</v-tab>
      </v-tabs>
    </v-card>
    <div class="py-6" />

    <v-tabs-items v-model="tab">
      <v-tab-item>
        <v-container fluid>
          <v-form>
            <v-card>
              <template v-slot:heading>Credentials</template>
              <v-card-text>
                <v-text-field
                  v-model="username"
                  label="Username"
                ></v-text-field>

                <p>Password change</p>

                <v-text-field
                  v-model="password"
                  name="password"
                  prepend-icon="mdi-lock"
                  type="password"
                  :rules="[
                    (value) => !!value || 'Please type password.',
                    (value) =>
                      (value && value.length >= 6) || 'minimum 6 characters',
                  ]"
                  label="Password"
                ></v-text-field>

                <v-text-field
                  v-model="confirmPassword"
                  label="Confirm Password"
                  name="confirmPassword"
                  prepend-icon="mdi-lock"
                  type="password"
                  :rules="[
                    !!confirmPassword || 'type confirm password',
                    password === confirmPassword ||
                      'The password confirmation does not match.',
                  ]"
                />

                <p class="text-center">
                  <v-btn class="mt-3" @click="saveCredentialsTab">
                    Update
                  </v-btn>
                </p>
              </v-card-text>
            </v-card>
          </v-form>
        </v-container>
      </v-tab-item>
      <v-tab-item>
        <v-container fluid>
          <v-form>
            <v-overlay v-if="getLoadingStatus" class="text-center">
              <v-progress-circular
                :size="70"
                indeterminate
              ></v-progress-circular>
            </v-overlay>

            <base-material-card>
              <template v-slot:heading>User Profile</template>
              <v-card-text>
                <v-file-input
                  accept="image/png, image/jpeg, image/bmp"
                  v-model="avatar"
                  placeholder="Pick an avatar"
                  prepend-icon="mdi-camera"
                  label="Avatar"
                ></v-file-input>

                <v-text-field
                  v-model="full_name"
                  label="Full Name"
                ></v-text-field>

                <v-text-field v-model="email" label="Email"></v-text-field>

                <v-dialog
                  ref="dialog"
                  v-model="dobModal"
                  :return-value.sync="dob"
                  persistent
                  width="290px"
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-text-field
                      v-model="dob"
                      label="Date of birth"
                      readonly
                      v-bind="attrs"
                      v-on="on"
                    ></v-text-field>
                  </template>
                  <v-date-picker v-model="dob" scrollable>
                    <v-spacer></v-spacer>
                    <v-btn text color="primary" @click="dobModal = false">
                      Cancel
                    </v-btn>
                    <v-btn text color="primary" @click="$refs.dialog.save(dob)">
                      OK
                    </v-btn>
                  </v-date-picker>
                </v-dialog>

                <p>Preferred unit</p>
                <v-radio-group v-model="useMetrics">
                  <v-radio label="Metrics" v-bind:value="true"></v-radio>
                  <v-radio label="U.S Standard" v-bind:value="false"></v-radio>
                </v-radio-group>

                <p>Sex</p>
                <v-radio-group v-model="sex">
                  <v-radio label="Male" value="M"></v-radio>
                  <v-radio label="Female" value="F"></v-radio>
                </v-radio-group>

                <v-slider
                  v-model="height"
                  hint="(centimeters)"
                  label="Height"
                  persistent-hint
                  max="210"
                  min="40"
                  thumb-label="always"
                  ticks
                ></v-slider>

                <p class="text-center">
                  <v-btn class="mt-3" @click="save"> Update </v-btn>
                </p>
              </v-card-text>
            </base-material-card>
          </v-form>
        </v-container>
      </v-tab-item>
    </v-tabs-items>
  </div>
</template>

<script>
import { validationMixin } from 'vuelidate'
import { required, minLength, maxLength, email } from 'vuelidate/lib/validators'
import { mapActions, mapGetters, mapState } from 'vuex'

export default {
  data: () => ({
    model: null,
    dobModal: false,
    tab: null,
    show: false,
    avatar: [],
    // Username & Password
    currentPassword: '',
    password: '',
    confirmPassword: '',
  }),

  mounted() {
    this.$store.dispatch('user/pullSettings')
  },

  computed: {
    ...mapGetters('user', ['getLoadingStatus']),

    username: {
      get() {
        if (this.$store.getters['user/getSettings']) {
          return this.$store.getters['user/getSettings'].username
        }
      },
      set(val) {
        this.$store.commit('user/update', {
          value: val,
          element: 'username',
        })
      },
    },
    full_name: {
      get() {
        if (this.$store.getters['user/getSettings'].user_settings) {
          return this.$store.getters['user/getSettings'].user_settings.full_name
        }
      },
      set(value) {
        this.$store.commit('user/setFormInput', {
          field: 'full_name',
          value: value,
        })
      },
    },
    email: {
      get() {
        if (this.$store.getters['user/getSettings'].user_settings) {
          return this.$store.getters['user/getSettings'].user_settings.email
        }
      },
      set(value) {
        this.$store.commit('user/setFormInput', {
          field: 'email',
          value: value,
        })
      },
    },
    dob: {
      get() {
        if (this.$store.getters['user/getSettings'].user_settings) {
          return this.$store.getters['user/getSettings'].user_settings.birthday
        }
      },
      set(value) {
        this.$store.commit('user/setFormInput', {
          field: 'birthday',
          value: value,
        })
      },
    },
    useMetrics: {
      get() {
        if (this.$store.getters['user/getSettings'].user_settings) {
          return this.$store.getters['user/getSettings'].user_settings
            .use_metric
        }
      },
      set(value) {
        this.$store.commit('user/setFormInput', {
          field: 'use_metric',
          value: value,
        })
      },
    },
    sex: {
      get() {
        if (this.$store.getters['user/getSettings'].user_settings) {
          return this.$store.getters['user/getSettings'].user_settings.sex
        }
      },
      set(value) {
        this.$store.commit('user/setFormInput', {
          field: 'sex',
          value: value,
        })
      },
    },
    height: {
      get() {
        if (this.$store.getters['user/getSettings'].user_settings) {
          return this.$store.getters['user/getSettings'].user_settings.height
        }
      },
      set(value) {
        this.$store.commit('user/setFormInput', {
          field: 'height',
          value: value,
        })
      },
    },
    emailErrors() {
      const errors = []
      if (!this.$v.email.$dirty) return errors
      !this.$v.email.email && errors.push('Must be valid e-mail')
      !this.$v.email.required && errors.push('E-mail is required')
      return errors
    },
  },

  methods: {
    ...mapActions('user', ['updatePreferences', 'uploadAvatar', 'update']),

    save() {
      this.updatePreferences(this.avatar)
    },

    saveCredentialsTab() {
      this.update({
        username: this.username,
        password: this.password,
      })
    },
  },

  validations: {
    email: { required, email },
  },
}
</script>
