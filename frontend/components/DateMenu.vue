<template>
  <v-card v-bind="$attrs" class="mb-3">
    <v-card-text class="text-center">
      <v-menu
        ref="menu"
        v-model="dateMenu"
        :close-on-content-click="false"
        :return-value.sync="internalDate"
        transition="scale-transition"
        offset-y
        min-width="290px"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-btn class="ma-1" color="green" dark @click="moveDate('+1')">
            <v-icon dark center>mdi-arrow-left</v-icon>
          </v-btn>

          <v-btn class="ma-1" color="green" dark @click="moveDate('-1')">
            <v-icon dark center>mdi-arrow-right</v-icon>
          </v-btn>

          <v-btn
            color="green"
            dark
            v-bind="attrs"
            v-on="on"
            @click="dateMenu = true"
            >{{ internalDate | formatDate }}
          </v-btn>
        </template>
        <v-date-picker
          v-model="internalDate"
          @input="$refs.menu.save(internalDate)"
        >
        </v-date-picker>
      </v-menu>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  name: 'DateMenu',

  props: {
    date: {
      type: String,
    },
  },

  data() {
    return {
      dateMenu: false,
      internalDate: this.date,
    }
  },

  watch: {
    internalDate() {
      this.$emit('update-date', this.internalDate)
    },
  },

  methods: {
    moveDate(direction) {
      var d = new Date(this.date)
      d.setDate(d.getDate() - direction)
      this.internalDate = d.toISOString().substr(0, 10)
      this.$emit('move-date', d.toISOString().substr(0, 10))
    },
  },
}
</script>
