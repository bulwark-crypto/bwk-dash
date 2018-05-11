
const { breakdownHash } = require('../lib/breakdown');
const PropTypes = require('prop-types');
const React = require('react');

const Card = require('../ui/Card');

const Network = (props) => (
  <Card
    icon="network"
    items={[
      <div><strong>Block Height:</strong> { props.blocks }</div>,
      <div><strong>Hash Rate:</strong> { breakdownHash(props.networkHashPS) }</div>,
      <div><strong>Difficulty:</strong> { props.difficulty }</div>
    ]}
    title="Network Information" />
);

Network.propTypes = {
  blocks: PropTypes.number.isRequired,
  difficulty: PropTypes.number.isRequired,
  networkHashPS: PropTypes.number.isRequired
};

module.exports = Network;
