import React from 'react';
import PropTypes from 'prop-types';
import moment from 'moment';

const Status = (props) => {
  if(props.status === null) {
    return null;
  }

  const lastChange = props.status.state && props.status.state.lastchange
    ? moment.unix(props.status.state.lastchange)
    : null;

  const currentOpenTime = moment.duration(props.currentTime - props.status.state.lastchange, "seconds");

  return (
    <div>
      <div>
        {props.status.space}
      </div>
      <div>
        {props.status.state && props.status.state.open ? "offen" : "geschlossen"}
      </div>
      <div>
        {"since " + currentOpenTime.humanize()}
      </div>
      <div>
        {lastChange ? lastChange.format('YYYY-MM-DD HH:mm') : null}
      </div>
      <div>
        {props.status.location.address} ({props.status.location.lon}, {props.status.location.lat})
      </div>
    </div>
  );
};

Status.propTypes = {
  status: PropTypes.object,
  currentTime: PropTypes.number,
};

Status.defaultProps = {
  status: null,
  currentTime: 0,
};

export default Status;