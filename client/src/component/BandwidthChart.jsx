
const PropTypes = require('prop-types');
const React = require('react');

const Card = require('../ui/Card');
const Graph = require('../ui/Graph');

const BandwidthChart = (props) => {
  const data = [
    props.incomingData,
    props.outgoingData,
    //props.totalData
  ];

  return (
    <div className="bandwidth-chart">
      <Card title="Bandwidth Chart">
        <Graph data={ data } />
      </Card>
    </div>
  );
};

BandwidthChart.propTypes = {
  incomingData: PropTypes.number.isRequired,
  outgoingData: PropTypes.number.isRequired,
  totalData: PropTypes.number.isRequired
};

module.exports = BandwidthChart;
