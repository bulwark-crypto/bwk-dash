
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
      <Card
        icon="user"
        items={[
          <QR code={ this.props.donationAddress } />
        ]}
        title="QR Code" />
    );
  };
}

module.exports = Donate;
