
const PropTypes = require('prop-types');
const React = require('react');

const Card = require('../ui/Card');

const MemPool = (props) => (
  <div className="mempool">
    <Card title="MemPool Information">
      Transactions: { props.transactions }<br />
      Size: { props.size }<br />
      TODO find way to get total RAM - minus mempool bytes!
    </Card>
  </div>
);

MemPool.propTypes = {
  size: PropTypes.number.isRequired,
  transactions: PropTypes.number.isRequired
};

module.exports = MemPool;
