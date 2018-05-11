

const PropTypes = require('prop-types');
const React = require('react');

const {
  Card, CardContent, CardFooter, CardStatus
} = require('carbon-components-react');

const UICard = (props) => {
  return (
    <div className="bx--col-sm-12 bx--col-md-6 bx--col-lg-4">
      <Card>
        <CardFooter>
        { !!props.active && <CardStatus /> }
        </CardFooter>
        <CardContent
          cardIcon={ props.icon }
          cardInfo={ props.items }
          cardTitle={ props.title } />
      </Card>
    </div>
  );
};

UICard.propTypes = {
  active: PropTypes.bool,
  icon: PropTypes.string,
  items: PropTypes.array,
  title: PropTypes.string.isRequired
};

module.exports = UICard;
