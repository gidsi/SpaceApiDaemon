import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import Component from '../components/HistoryChart';

const mapStateToProps = (store) => ({
  history: store.history,
});

const mapDispatchToProps = {};

const History = (props) => {
  return (
    <Component
      history={props.history.items.filter(
        historyElement => historyElement.from > props.history.filter
      )}
    />
  );
};

History.propTypes = {
  history: PropTypes.object.isRequired,
};

export default connect(mapStateToProps, mapDispatchToProps)(History);