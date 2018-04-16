
const PropTypes = require('prop-types');
const React = require('react');

const Card = require('../ui/Card');

const Bandwidth = (props) => (
  <div className="bandwidth">
    <Card title="Bandwidth Summary">
      Sent: { props.outgoingData }<br />
      Recv: { props.incomingData }<br />
      Total: { props.totalData }<br />
    </Card>
  </div>
);

Bandwidth.propTypes = {
  incomingData: PropTypes.number.isRequired,
  outgoingData: PropTypes.number.isRequired,
  totalData: PropTypes.number.isRequired
};

module.exports = Bandwidth;
