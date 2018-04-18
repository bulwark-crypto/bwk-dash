
const PropTypes = require('prop-types');
const QRCode = require('qrcode');
const React = require('react');

const Card = require('../ui/Card');

class Donate extends React.Component {
  static propTypes = {
    donationAddress: PropTypes.string.isRequired
  };

  constructor(props) {
    super(props);
    this.qrcode = null;
  };

  componentDidMount() {
    this.qrcode = new QRCode(
      document.getElementById('donate-qr'),
      this.props.donationAddress
    );
  };

  render() {
    return (
      <div className="donate">
        <Card title="Donate Address">
          Donate QR:<br />
          <div id="donate-qr" />
        </Card>
      </div>
    );
  };
}

module.exports = Donate;
