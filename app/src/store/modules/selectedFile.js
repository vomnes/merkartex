/* eslint no-shadow: ["error", { "allow": ["state"] }] */
const state = {
  name: '',
  loaded: 0,
  size: 0,
  valid: false,
};

const getters = {
  stateSelectedFile: (state) => state,
  selectedFileLoadedValue: (state) => state.loaded,
};

const actions = {
  setSelectedFileValue({ commit }, { type, value }) {
    commit('SET_VALUE', { type, value });
  },
  selectedFileReset({ commit }) {
    commit('RESET');
  },
};

const mutations = {
  SET_VALUE(state, { type, value }) {
    state[type] = value;
  },
  RESET(state) {
    state.name = '';
    state.loaded = 0;
    state.size = 0;
    state.valid = false;
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
