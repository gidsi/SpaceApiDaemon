import React from 'react';
import PropTypes from 'prop-types';
import moment from 'moment';
import injectSheet from 'react-jss';
import head from 'lodash.head';
import last from 'lodash.last';
import PercentageBar from './PercentageBar';

const getLongestTimeOpen = (history) => history.reduce(
  (highest, historyElement) =>
    (historyElement.difference && historyElement.difference > highest && historyElement.open)
      ? historyElement.difference
      : highest,
    0);

const getLongestTimeClosed = (history) => history.reduce(
  (highest, historyElement) =>
    (historyElement.difference && historyElement.difference > highest && !historyElement.open)
      ? historyElement.difference
      : highest,
  0);

const medianTimeOpen = (history) => {
  const list = history
    .filter(historyElement => historyElement.open)
    .sort((a, b) => a.difference - b.difference);

  const bar = list[Math.round(list.length / 2)];
  return bar ? bar.difference : 0;
};

const medianTimeClosed = (history) => {
  const list = history
    .filter(historyElement => !historyElement.open)
    .sort((a, b) => a.difference - b.difference);

  const bar = list[Math.round(list.length / 2)];
  return bar ? bar.difference : 0;
};

const percentOpen = (history) => {
  const sumTime = history.reduce((reduced, current) => ({
      timeOpen: current.open && current.difference ? reduced.timeOpen + current.difference : reduced.timeOpen,
      timeClosed: !current.open && current.difference ? reduced.timeClosed + current.difference : reduced.timeClosed,
    }), { timeOpen: 0, timeClosed: 0 });
  const timeComplete = sumTime.timeOpen + sumTime.timeClosed;

  return {
    ...sumTime,
    timeOpenPercent: sumTime.timeOpen * 100 / timeComplete,
    timeClosedPercent: sumTime.timeClosed * 100 / timeComplete,
  }
};

const style = {
  table: {
    marginTop: '30px',
    width: '100%',
  },
  row: {
    borderTop: '1px #ddd solid',
    '&:hover': {
      backgroundColor: '#eee',
    },
    '& td:first-child': {
      width: '40%',
    },
  },
  cell: {

  },
};

const History = (props) => {
  const longestTimeOpen = moment.duration(getLongestTimeOpen(props.history), 'seconds');
  const longestTimeClosed = moment.duration(getLongestTimeClosed(props.history), 'seconds');
  const openingTime = percentOpen(props.history);
  const medianTimeOpenObj = moment.duration(medianTimeOpen(props.history), 'seconds');
  const medianTimeClosedObj = moment.duration(medianTimeClosed(props.history), 'seconds');

  return (
    <table className={props.classes.table}>
      <tbody>
        <tr className={props.classes.row}>
          <td className={props.classes.cell}>
            longest time open
          </td>
          <td>
            {longestTimeOpen.humanize()}
          </td>
        </tr>
        <tr className={props.classes.row}>
          <td>
            longest time closed
          </td>
          <td>
            {longestTimeClosed.humanize()}
          </td>
        </tr>
        <tr className={props.classes.row}>
          <td>
            open / closed ratio
          </td>
          <td>
            <PercentageBar
              goodValue={Math.round(openingTime.timeOpenPercent)}
              goodLegend={'open'}
              goodColor={head(props.chartGradient)}
              badValue={Math.round(openingTime.timeClosedPercent)}
              badLegend={'closed'}
              badColor={last(props.chartGradient)}
            />
          </td>
        </tr>
        <tr className={props.classes.row}>
          <td>
            median time open
          </td>
          <td>
            {medianTimeOpenObj.humanize()}
          </td>
        </tr>
        <tr className={props.classes.row}>
          <td>
            median time closed
          </td>
          <td>
            {medianTimeClosedObj.humanize()}
          </td>
        </tr>
      </tbody>
    </table>
  );
};

History.propTypes = {
  history: PropTypes.array,
  chartGradient: PropTypes.array.isRequired,
};

export default injectSheet(style)(History);