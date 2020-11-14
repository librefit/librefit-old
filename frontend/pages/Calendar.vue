<template>
  <v-row class="fill-height">
    <v-col>
      <v-sheet height="64">
        <v-toolbar flat color="white">
          <v-btn outlined class="mr-4" color="grey darken-2" @click="setToday">
            Today
          </v-btn>
          <v-btn fab text small color="grey darken-2" @click="prev">
            <v-icon small>mdi-chevron-left</v-icon>
          </v-btn>
          <v-btn fab text small color="grey darken-2" @click="next">
            <v-icon small>mdi-chevron-right</v-icon>
          </v-btn>
          <v-toolbar-title v-if="$refs.calendar">
            {{ $refs.calendar.title }}
          </v-toolbar-title>
          <v-spacer></v-spacer>
        </v-toolbar>
      </v-sheet>
      <v-sheet height="600">
        <v-calendar
          ref="calendar"
          v-model="focus"
          color="primary"
          type="month"
          :events="events"
          :event-color="getEventColor"
          @click:more="viewDay"
          @click:date="viewDay"
          @change="updateRange"
        ></v-calendar>
      </v-sheet>
    </v-col>
  </v-row>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'

export default {
  data: () => ({
    focus: '',
    events: []
  }),

  mounted() {
    this.$refs.calendar.checkChange()
    this.$store.dispatch('measures/loadMeasures')
    this.$store.dispatch('fluids/statsFluids')
  },

  created() {
    this.unsubscribe = this.$store.subscribe((mutation, state) => {
      const events = []
      var end = new Date()
      var start = new Date(end.getFullYear(), end.getMonth(), 1)

      if (mutation.type === 'measures/initMeasures') {
        // this.$store.getters['measures/getByMonth'](start, end).forEach(data => {
        //   events.push({
        //     name: `Checked-in Weight: ${data.value}`,
        //     start: new Date(`${data.date}T00:00:00`),
        //     end: new Date(`${data.date}T23:59:59`),
        //     color: 'orange',
        //     timed: false
        //   })
        // })
      }
      
      if (mutation.type === 'fluids/statsFluids') {
       // this.$store.getters['fluids/getByMonth'](start, end).forEach(data => {
       //   events.push({
       //     name: `Total Fluids: ${data.value} ml`,
       //     start: new Date(`${data.date}T00:00:00`),
       //     end: new Date(`${data.date}T23:59:59`),
       //     color: 'cyan',
       //     timed: false
       //   })
       // })
      }

      //this.events = events
    })
  },

  // TODO: Do I need this within beforeDestroy?
  beforeDestroy() {
    this.unsubscribe()
  },

  methods: {
    ...mapGetters('measures', ['getByMonth']),

    viewDay({ date }) {
      this.focus = date
      this.type = 'day'
    },
    getEventColor(event) {
      return event.color
    },
    setToday() {
      this.focus = ''
    },
    prev() {
      this.$refs.calendar.prev()
    },
    next() {
      this.$refs.calendar.next()
    },
    updateRange({ start, end }) {
      const events = []
      
      console.log("Tengo:", start)
      console.log(end)
      
      var startDate = new Date(start.date)
      var endDate = new Date(end.date)

      this.$store.getters['fluids/getByMonth'](startDate, endDate).forEach(data => {
        events.push({
          name: `Total Fluids: ${data.value} ml`,
          start: new Date(`${data.date}T00:00:00`),
          end: new Date(`${data.date}T23:59:59`),
          color: 'cyan',
          timed: false
        })
      })

      this.$store.getters['measures/getByMonth'](startDate, endDate).forEach(data => {
        events.push({
          name: `Checked-in Weight: ${data.value}`,
          start: new Date(`${data.date}T00:00:00`),
          end: new Date(`${data.date}T23:59:59`),
          color: 'orange',
          timed: false
        })
      })

      this.events = events
    }
  }
}
</script>
