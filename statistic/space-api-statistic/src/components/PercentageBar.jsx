import React from 'react';
import PropTypes from 'prop-types';
import injectSheet from 'react-jss';

const style = {
  container: {
    display: 'flex',
    width: '100%',
  },
  bar: {
    fontSize: 'smaller',
    color: '#ffffff',
    textAlign: 'center',
    boxSizing: 'border-box',
    overflow: 'hidden',
    whiteSpace: 'nowrap',
    transition: 'all 300ms',
  },
  good: {
    extend: 'bar',
    borderRadius: '3px 0 0 3px',
  },
  bad: {
    extend: 'bar',
    borderRadius: '0 3px 3px 0',
  },
};


const PercentageBar = (props) => {
  const total = props.goodValue + props.badValue;

  if (total === 0) {
    return null;
  }

  const goodPercent = (props.goodValue * 100) / total;
  const badPercent = (props.badValue * 100) / total;

  return (
    <div className={props.classes.container}>
      {(goodPercent) ? <div
        style={{ width: `${goodPercent}%`, backgroundColor: props.goodColor }}
        className={props.classes.good}
      >
        {`${props.goodValue}% ${props.goodLegend}`}
      </div> : null}
      {badPercent ? <div
        style={{ width: `${badPercent}%`, backgroundColor: props.badColor }}
        className={props.classes.bad}
      >
        {`${props.badValue}% ${props.badLegend}`}
      </div> : null}
    </div>
  );
};

PercentageBar.propTypes = {
  goodValue: PropTypes.number,
  goodLegend: PropTypes.string,
  goodColor: PropTypes.string,
  badValue: PropTypes.number,
  badLegend: PropTypes.string,
  badColor: PropTypes.string,
  classes: PropTypes.object.isRequired, // eslint-disable-line react/forbid-prop-types
};

PercentageBar.defaultProps = {
  goodValue: 0,
  goodLegend: '',
  goodColor: '#5CB85C',
  badValue: 0,
  badLegend: '',
  badColor: '#D9534E',
};

export default injectSheet(style)(PercentageBar);
