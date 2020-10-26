const getRandomLatLng = () => ([
  -90 + 180 * Math.random(),
  -180 + 360 * Math.random(),
]);

const getLocationRadiusKM = (x0, y0, radius) => {
  const random = Math.random();
  // Convert radius from meters to degrees
  const radiusInDegrees = (radius * 1000.0) / 111000.0;

  const u = random;
  const v = random;
  const w = radiusInDegrees * Math.sqrt(u);
  const t = 2 * Math.PI * v;
  const x = w * Math.cos(t);
  const y = w * Math.sin(t);

  // Adjust the x-coordinate for the shrinking of the east-west distances
  const newX = x / Math.cos((y0 * Math.PI) / 180);

  const foundLongitude = y + y0;
  const foundLatitude = newX + x0;
  return [foundLatitude, foundLongitude];
};

export {
  getRandomLatLng,
  getLocationRadiusKM,
};
