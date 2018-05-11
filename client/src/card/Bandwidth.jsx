
const { breakdownBytes } = require('../lib/breakdown');
const PropTypes = require('prop-types');
const React = require('react');

const Card = require('../ui/Card');

const Bandwidth = (props) => (
  <Card
    icon="crash"
    items={[
      <div><strong>Sent:</strong> { breakdownBytes(props.outgoingData) }</div>,
      <div><strong>Recv:</strong> { breakdownBytes(props.incomingData) }</div>,
      <div><strong>Total:</strong> { breakdownBytes(props.totalData) }</div>
    ]}
    title="Bandwidth Summary" />
);

Bandwidth.propTypes = {
  incomingData: PropTypes.number.isRequired,
  outgoingData: PropTypes.number.isRequired,
  totalData: PropTypes.number.isRequired
};

module.exports = Bandwidth;
