
const PropTypes = require('prop-types');
const React = require('react');

const Card = require('../ui/Card');

const Node = (props) => (
  <Card
    icon="infrastructure"
    items={[
      'Bulwark Node',
      <div><strong>Subversion:</strong> { props.subversion }</div>,
      <div><strong>Protocol:</strong> { props.protocol }</div>,
      <div><strong>IP Address:</strong> { props.ip }</div>
    ]}
    title="Node Information" />
);

Node.propTypes = {
  ip: PropTypes.string.isRequired,
  protocol: PropTypes.number.isRequired,
  subversion: PropTypes.string.isRequired,
  version: PropTypes.number.isRequired
};

module.exports = Node;
