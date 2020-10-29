import axios from '~/plugins/axios'
import headers from '~/plugins/headers'

export const state = () => ({
  allFluids: [],
  stats: []
});

export const getters = {
  getAllFluids: state => {
    return state.allFluids;
  },
  // getStats: state => {
  //   var d = new Date
  //   let bigCities = state.fluids.filter(function (e) {
  //     return e.population > 3000000;
  //   });
  // },
  // totalLiquidsToday: state => {
  //   var d = new Date;
  //   console.log(state)
  //   console.log(d)
  //   state.fluids.forEach(function(item) {
  //     console.log(item['date'])
  //   })
  // },
  // getChartOptions: state => {
  //   return {
  //     chart: {
  //       id: "Weight Chart"
  //     },
  //     xaxis: {
  //       categories: state.Fluids.map(function(a) {
  //         return a.date;
  //       })
  //     }
  //   };
  // },
  // getChartSeries: state => {
  //   return {
  //     name: "series-1",
  //     data: state.Fluids.map(function(a) {
  //       return a.value;
  //     })
  //   };
  // }
};

export const actions = {
  loadFluids({ commit }) {
    axios.get('/fluids', headers()).then(
      response => {
        commit("initFluids", response.data);
      },
      error => {
        commit("showMessage", { content: error, color: "red" });
        commit("snackbar/showMessage", { content: error, color: "red" }, { root: true });
      }
    );
  },
  statsFluids({ state }) {
    var date = new Date;
    var today = date.toISOString().substr(0, 10)
    var totalFluidsToday = 0
    state.fluids.forEach(function(fluid) {
      if (fluid.date == today) {
        console.log("TRhis is today")
        totalFluidsToday += fluid.value
      }
    });
    console.log(totalFluidsToday)

  },
  addFluid({ commit }, payload) {
    axios.post('/fluid', payload, headers()).then(
      response => {
        commit("appendFluid", response.data);
      },
      error => {
        commit("snackbar/showMessage", { content: error, color: "red" }, { root: true });
      }
    );
  },
  getFluidsByRange({ commit }, payload) {
    axios.post('/fluids', payload, headers()).then(
      response => {
        commit("removeAllFluids");
        commit("initFluids", response.data);
      },
      error => {
        commit("snackbar/showMessage", { content: error, color: "red" }, { root: true });
      }
    );
  },
  deleteFluid: ({ commit }, payload) => {
    axios.delete('/fluid/' + payload.id, headers()).then(
      () => {
        commit("removeFluid", payload);
      },
      error => {
        commit("snackbar/showMessage", { content: error, color: "red" }, { root: true });
      }
    );
  },
  updateFluid: ({ commit }, payload) => {
    axios.put('/fluid/' + payload.id, payload, headers()).then(
      response => {
        commit("updateFluid", response.data);
      },
      error => {
        commit("snackbar/showMessage", { content: error, color: "red" }, { root: true });
      }
    );
  },
};

export const mutations = {
  initFluids: (state, payload) => {
    state.allFluids = payload;
  },
  appendFluid: (state, payload) => {
    state.allFluids.push(payload);
  },
  updateFluid: (state, payload) => {
    const index = state.allFluids.findIndex(x => x.id == payload.id);
    state.allFluids.splice(index, 1, payload);
  },
  removeFluid: (state, payload) => {
    const index = state.allFluids.indexOf(payload);
    state.allFluids.splice(index, 1);
  },
  removeAllFluids: (state) => {
    state.allFluids = [];
  }
};

