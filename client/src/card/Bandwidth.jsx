
const { breakdownBytes } = require('../lib/breakdown');
const PropTypes = require('prop-types');
const React = require('react');

const Card = require('../ui/Card');

const Bandwidth = (props) => (
  <div className="bandwidth">
    <Card title="Bandwidth Summary">
      Sent: { breakdownBytes(props.outgoingData) }<br />
      Recv: { breakdownBytes(props.incomingData) }<br />
      Total: { breakdownBytes(props.totalData) }<br />
    </Card>
  </div>
);

Bandwidth.propTypes = {
  incomingData: PropTypes.number.isRequired,
  outgoingData: PropTypes.number.isRequired,
  totalData: PropTypes.number.isRequired
};

module.exports = Bandwidth;
