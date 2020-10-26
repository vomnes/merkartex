const uploadFile = (conf, data) => {
  const formData = new FormData();
  formData.append('map', data.map, data.filename);
  return fetch(
    `${conf.BACK_URL}/upload/file`,
    {
      method: 'POST',
      headers: {},
      body: formData,
    },
  );
};

export default uploadFile;
