
const PropTypes = require('prop-types');
const React = require('react');

const Card = require('../ui/Card');

const Donate = (props) => (
  <div className="donate">
    <Card title="Donate Address">
      Donate QR: { props.donationAddress }
    </Card>
  </div>
);

Donate.propTypes = {
  donationAddress: PropTypes.string.isRequired
};

module.exports = Donate;
