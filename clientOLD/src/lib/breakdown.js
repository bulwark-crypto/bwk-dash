
const breakdown = (n, labels) => {
  let idx = 0;
  while(n >= 1000) {
    n = n / 1000;
    idx++;
  }
  return `${ n.toFixed(2) } ${ labels[idx] }`;
};

const breakdownBytes = (n) => {
  return breakdown(n, ['B', 'kB', 'MB', 'GB', 'TB']);
};

const breakdownHash = (n) => {
  return breakdown(n, ['H/s', 'kH/s', 'MH/s', 'GH/s', 'TH/s']);
};

module.exports = {
  breakdown,
  breakdownBytes,
  breakdownHash
};
