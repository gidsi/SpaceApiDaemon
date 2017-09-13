import React from 'react';
import PropTypes from 'prop-types';
import moment from 'moment';
import injectSheet from 'react-jss';
import HistoryPropTypes from '../propTypes/history';
import RadioButton from './RadioButton';

const getMinMaxForHistory = history => ({
  minimumAmount: history[0] ? history[0].from : 0,
  maximumAmount: history[history.length - 2] ? history[history.length - 2].from : 0,
});

const style = {
  container: {
    marginTop: '30px',
    display: 'flex',
    width: '100%',
  },
  filterLabel: {
    paddingTop: '3px',
    paddingRight: '10px',
  },
};

const Filter = (props) => {
  const minMax = getMinMaxForHistory(props.history);

  const buttonList = [
    {
      id: '1',
      display: 'all time',
      value: minMax.minimumAmount,
    },
    {
      id: '2',
      display: '1 year',
      value: moment().subtract(1, 'year').unix(),
    },
    {
      id: '3',
      display: '6 month',
      value: moment().subtract(6, 'month').unix(),
    },
    {
      id: '4',
      display: '2 weeks',
      value: moment().subtract(2, 'week').unix(),
    },
  ];

  return (
    <div className={props.classes.container}>
      <div className={props.classes.filterLabel}>
        Filter
      </div>
      {buttonList.map(button => (
        <RadioButton
          display={button.display}
          group={'foo'}
          key={button.id}
          id={`foo-${button.id}`}
          value={`${button.value}`}
          checked={moment.unix(props.filterValue).format('YYYY-MM-DD') === moment.unix(parseInt(button.value, 10)).format('YYYY-MM-DD')}
          onChange={(e) => { props.setFilter(parseInt(e.target.value, 10)); }}
        />
      ))}
    </div>
  );
};

Filter.propTypes = {
  filterValue: PropTypes.number,
  // eslint-disable-next-line
  setFilter: PropTypes.func, // line is used, linterbug
  history: PropTypes.arrayOf(HistoryPropTypes.historyElement),
  classes: PropTypes.object.isRequired, // eslint-disable-line react/forbid-prop-types
};

Filter.defaultProps = {
  filterValue: 0,
  setFilter: () => {},
  history: [],
};

export default injectSheet(style)(Filter);
