/* eslint no-shadow: ["error", { "allow": ["state"] }] */

import Vue from 'vue';
import { ArrayLib } from 'assets/library';

const state = {
  title: '',
  list: [],
  selected: [],
  lastSelected: 0,
  hasChanges: false,
};

const getters = {
  placemarkIsSelected: (state) => (index) => state.selected[index] === true,
  hasPlacemarksSelection: (state) => state.selected.find((isSelected) => isSelected === true),
  getPlacemarks: (state) => state.list,
  getTitle: (state) => state.title,
  getHasChanges: (state) => state.hasChanges,
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
  setTitle({ commit }, title) {
    commit('SET_TITLE', title);
  },
  removePlacemark({ commit }, index) {
    commit('REMOVE_PLACEMARK', index);
    commit('SET_HAS_CHANGES', true);
  },
  updatePlacemark({ commit }, { id, data }) {
    commit('UPDATE_PLACEMARK', { id, data });
    commit('SET_HAS_CHANGES', true);
  },
  pushPlacemark({ commit }, data) {
    const newData = data;
    if (state.list.length === 0) {
      newData.id = 0;
    } else {
      newData.id = state.list[state.list.length - 1].id + 1;
    }
    commit('PUSH_NEW_PLACEMARK', newData);
    commit('SET_HAS_CHANGES', true);
  },
  editMultiplePlacemarks({ commit }, { title, icon }) {
    state.selected.forEach((item, i) => {
      if (item === true) {
        const newData = state.list[i];
        newData.name = (title.before ? title.before : '') + newData.name + (title.after ? title.after : '');
        if (icon.style !== '') {
          newData.icon.style = icon.style;
        }
        if (icon.category !== '') {
          newData.icon.category = icon.category;
        }
        commit('UPDATE_PLACEMARK_BY_INDEX', i, newData);
      }
    });
    commit('SET_HAS_CHANGES', true);
  },
  toggleHasChanges({ commit }, value) {
    commit('SET_HAS_CHANGES', value);
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
  SET_TITLE(state, title) {
    state.title = title;
  },
  REMOVE_PLACEMARK(state, index) {
    ArrayLib.removeItemByIndex(state.list, index);
  },
  UPDATE_PLACEMARK(state, { id, data }) {
    for (let index = 0; index < state.list.length; index += 1) {
      if (state.list[index].id === id) {
        Vue.set(state.list, index, data);
        return;
      }
    }
  },
  PUSH_NEW_PLACEMARK(state, data) {
    state.list.push(data);
  },
  UPDATE_PLACEMARK_BY_INDEX(state, { index, data }) {
    Vue.set(state.list, index, data);
  },
  SET_HAS_CHANGES(state, value) {
    state.hasChanges = value;
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
