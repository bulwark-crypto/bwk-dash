
const PropTypes = require('prop-types');
const React = require('react');

const Card = require('../ui/Card');

const Status = (props) => (
  <div className="status">
    <Card title="Node Status">
      Status: { props.status }<br />
      Staking: { props.stakingStatus }<br />
    </Card>
  </div>
);

Status.propTypes = {
  stakingStatus: PropTypes.string.isRequired,
  status: PropTypes.string.isRequired
};

module.exports = Status;
