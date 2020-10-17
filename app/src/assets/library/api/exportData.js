const exportData = (conf, data) => fetch(
  `${conf.BACK_URL}/export/kmz`,
  {
    method: 'POST',
    headers: {},
    body: data,
  },
);

export default exportData;
