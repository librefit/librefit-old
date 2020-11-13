<template>
  <div>
    <p class="text-center mb-4 text-h4">
      Weight tracking
    </p>
    
    <div class="py-4" />

    <base-material-card>
      <template v-slot:heading>Graph</template>
      <v-card-text>
        <LineChart
          :chartData="getLineChartData"
          :options="getLineChartOptions"
        />
      </v-card-text>
    </base-material-card>

    <div class="py-6" />

    <base-material-card>
      <template v-slot:heading>Weight</template>
      <v-data-table
        :headers="headers"
        :items="getAllMeasures"
        sort-by="date"
        sort-desc
        class="elevation-1"
      >
        <template v-slot:top>
          <v-toolbar flat color="white">
            <v-dialog v-model="dialog" max-width="500px">
              <template v-slot:activator="{ on, attrs }">
                <v-btn class="mb-2" v-bind="attrs" v-on="on">Register</v-btn>
              </template>
              <v-card>
                <v-card-title>
                  <span class="headline">{{ formTitle }}</span>
                </v-card-title>

                <v-card-text>
                  <v-container>
                    <v-row>
                      <v-col cols="12" sm="6" md="4">
                        <v-menu
                          v-model="menuCalendar"
                          :nudge-right="40"
                          transition="scale-transition"
                          offset-y
                          min-width="290px"
                        >
                          <template v-slot:activator="{ on, attrs }">
                            <v-text-field
                              v-model="editedItem.date"
                              label="Date"
                              prepend-icon="mdi-calendar"
                              v-bind="attrs"
                              v-on="on"
                            ></v-text-field>
                          </template>
                          <v-date-picker
                            v-model="editedItem.date"
                            @input="menuCalendar = false"
                          ></v-date-picker>
                        </v-menu>
                      </v-col>
                      <v-col cols="12" sm="6" md="4">
                        <v-text-field
                          v-model="editedItem.value"
                          label="Value"
                        ></v-text-field>
                      </v-col>
                    </v-row>
                  </v-container>
                </v-card-text>

                <v-card-actions>
                  <v-spacer></v-spacer>
                  <v-btn color="blue darken-1" text @click="close"
                    >Cancel</v-btn
                  >
                  <v-btn color="blue darken-1" text @click="save">Save</v-btn>
                </v-card-actions>
              </v-card>
            </v-dialog>
          </v-toolbar>
        </template>
        <template v-slot:item.actions="{ item }">
          <v-icon small class="mr-2" @click="editItem(item)">
            mdi-pencil
          </v-icon>
          <v-icon small @click="deleteItem(item)">
            mdi-delete
          </v-icon>
        </template>
      </v-data-table>
    </base-material-card>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'

export default {
  data: () => ({
    menuCalendar: false,
    dialog: false,
    headers: [
      { text: 'Date', value: 'date' },
      { text: 'value', value: 'value' },
      { text: 'Actions', value: 'actions', sortable: false }
    ],
    editedIndex: -1,
    editedItem: {
      date: new Date().toISOString().substr(0, 10),
      value: 0
    },
    defaultItem: {
      date: new Date().toISOString().substr(0, 10),
      value: 0
    }
  }),

  mounted() {
    this.$store.dispatch('measures/loadMeasures')
  },

  computed: {
    formTitle() {
      return this.editedIndex === -1 ? 'New weight' : 'Edit weight'
    },
    ...mapGetters('measures', [
      'getAllMeasures',
      'getLineChartData',
      'getLineChartOptions'
    ])
  },

  methods: {
    ...mapActions('measures', ['addMeasure', 'deleteMeasure', 'updateMeasure']),

    editItem(item) {
      this.editedIndex = 10
      this.editedItem = Object.assign({}, item)
      this.editedItem.value = String(this.editedItem.value)
      this.dialog = true
    },

    deleteItem(item) {
      confirm('Are you sure you want to delete this item?') &&
        this.deleteMeasure(item)
    },

    close() {
      this.dialog = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      })
    },

    save() {
      const { date, value } = this.editedItem
      const payload = { date, value, type: 'weight' }

      if (this.editedIndex > -1) {
        this.updateMeasure(this.editedItem)
      } else {
        this.addMeasure(payload)
      }
      this.close()
    }
  }
}
</script>
