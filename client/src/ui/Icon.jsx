
const PropTypes = require('prop-types');
const React = require('react');

const Icon = (props) => (
  <div className="material-icons">{ props.name }</div>
);

Icon.propTypes = {
  name: PropTypes.string.isRequired
};

module.exports = Icon;
