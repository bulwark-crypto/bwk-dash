
require('babel-polyfill');
const React = require('react');
const ReactDOM = require('react-dom');

const App = require('./container/App');

ReactDOM.render(
  <App />,
  document.getElementById('app-root')
);
