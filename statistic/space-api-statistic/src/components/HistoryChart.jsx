import React from 'react';
import PropTypes from 'prop-types';
import range from 'lodash.range';
import moment from 'moment';
import injectSheet from 'react-jss';
import classnames from 'classnames';
import gradient from '../gradient-color';

const gradientColor = gradient(["#D9534E","#D47551","#D09453","#CCB054","#C6C856","#A8C458","#8CC059","#72BC5A","#5CB85C"], 101);
const weekdayNames = ['Mo', 'Tu', 'We', 'Th', 'Fr', 'Sa', 'Su'];
const getPercentageOpen = day => Math.round(day.open * 100 / (day.open + day.close)) | 0;

const getColorStyles = gradientColor => gradientColor.reduce(
  (acc, color, index) => ({
    ...acc,
    ['color-block-' + index]: {
      borderColor: color,
      backgroundColor: color,
      color: color,
    },
  }), style);

const initHourWeekdayHistoryGrid = () => range(7).map(
  weekday => range(24).map(
    hour => ({
      weekday,
      hour,
      open: 0,
      close: 0,
    })
  ));

const formatHistoryIntoGrid = (history) => {
  const hourWeekdayHistoryGrid = initHourWeekdayHistoryGrid();
  const openTimeArray = history.filter(historyElement => historyElement.open);
  if(openTimeArray.length > 0) {
    const hourWalker = moment.unix(openTimeArray[0].from);
    const now = moment();
    while(hourWalker < now) {
      const hourWalkerUnix = hourWalker.unix();
      const element = openTimeArray.filter(openTime => openTime.from <= hourWalkerUnix && openTime.till >= hourWalkerUnix);
      if(element.length > 0) {
        hourWeekdayHistoryGrid[hourWalker.isoWeekday() - 1][hourWalker.hour()].open++;
      } else {
        hourWeekdayHistoryGrid[hourWalker.isoWeekday() - 1][hourWalker.hour()].close++;
      }
      hourWalker.add(1, 'hours');
    }
  }
  return hourWeekdayHistoryGrid;
};

const style = {
  container: {
    marginTop: '30px',
    display: 'flex',
    flexWrap: 'wrap',
    flexDirection: 'column',
  },
  legendHours: {
    display: 'flex',
    flexDirection: 'row',
  },
  daysRow: {
    display: 'flex',
    flexDirection: 'row',
  },
  legendColor: {
    display: 'flex',
    marginTop: '10px',
    position: 'relative',
    maxWidth: 'calc(100% - 30px)',
    width: 'calc(100% - 30px)',
    marginLeft: '30px',
    borderRadius: '5px',
    overflow: 'hidden',
  },
  legendColorText: {
    display: 'flex',
    justifyContent: 'space-between',
    position: 'absolute',
    width: 'calc(100% - 10px)',
    marginLeft: '5px',
    marginTop: '3px',
    color: '#fff',
  },
  legendColorBlock: {
    height: '30px',
    float: 'left',
    width: '1%',
  },
  block: {
    width: '100%',
    height: '28px',
    marginTop: '1px',
    marginLeft: '1px',
    borderRadius: '2px',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    overflow: 'hidden',
    border: '1px transparent solid',
    color: 'black',
    transition: 'color 700ms, background-color 700ms',
    cursor: 'pointer',
    '&:hover': {
      color: 'white !important',
    },
  },
  legendDays: {
    minWidth: '30px',
    maxWidth: '30px',
  },
};

const HistoryChart = (props) => {

  const colorStyles = getColorStyles(gradient(props.chartGradient, 101));

  return (
    <div className={props.classes.container}>
      <div>
        <div className={props.classes.legendHours}>
          <div className={classnames(props.classes.block, props.classes.legendDays)}>
          </div>
          {initHourWeekdayHistoryGrid()[0].map((day, index) => (
            <div
              key={'hour_legend_' + index}
              className={props.classes.block}
              title={index}
            >
              {index}
            </div>
          ))}
        </div>
      </div>
      {formatHistoryIntoGrid(props.history).map((weekdays, index) => (
        <div key={'day_legend_' + index} >
          <div className={props.classes.daysRow}>
            <div
              className={classnames(props.classes.legendDays, props.classes.block)}
              title={index}
            >
              {weekdayNames[index]}
            </div>
            {weekdays.map((day, index) => (
              <div
                key={day.weekday + '_' + day.hour + '_' + getPercentageOpen(day)}
                style={colorStyles['color-block-' + getPercentageOpen(day)]}
                className={props.classes.block}
              >
                {getPercentageOpen(day)}
              </div>
            ))}
          </div>
        </div>
      ))}
      <div className={props.classes.legendColor}>
        {gradientColor.map((color, index) => (
          <div
            key={'percent_legend_' + index}
            title={index}
            style={colorStyles['color-block-' + index]}
            className={props.classes.legendColorBlock}
          />
        ))}
        <div className={props.classes.legendColorText}>
          <div>
            0%
          </div>
          <div>
            50%
          </div>
          <div>
            100%
          </div>
        </div>
      </div>
    </div>
  );
}

HistoryChart.propTypes = {
  history: PropTypes.array,
  chartGradient: PropTypes.array.isRequired,
};

HistoryChart.default = {
  history: [],
};

export default injectSheet(style)(HistoryChart);