import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { actions as historyActions } from '../redux/reducers/history';
import Component from '../components/Filter';

const mapStateToProps = (store) => ({
  history: store.history,
});

const mapDispatchToProps = {
  ...historyActions,
};

const Filter = (props) => {
  return (
    <Component
      filterValue={props.history.filter}
      setFilter={props.setFilter}
      history={props.history.items}
    />
  );
};

Filter.propTypes = {
  history: PropTypes.object.isRequired,
  setFilter: PropTypes.func.isRequired,
};

export default connect(mapStateToProps, mapDispatchToProps)(Filter);