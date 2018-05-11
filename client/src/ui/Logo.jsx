
const img = require('../img/whitelogo.svg');
const React = require('react');

const Logo = (props) => (
  <div className="logo responsive-img">
    <img
      alt="Bulwark Home Node Dashboard"
      src={ img }
      title="Bulwark Home Node Dashboard" />
  </div>
);

module.exports = Logo;
