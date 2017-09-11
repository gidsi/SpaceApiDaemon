import React from 'react';
import { connect } from 'react-redux';
import Component from '../components/HistoryChart';
import HistoryPropTypes from '../propTypes/history';
import ConfigPropTypes from '../propTypes/config';

const mapStateToProps = store => ({
  history: store.history,
  config: store.config,
});

const mapDispatchToProps = {};

const History = props => (
  <Component
    history={props.history.items.filter(
      historyElement => historyElement.from > props.history.filter,
    )}
    chartGradient={props.config.chartGradient}
  />
);

History.propTypes = {
  history: HistoryPropTypes.history.isRequired,
  config: ConfigPropTypes.isRequired,
};

export default connect(mapStateToProps, mapDispatchToProps)(History);
