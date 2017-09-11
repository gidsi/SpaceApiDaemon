import PropTypes from 'prop-types';

export default PropTypes.shape(
  {
    space: PropTypes.string.isRequired,
    logo: PropTypes.string.isRequired,
    url: PropTypes.string.isRequired,
    location: PropTypes.shape({
      address: PropTypes.string.isRequired,
      lat: PropTypes.number.isRequired,
      lon: PropTypes.number.isRequired,
    }),
    state: {
      open: PropTypes.bool.isRequired,
      lastchange: PropTypes.number,
    },
  },
);
