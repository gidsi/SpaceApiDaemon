import PropTypes from 'prop-types';

const historyElement = PropTypes.shape({
  difference: PropTypes.number,
  from: PropTypes.number,
  till: PropTypes.number,
  open: PropTypes.bool,
});

const history = PropTypes.shape({
  items: PropTypes.arrayOf(historyElement).isRequired,
  filter: PropTypes.number.isRequired,
  time: PropTypes.number.isRequired,
});

export default {
  historyElement,
  history,
};
