const removeItemByIndex = (arr, index) => {
  if (index > -1) {
    arr.splice(index, 1);
  }
  return arr;
};

/* eslint "import/prefer-default-export": "off" */
export {
  removeItemByIndex,
};
