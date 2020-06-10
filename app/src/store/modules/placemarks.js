/* eslint no-shadow: ["error", { "allow": ["state"] }] */

import Vue from 'vue';

const state = {
  list: [0, 1, 2, 3, 4, 5],
  selected: [false, false, false, false, false, false],
  lastSelected: 0,
};

const getters = {
  placemarkIsSelected: (state) => (index) => state.selected[index] === true,
  hasPlacemarksSelection: (state) => state.selected.find((isSelected) => isSelected === true),
};

const actions = {
  /* modeSelectMultiOn */
  setPlacemarkSelectStatus({ commit }, { index, status }) {
    commit('TOGGLE_SELECT', { index, status });
    commit('SET_LAST_INDEX_SELECTED', index);
  },
  /* modeSelectRangeOn */
  selectPlacemarksRange({ commit }, target) {
    let tmpIndex = state.lastSelected;
    let tmpTarget = target;
    if (tmpTarget < tmpIndex) {
      [tmpTarget, tmpIndex] = [tmpIndex, tmpTarget];
    }
    let index;
    for (index = tmpIndex; index <= tmpTarget; index += 1) {
      commit('TOGGLE_SELECT', { index, status: true });
    }
    commit('SET_LAST_INDEX_SELECTED', index - 1);
  },
  unselectAllPlacemarks({ commit }, selectedIndex) {
    for (let index = 0; index < state.selected.length; index += 1) {
      commit('TOGGLE_SELECT', { index, status: false });
    }
    commit('SET_LAST_INDEX_SELECTED', selectedIndex);
  },
};

const mutations = {
  TOGGLE_SELECT(state, { index, status }) {
    Vue.set(state.selected, index, status);
  },
  SET_LAST_INDEX_SELECTED(state, index) {
    state.lastSelected = index;
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
