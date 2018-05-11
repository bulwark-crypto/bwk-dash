
const PropTypes = require('prop-types');
const React = require('react');

const Card = require('../ui/Card');

const NodeDetail = (props) => (
  <Card
    active={ props.status !== 'Offline' }
    icon="mobile"
    items={[
      <div><strong>Network:</strong> { props.network }</div>,
      <div><strong>Connections:</strong> { props.connections }</div>,
      <div><strong>Blocks:</strong> { props.blocks }</div>
    ]}
    title="Node Details" />
);

NodeDetail.propTypes = {
  blocks: PropTypes.number.isRequired,
  connections: PropTypes.number.isRequired,
  network: PropTypes.string.isRequired,
  status: PropTypes.string.isRequired
};

module.exports = NodeDetail;
