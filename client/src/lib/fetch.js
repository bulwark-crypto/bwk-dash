
require('whatwg-fetch');

/**
 * getInfo will make a request to the api
 * in order to get the coin info.
 */
const getInfo = () => {
  const options = {
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json'
    },
    method: 'GET'
  };

  return fetch(`/api/info`, options)
    .then(res => res.json());
};

module.exports = { getInfo };
