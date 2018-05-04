
const PropTypes = require('prop-types');
const QRCode = require('qrcode');
const React = require('react');

const Card = require('../ui/Card');

class QR extends React.Component {
  static propTypes = {
    code: PropTypes.string.isRequired
  };

  componentDidMount() {
    this.drawQRCode();
  };

  componentDidUpdate(prevProps) {
    if (prevProps.code !== this.props.code) {
      this.drawQRCode();
    }
  };

  drawQRCode = () => {
    QRCode.toCanvas(
      document.getElementById('donate-qr'),
      this.props.code,
      {},
      (err) => {
        if (err) {
          console.log(err);
        }
      }
    );
  };

  render() {
    return (
      <canvas id="donate-qr" />
    );
  };
}

module.exports = QR;
