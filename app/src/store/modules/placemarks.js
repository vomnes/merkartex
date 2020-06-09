const state = {
  list: [],
  keypress: {
    shift: false,
    window: false,
  },
};

const getters = {};

const actions = {
  setKeypressStatus({ commit }, { type, status }) {
    commit('MANAGE_KEYPRESS', { type, status });
  },
};

const mutations = {
  /* eslint no-shadow: ["error", { "allow": ["state"] }] */
  MANAGE_KEYPRESS(state, { type, status }) {
    state.keypress[type] = status;
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
