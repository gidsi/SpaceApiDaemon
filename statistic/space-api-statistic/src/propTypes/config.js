import PropTypes from 'prop-types';

export default PropTypes.shape(
  {
    apiUrl: PropTypes.string.isRequired,
    displayHistory: PropTypes.bool.isRequired,
    displayHistoryChart: PropTypes.bool.isRequired,
    displayFilter: PropTypes.bool.isRequired,
    displayStatus: PropTypes.bool.isRequired,
    chartGradient: PropTypes.arrayOf(PropTypes.string),
  },
);
