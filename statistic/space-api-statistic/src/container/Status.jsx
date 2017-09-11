import React from 'react';
import { connect } from 'react-redux';
import Component from '../components/Status';
import HistoryPropTypes from '../propTypes/history';
import StatusPropTypes from '../propTypes/status';

const mapStateToProps = store => ({
  history: store.history,
  status: store.status,
});

const mapDispatchToProps = {};

const Status = props => (
  <Component
    status={props.status}
    currentTime={props.history.time}
  />
);

Status.propTypes = {
  status: StatusPropTypes,
  history: HistoryPropTypes.history.isRequired,
};

Status.defaultProps = {
  status: null,
};

export default connect(mapStateToProps, mapDispatchToProps)(Status);
