import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import Component from '../components/Status';

const mapStateToProps = (store) => ({
  history: store.history,
  current: store.status,
});

const mapDispatchToProps = {};

const Status = (props) => {
  return (
    <Component
      status={props.current}
      currentTime={props.history.time}
    />
  );
};

Status.propTypes = {
  history: PropTypes.object.isRequired,
};

export default connect(mapStateToProps, mapDispatchToProps)(Status);