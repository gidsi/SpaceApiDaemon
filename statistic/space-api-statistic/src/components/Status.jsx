import React from 'react';
import PropTypes from 'prop-types';
import moment from 'moment';
import injectSheet from 'react-jss';
import StatusPropTypes from '../propTypes/status';

const addLeadingZero = number => (number < 10 ? `0${number}` : number);

const toHHMMSS = (secondsCountString) => {
  const secondsCount = parseInt(secondsCountString, 10);
  const hours = Math.floor(secondsCount / 3600);
  const minutes = Math.floor((secondsCount - (hours * 3600)) / 60);
  const seconds = secondsCount - (hours * 3600) - (minutes * 60);

  return `${addLeadingZero(hours)}:${addLeadingZero(minutes)}:${addLeadingZero(seconds)}`;
};

const style = {
  table: {
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

const Status = (props) => {
  if (props.status === null) {
    return null;
  }

  const lastChange = props.status.state && props.status.state.lastchange
    ? moment.unix(props.status.state.lastchange)
    : null;

  const currentOpenTime = moment.duration(props.currentTime - props.status.state.lastchange, 'seconds');

  return (
    <table className={props.classes.table}>
      <tbody>
        <tr className={props.classes.row}>
          <td className={props.classes.cell}>
            space status
          </td>
          <td>
            {props.status.state && props.status.state.open ? 'open' : 'closed'}
          </td>
        </tr>
        {props.status.state.lastchange && <tr className={props.classes.row}>
          <td className={props.classes.cell}>
            current status since
          </td>
          <td>
            {lastChange ? lastChange.format('YYYY-MM-DD HH:mm:ss') : null}
          </td>
        </tr>}
        {props.status.state.lastchange && <tr className={props.classes.row}>
          <td className={props.classes.cell}>
            current status time
          </td>
          <td>
            {`${toHHMMSS(currentOpenTime.asSeconds())}h`}
          </td>
        </tr>}
      </tbody>
    </table>
  );
};

Status.propTypes = {
  status: StatusPropTypes,
  classes: PropTypes.object.isRequired, // eslint-disable-line react/forbid-prop-types
  currentTime: PropTypes.number,
};

Status.defaultProps = {
  status: null,
  currentTime: 0,
};

export default injectSheet(style)(Status);
