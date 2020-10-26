const uploadFromMyMaps = (conf, myMapsURL) => fetch(
  `${conf.BACK_URL}/upload/file`,
  {
    method: 'POST',
    headers: {},
    body: JSON.stringify({
      'google-my-maps-link': myMapsURL,
    }),
  },
);

export default uploadFromMyMaps;
