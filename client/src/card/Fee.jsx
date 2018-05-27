
const PropTypes = require('prop-types');
const React = require('react');

const Card = require('../ui/Card');

const Fee = (props) => (
  <Card
    icon="dollars"
    items={[
      <div><strong>Min:</strong> { props.min < 0 ? 'N/A' : props.min }</div>,
      <div><strong>Mid:</strong> { props.mid < 0 ? 'N/A' : props.mid }</div>,
      <div><strong>Max:</strong> { props.max < 0 ? 'N/A' : props.max }</div>
    ]}
    title="Network Fee Summary" />
);

Fee.propTypes = {
  max: PropTypes.number.isRequired,
  mid: PropTypes.number.isRequired,
  min: PropTypes.number.isRequired,
};

module.exports = Fee;
