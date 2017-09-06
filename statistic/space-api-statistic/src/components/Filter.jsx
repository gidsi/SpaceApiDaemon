import React from 'react';
import PropTypes from 'prop-types';
import moment from 'moment';
import injectSheet from 'react-jss';

const getMinMaxForHistory = (history) => ({
  minimumAmount: history[0] ? history[0].from : 0,
  maximumAmount: history[history.length-2] ? history[history.length-2].from : 0,
});

const style = {
  container: {
    display: 'flex',
    width: '100%',
  },
  foo: {
    whiteSpace: 'nowrap',
    marginLeft: '10px',
    marginTop: '3px',
  },
  baz: {
    width: '100%',
  },
};

const Filter = (props) => {
  const minMax = getMinMaxForHistory(props.history);
  return (
    <div className={props.classes.container}>
      <input
        className={props.classes.baz}
        value={props.filterValue}
        onChange={(e) => { props.setFilter(parseInt(e.nativeEvent.target.value, 10))}}
        type="range"
        min={minMax.minimumAmount}
        max={minMax.maximumAmount}
        step="84600"
      />
      <div className={props.classes.foo}>
        {moment.unix(props.filterValue).format("YYYY-MM-DD")}
      </div>
    </div>
  );
};

Filter.propTypes = {
  filterValue: PropTypes.number,
  setFilter: PropTypes.func,
  history: PropTypes.array,
};

Filter.defaultProps = {
  filterValue: 0,
  setFilter: () => {},
  history: [],
};

export default injectSheet(style)(Filter);