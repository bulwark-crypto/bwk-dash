
const PropTypes = require('prop-types');
const React = require('react');

const Card = require('../ui/Card');
const QR = require('../ui/QR');

class Donate extends React.Component {
  static propTypes = {
    donationAddress: PropTypes.string.isRequired
  };

  render() {
    return (
      <div className="donate">
        <Card title="Donate Address">
          Donate QR:<br />
          <QR code={ this.props.donationAddress } />
        </Card>
      </div>
    );
  };
}

module.exports = Donate;
