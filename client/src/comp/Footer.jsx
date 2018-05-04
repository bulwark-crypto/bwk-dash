
const PropTypes = require('prop-types');
const React = require('react');

const Footer = (props) => (
  <div className="footer">
    <div>BWK-Dash: Bulwark Node Dashboard</div>
    <div>Copyright &copy; 2019 Bulwark Cryptocurrency</div>
    <div>Donate to: { props.donation }</div>
  </div>
);

Footer.propTypes = {
  donation: PropTypes.string.isRequired
};

module.exports = Footer;
