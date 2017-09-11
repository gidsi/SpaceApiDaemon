import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import Component from '../components/History';

const mapStateToProps = (store) => ({
  history: store.history,
  config: store.config,
});

const mapDispatchToProps = {};

const History = (props) => {
  return (
    <Component
      history={props.history.items.filter(
        historyElement => historyElement.from > props.history.filter
      )}
      chartGradient={props.config.chartGradient}
    />
  );
};

History.propTypes = {
  history: PropTypes.object.isRequired,
  config: PropTypes.object.isRequired,
};

export default connect(mapStateToProps, mapDispatchToProps)(History);