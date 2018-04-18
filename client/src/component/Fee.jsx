
const PropTypes = require('prop-types');
const React = require('react');

const Card = require('../ui/Card');

const Fee = (props) => (
  <div className="fee">
    <Card title="Network Fee Summary">
      Min: { props.min }<br />
      Mid: { props.mid }<br />
      Max: { props.max }
    </Card>
  </div>
);

Fee.propTypes = {
  max: PropTypes.number.isRequired,
  mid: PropTypes.number.isRequired,
  min: PropTypes.number.isRequired,
};

module.exports = Fee;
