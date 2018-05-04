

const PropTypes = require('prop-types');
const React = require('react');

const { ControlledCard } = require('carbon-components-react');
const Icon = require('./Icon');

const Card = (props) => (
  <div className="cute-100 cute-6-phone cute-6-tablet cute-4-laptop">
    <div className="card">
      <div className="card-title">{ props.title }</div>
      <div className="card-content">
        { props.icon &&
          <div className="card-content-icon">
            <Icon name={ props.icon } />
          </div>
        }
        { props.children }
      </div>
    </div>
  </div>
);

Card.propTypes = {
  children: PropTypes.node,
  icon: PropTypes.string,
  title: PropTypes.string.isRequired
};

module.exports = Card;
