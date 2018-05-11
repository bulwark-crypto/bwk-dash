
const { breakdownBytes } = require('../lib/breakdown');
const PropTypes = require('prop-types');
const React = require('react');

const Card = require('../ui/Card');

const MemPool = (props) => (
  <Card
    icon="block-chain"
    items={[
      <div><strong>Transactions:</strong> { props.transactions }</div>,
      <div><strong>Size:</strong> { breakdownBytes(props.size) }</div>,
      <div><strong>RAM:</strong> { breakdownBytes(props.max) }</div>
    ]}
    title="Memool Information" />
);

MemPool.propTypes = {
  size: PropTypes.number.isRequired,
  transactions: PropTypes.number.isRequired
};

module.exports = MemPool;
