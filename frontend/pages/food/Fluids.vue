<template>
  <v-card>
    <v-card-title>
      <h1>Fluids intake log</h1>
    </v-card-title>

    <v-card-title>
      <v-menu
        ref="menu"
        v-model="dateRangeMenu"
        :close-on-content-click="false"
        :return-value.sync="rangeDates"
        transition="scale-transition"
        offset-y
        min-width="290px"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-text-field
            v-model="dateRangeText"
            label="Date Range"
            prepend-icon="mdi-calendar"
            readonly
            v-bind="attrs"
            v-on="on"
          ></v-text-field>
        </template>
        <v-date-picker v-model="rangeDates" range>
          <v-spacer></v-spacer>
          <v-btn text color="primary" @click="dateRangeMenu = false"
            >Cancel</v-btn
          >
          <v-btn text color="primary" @click="filterByDates">OK</v-btn>
        </v-date-picker>
      </v-menu>
    </v-card-title>

    <v-card-title>
      <span class="mr-2">Quick add</span>
      <v-btn
        class="ma-2"
        rounded
        color="primary"
        dark
        v-on:click="quickAdd('250')"
        >+ 250 ml</v-btn
      >
      <v-btn
        class="ma-2"
        rounded
        color="primary"
        dark
        v-on:click="quickAdd('500')"
        >+ 500 ml</v-btn
      >
      <v-btn
        class="ma-2"
        rounded
        color="primary"
        dark
        v-on:click="quickAdd('1000')"
        >+ 1000 ml</v-btn
      >
      <v-btn class="ma-2" rounded color="primary" dark @click="dialog = true"
        >Custom</v-btn
      >
    </v-card-title>

    <v-data-table
      :headers="headers"
      :items="getAllFluids"
      sort-by="date"
      sort-desc
      class="elevation-1"
    >
      <template v-slot:item.date="{ item }">
        <span>{{ new Date(item.date).toLocaleString() }}</span>
      </template>
      <template v-slot:top>
        <v-toolbar flat color="white">
          <v-spacer></v-spacer>
          <v-dialog v-model="dialog" max-width="500px">
            <v-card>
              <v-card-title>
                <span class="headline">{{ formTitle }}</span>
              </v-card-title>

              <v-card-text>
                <v-container>
                  <v-row>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field
                        v-model="editedItem.type"
                        label="Liquid"
                      ></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-datetime-picker
                        label="Select Datetime"
                        v-model="editedItem.date"
                      >
                        <template v-slot:dateIcon>
                          <v-icon>mdi-calendar</v-icon>
                        </template>
                        <template v-slot:timeIcon>
                          <v-icon>mdi-clock-outline</v-icon>
                        </template>
                      </v-datetime-picker>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field
                        v-model="editedItem.value"
                        :rules="[rules.positive, rules.required]"
                        label="Amount (ml)"
                      ></v-text-field>
                    </v-col>
                  </v-row>
                </v-container>
              </v-card-text>

              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="blue darken-1" text @click="close">Cancel</v-btn>
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
  </v-card>
</template>

<script>
import { mapActions, mapGetters } from "vuex";

export default {
  data: () => ({
    rules: {
      positive: value => value > 0 || "Must be more than 0",
      required: value => !!value || "Required."
    },
    rangeDates: [
      new Date().toISOString().substr(0, 10),
      new Date().toISOString().substr(0, 10)
    ],
    dateRangeMenu: false,
    menuCalendar: false,
    dialog: false,
    headers: [
      {
        text: "water",
        align: "start",
        sortable: false,
        value: "type"
      },
      { text: "Date", value: "date" },
      { text: "Amount (ml)", value: "value" },
      { text: "Actions", value: "actions", sortable: false }
    ],
    editedIndex: -1,
    editedItem: {
      type: "water",
      date: new Date(),
      value: null
    },
    defaultItem: {
      type: "water",
      date: new Date(),
      value: null
    }
  }),

  mounted() {
    this.$store.dispatch("fluids/loadFluids");
    // this.$store.subscribe(mutation => {
    //   if (mutation.type === "initFluids") {
    //     this.$store.dispatch("statsFluids");
    //   }
    // });
  },

  computed: {
    formTitle() {
      return this.editedIndex === -1 ? "Register fluid intake" : "Edit Item";
    },
    dateRangeText() {
      return this.rangeDates.join(" ~ ");
    },
    ...mapGetters('fluids', ["getAllFluids"])
  },

  watch: {
    dialog(val) {
      val || this.close();
    }
  },

  methods: {
    ...mapActions('fluids', [
      "addFluid",
      "deleteFluid",
      "updateFluid",
      "getFluidsByRange"
    ]),

    filterByDates() {
      const payload = { from: this.rangeDates[0], to: this.rangeDates[1] };
      this.getFluidsByRange(payload);
      this.dateRangeMenu = false;
    },

    editItem(item) {
      this.editedIndex = 10;
      this.editedItem = Object.assign({}, item)
      this.editedItem.date = new Date(item.date);
      this.editedItem.value = String(item.value);
      this.dialog = true;
    },

    deleteItem(item) {
      confirm("Are you sure you want to delete this item?") &&
        this.deleteFluid(item);
    },

    quickAdd(amount) {
      const payload = {
        type: "water",
        date: new Date(),
        value: amount
      };
      this.addFluid(payload);
    },

    close() {
      this.dialog = false;
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem);
        this.editedIndex = -1;
      });
    },

    save() {
      if (this.editedIndex > -1) {
        this.updateFluid(this.editedItem);
      } else {
        this.addFluid(this.editedItem);
      }
      this.close();
    }
  }
};
</script>
