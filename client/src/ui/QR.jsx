
const PropTypes = require('prop-types');
const React = require('react');

const QR = (props) => (
  <div className="qr-code">{ props.code }</div>
);

QR.propTypes = {
  code: PropTypes.string.isRequired
};

module.exports = QR;
