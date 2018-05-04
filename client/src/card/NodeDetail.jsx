
const PropTypes = require('prop-types');
const React = require('react');

const Card = require('../ui/Card');

const NodeDetail = (props) => (
  <div className="node-detail">
    <Card title="Node Details">
      Network: { props.network }<br />
      Connections: { props.connections }<br />
      Blocks: { props.blocks }<br />
    </Card>
  </div>
);

NodeDetail.propTypes = {
  blocks: PropTypes.number.isRequired,
  connections: PropTypes.number.isRequired,
  network: PropTypes.string.isRequired
};

module.exports = NodeDetail;
