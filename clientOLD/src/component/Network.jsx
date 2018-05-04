
const { breakdownHash } = require('../lib/breakdown');
const PropTypes = require('prop-types');
const React = require('react');

const Card = require('../ui/Card');

const Network = (props) => (
  <div className="network">
    <Card title="Network Information">
      Block Height: { props.blocks }<br />
      Hash Rate: { breakdownHash(props.networkHashPS) }<br />
      Difficulty: { props.difficulty }<br />
    </Card>
  </div>
);

Network.propTypes = {
  blocks: PropTypes.number.isRequired,
  difficulty: PropTypes.number.isRequired,
  networkHashPS: PropTypes.number.isRequired
};

module.exports = Network;
