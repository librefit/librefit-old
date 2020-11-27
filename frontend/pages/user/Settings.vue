<template>
  <v-container fluid>
    <v-form>
      <!-- <p>asdasd: {{ ale }} </p> -->
      <base-material-card>
        <template v-slot:heading>User Profile</template>
        <v-card-text>
          <v-text-field v-model="username" label="Username"></v-text-field>
          <v-text-field v-model="name" label="First name"></v-text-field>
          <v-text-field
            v-model="email"
            label="Email"
            @input="$v.email.$touch()"
            :error-messages="emailErrors"
            @blur="$v.email.$touch()"
          ></v-text-field>

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

          <v-select :items="unitItems" label="Preferred units" dense></v-select>
          <v-select :items="sexItems" label="Sex" dense></v-select>

          <v-row align="center">
            <p>Height:</p>
            <v-col class="d-flex" cols="4" sm="2">
              <v-select :items="mItems" filled label="m" dense></v-select>
            </v-col>

            <v-col class="d-flex" cols="4" sm="2">
              <v-select :items="mCm" filled label="cm" dense></v-select>
            </v-col>
          </v-row>

          <p class="text-center">
            <v-btn class="mt-3" type="submit"> Save </v-btn>
          </p>
        </v-card-text>
      </base-material-card>
    </v-form>
  </v-container>
</template>

<script>
import { validationMixin } from 'vuelidate'
import { required, minLength, maxLength, email } from 'vuelidate/lib/validators'
import { mapActions, mapGetters, mapState } from 'vuex'

export default {
  data: () => ({
    unitItems: ['U.S Standard', 'Metric'],
    mItems: [...Array(3).keys()],
    mCm: [...Array(100).keys()],
    sexItems: ['Male', 'Female'],
    dobModal: false,
    cards: ['Today', 'Yesterday'],
    drawer: null
  }),

  mounted() {
    this.$store.dispatch('user/pullSettings')
  },

  computed: {
    username: {
      get() {
        if (this.$store.getters['user/getSettings']) {
          return this.$store.getters['user/getSettings'].username
        }
      },
      set(value) {
        this.$store.commit('user/setUsername', value)
      }
    },
    name: {
      get() {
        if (this.$store.getters['user/getSettings'].user_settings) {
          return this.$store.getters['user/getSettings'].user_settings.full_name
        }
      },
      set(value) {
        this.$store.commit('user/setFullName', value)
      }
    },
    email: {
      get() {
        if (this.$store.getters['user/getSettings'].user_settings) {
          return this.$store.getters['user/getSettings'].user_settings.email
        }
      },
      set(value) {
        this.$store.commit('user/setEmail', value)
      }
    },
    dob: {
      get() {
        if (this.$store.getters['user/getSettings'].user_settings) {
          return this.$store.getters['user/getSettings'].user_settings.birthday
        }
      },
      set(value) {
        this.$store.commit('user/setDob', value)
      }
    },
    useMetrics: {
      get() {
        if (this.$store.getters['user/getSettings'].user_settings) {
          return this.$store.getters['user/getSettings'].user_settings
            .use_metric
        }
      },
      set(value) {
        this.$store.commit('user/setUseMetric', value)
      }
    },

    emailErrors() {
      const errors = []
      if (!this.$v.email.$dirty) return errors
      !this.$v.email.email && errors.push('Must be valid e-mail')
      !this.$v.email.required && errors.push('E-mail is required')
      return errors
    }
  },

  validations: {
    email: { required, email }
  }
}
</script>
