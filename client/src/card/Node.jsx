
const PropTypes = require('prop-types');
const React = require('react');

const Card = require('../ui/Card');

const Node = (props) => (
  <div className="node">
    <Card title="Node Information">
      Bulwark Node<br />
      Subversion: { props.subversion }<br />
      Protocol: { props.protocol }<br />
      IP: { props.ip }<br />
    </Card>
  </div>
);

Node.propTypes = {
  ip: PropTypes.string.isRequired,
  protocol: PropTypes.number.isRequired,
  subversion: PropTypes.string.isRequired,
  version: PropTypes.number.isRequired
};

module.exports = Node;
