
const PropTypes = require('prop-types');
const QRCode = require('qrcode');
const React = require('react');

class QR extends React.Component {
  static propTypes = {
    code: PropTypes.string.isRequired
  };

  componentDidMount() {
    if (!!this.props.code) {
      this.drawQRCode();
    }
  };

  componentDidUpdate(prevProps) {
    if (!!this.props.code && prevProps.code !== this.props.code) {
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
    if (!this.props.code) {
      return null;
    }

    return (
      <canvas id="donate-qr" />
    );
  };
}

module.exports = QR;
