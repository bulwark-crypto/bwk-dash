
const PropTypes = require('prop-types');
const React = require('react');

const Card = require('../ui/Card');

const Status = (props) => (
  <Card
    icon="power"
    items={[
      <div><strong>Status:</strong> { props.status }</div>,
      <div><strong>Staking:</strong> { props.stakingStatus }</div>
    ]}
    title="Node Status" />
);

Status.propTypes = {
  stakingStatus: PropTypes.string.isRequired,
  status: PropTypes.string.isRequired
};

module.exports = Status;
