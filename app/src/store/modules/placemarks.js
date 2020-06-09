/* eslint no-shadow: ["error", { "allow": ["state"] }] */

import Vue from 'vue';

const state = {
  list: [0, 1, 2, 3, 4, 5],
  selected: [false, false, false, false, false, false],
  lastSelected: 0,
};

const getters = {
  placemarkIsSelected: (state) => (index) => state.selected[index] === true,
};

const actions = {
  setPlacemarkSelectStatus({ commit }, { index, status }) {
    commit('TOGGLE_SELECT', { index, status });
    commit('SET_LAST_INDEX_SELECTED', index);
  },
  selectPlacemarksRange({ commit }, { start, end }) {
    let index = 0;
    for (index = start; index <= end || index < state.selected.length; index += 1) {
      commit('TOGGLE_SELECT', { index, status: true });
    }
    commit('SET_LAST_INDEX_SELECTED', index);
  },
  unselectAllPlacemarks({ commit }) {
    for (let index = 0; index < state.selected.length; index += 1) {
      commit('TOGGLE_SELECT', { index, status: false });
    }
  },
};

const mutations = {
  TOGGLE_SELECT(state, { index, status }) {
    Vue.set(state.selected, index, status);
  },
  SET_LAST_INDEX_SELECTED(state, index) {
    Vue.set(state.lastSelected, index);
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
