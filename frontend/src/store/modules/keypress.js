/* eslint no-shadow: ["error", { "allow": ["state"] }] */
const state = {
  shift: false,
  window: false,
};

const getters = {
  modeSelectRangeOn: (state) => state.shift,
  modeSelectMultiOn: (state) => state.window,
  modeSelectOn: (state) => state.shift || state.window,
};

const actions = {
  setKeypressStatus({ commit }, { type, status }) {
    commit('MANAGE_KEYPRESS', { type, status });
  },
};

const mutations = {
  MANAGE_KEYPRESS(state, { type, status }) {
    state[type] = status;
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
