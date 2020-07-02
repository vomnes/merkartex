/* eslint no-shadow: ["error", { "allow": ["state"] }] */

import Vue from 'vue';
import { ArrayLib } from 'assets/library';

const cloneDeep = require('clone-deep');

const state = {
  title: '',
  list: [],
  folders: [],
  selected: [],
  lastSelected: 0,
  hasChanges: false,
  length: 0,
};

const getters = {
  placemarkIsSelected: (state) => (index) => state.selected[index] === true,
  hasPlacemarksSelection: (state) => state.selected.find((isSelected) => isSelected === true),
  getPlacemarks: (state) => state.list,
  getFolders: (state) => state.folders,
  getTitle: (state) => state.title,
  getHasChanges: (state) => state.hasChanges,
};

const multiEdit = (elem, title, icon) => {
  const newData = elem;
  newData.name = (title.before ? title.before : '') + newData.name + (title.after ? title.after : '');
  if (icon.style !== '') {
    newData.icon.style = icon.style;
  }
  if (icon.category !== '') {
    newData.icon.category = icon.category;
  }
  return newData;
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
  setPlacemarksFolders({ commit }, folders) {
    commit('SET_FOLDERS', folders);
  },
  setTitle({ commit }, title) {
    commit('SET_TITLE', title);
  },
  setLength({ commit }, value) {
    commit('SET_LENGTH', value);
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
    for (let i = 0; state.folders && i < state.folders.length; i += 1) {
      for (let j = 0; state.folders[i].placemarks && j
        < state.folders[i].placemarks.length; j += 1) {
        const elem = cloneDeep(state.folders[i].placemarks[j]);
        if (elem.id < state.selected.length && state.selected[elem.id]) {
          const newData = multiEdit(elem, title, icon);
          commit('UPDATE_PLACEMARK_BY_INDEX', {
            array: state.folders[i].placemarks,
            index: j,
            data: newData,
          });
        }
      }
    }
    for (let index = 0; state.list && index < state.list.length; index += 1) {
      const elem = state.list[index];
      if (elem.id < state.selected.length && state.selected[elem.id]) {
        const newData = multiEdit(elem, title, icon);
        commit('UPDATE_PLACEMARK_BY_INDEX', { array: state.list, index, data: newData });
      }
    }
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
    for (let i = 0; i < state.length; i += 1) {
      state.selected.push(false);
    }
  },
  SET_FOLDERS(state, folders) {
    state.folders = folders;
  },
  SET_TITLE(state, title) {
    state.title = title;
  },
  REMOVE_PLACEMARK(state, index) {
    ArrayLib.removeItemByIndex(state.list, index);
  },
  UPDATE_PLACEMARK(state, { id, data }) {
    for (let i = 0; i < state.folders.length; i += 1) {
      for (let j = 0; j < state.folders[i].placemarks.length; j += 1) {
        if (state.folders[i].placemarks[j].id === id) {
          Vue.set(state.folders[i].placemarks, j, data);
          return;
        }
      }
    }
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
  UPDATE_PLACEMARK_BY_INDEX(state, { array, index, data }) {
    Vue.set(array, index, data);
  },
  SET_HAS_CHANGES(state, value) {
    state.hasChanges = value;
  },
  SET_LENGTH(state, value) {
    state.length = value;
  },
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
