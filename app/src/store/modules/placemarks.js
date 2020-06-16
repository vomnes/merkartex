/* eslint no-shadow: ["error", { "allow": ["state"] }] */

import Vue from 'vue';

const state = {
  list: [],
  selected: [],
  lastSelected: 0,
};

const getters = {
  placemarkIsSelected: (state) => (index) => state.selected[index] === true,
  hasPlacemarksSelection: (state) => state.selected.find((isSelected) => isSelected === true),
  getPlacemarks: (state) => state.list,
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
  setPlacemarks({ commit }, placemarks) {
    commit('SET_PLACEMARKS', placemarks);
  },
};

const mutations = {
  TOGGLE_SELECT(state, { index, status }) {
    Vue.set(state.selected, index, status);
  },
  SET_LAST_INDEX_SELECTED(state, index) {
    state.lastSelected = index;
  },
  SET_PLACEMARKS(state, data) {
    state.list = data;
    state.selected = [];
    for (let i = 0; i < data.length; i += 1) {
      state.selected.push(false);
    }
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
