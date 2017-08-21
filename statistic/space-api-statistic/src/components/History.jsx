import React from 'react';
import PropTypes from 'prop-types';
import moment from 'moment';

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

const History = (props) => {
  const longestTimeOpen = moment.duration(getLongestTimeOpen(props.history), 'seconds');
  const longestTimeClosed = moment.duration(getLongestTimeClosed(props.history), 'seconds');
  const openingTime = percentOpen(props.history);
  const medianTimeOpenObj = moment.duration(medianTimeOpen(props.history), 'seconds');
  const medianTimeClosedObj = moment.duration(medianTimeClosed(props.history), 'seconds');

  return (
    <div>
      <table>
        <tr>
          <td>
            Longest time open
          </td>
          <td>
            {longestTimeOpen.humanize()}
          </td>
        </tr>
        <tr>
          <td>
            Longest time closed
          </td>
          <td>
            {longestTimeClosed.humanize()}
          </td>
        </tr>
        <tr>
          <td>
            Time open %
          </td>
          <td>
            {Math.round(openingTime.timeOpenPercent)}%
          </td>
        </tr>
        <tr>
          <td>
            Time closed %
          </td>
          <td>
            {Math.round(openingTime.timeClosedPercent)}%
          </td>
        </tr>
        <tr>
          <td>
            Median Time Open
          </td>
          <td>
            {medianTimeOpenObj.humanize()}
          </td>
        </tr>
        <tr>
          <td>
            Median Time Closed
          </td>
          <td>
            {medianTimeClosedObj.humanize()}
          </td>
        </tr>
      </table>

    </div>
  );
};

History.propTypes = {
  history: PropTypes.array,
};

export default History;