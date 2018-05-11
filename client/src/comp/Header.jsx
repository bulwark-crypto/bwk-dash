
const React = require('react');

const Logo = require('../ui/Logo');

const Header = () => (
  <div className="header bx--row">
    <div className="bx--col-sm-12 bx--col-md-2">
      <Logo />
    </div>
    <div className="bx--col-sm-12 bx--col-md-4 bx--type-beta">
      Home Node Dashboard
    </div>
    <div className="bx--col-sm-12 bx--col-md-6">
      <a
        className="rectangle-2"
        href="#"
        target="_blank">
        Help
      </a>
      <a
        className="rectangle-2"
        href="https://explorer.bulwarkcrypto.com"
        target="_blank">
        Block Explorer
      </a>
      <a
        className="rectangle-2"
        href="https://github.com/bulwark-crypto"
        target="_blank">
        GitHub
      </a>
    </div>
  </div>
);

module.exports = Header;
