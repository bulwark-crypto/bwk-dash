
const { breakdownBytes } = require('../lib/breakdown');
const PropTypes = require('prop-types');
const React = require('react');

const Card = require('../ui/Card');

const MemPool = (props) => (
  <div className="mempool">
    <Card title="MemPool Information">
      Transactions: { props.transactions }<br />
      Size: { breakdownBytes(props.size) }<br />
      RAM Size: { breakdownBytes(props.max) }
    </Card>
  </div>
);

MemPool.propTypes = {
  size: PropTypes.number.isRequired,
  transactions: PropTypes.number.isRequired
};

module.exports = MemPool;
