import React from 'react';
import PropTypes from 'prop-types';
import gradient from 'gradient-color';
import range from 'lodash.range';
import moment from 'moment';
import injectSheet from 'react-jss';
import classnames from 'classnames';

const gradientColor = gradient(["#D9534E","#D47551","#D09453","#CCB054","#C6C856","#A8C458","#8CC059","#72BC5A","#5CB85C"], 101);
const getPercentageOpen = day => Math.round(day.value * 100 / (day.value + day.close));

const addColorClasses = style => gradientColor.reduce(
  (acc, color, index) => ({
    ...acc,
    ['color-block-' + index]: {
      borderColor: color,
      backgroundColor: color,
      color: color,
      '&:hover': {
        color: 'white',
      },
    },
  }), style);

const initHourWeekdayHistoryGrid = () => range(7).map(
  weekday => range(24).map(
    hour => ({
      weekday,
      hour,
      value: 0,
      close: 0,
    })
  ));

const formatHistoryIntoGrid = (history) => {
  const hourWeekdayHistoryGrid = initHourWeekdayHistoryGrid();
  const openTimeArray = history.filter(historyElement => historyElement.open);
  if(openTimeArray.length > 0) {
    const hourStart = moment.unix(openTimeArray[0].from);
    const now = moment();
    while(hourStart < now) {
      const hourStartUnix = hourStart.unix();
      const element = openTimeArray.filter(openTime => openTime.from <= hourStartUnix && openTime.till >= hourStartUnix);
      if(element.length > 0) {
        hourWeekdayHistoryGrid[hourStart.weekday()][hourStart.hour()].value++;
      } else {
        hourWeekdayHistoryGrid[hourStart.weekday()][hourStart.hour()].close++;
      }
      hourStart.add(1, 'hour');
    }
  }
  return hourWeekdayHistoryGrid;
};

const style = addColorClasses({
  container: {
    display: 'flex',
    flexWrap: 'wrap',
    flexDirection: 'column',
  },
  legendHours: {
    display: 'flex',
    flexDirection: 'row',
  },
  legendDays: {
    display: 'flex',
    flexDirection: 'row',
  },
  legendColor: {
    display: 'flex',
    marginTop: '10px',
  },
  legendColorBlock: {
    width: '100%',
    height: '30px',
    float: 'left',
  },
  block: {
    width: '100%',
    height: '38px',
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
  },
});

const HistoryChart = (props) => (
  <div className={props.classes.container}>
    <div>
      <div className={props.classes.legendHours}>
        <div className={props.classes.block}>
        </div>
        {initHourWeekdayHistoryGrid()[0].map((day, index) => (
          <div
            id={'day_legend_' + index}
            className={props.classes.block}
            title={index}
          >
            {index}
          </div>
        ))}
      </div>
    </div>
    {formatHistoryIntoGrid(props.history).map((weekday, index) => (
      <div id={index} >
        <div className={props.classes.legendDays}>
          <div
            className={props.classes.block}
            title={index}
          >
            {index}
          </div>
          {weekday.map((day, index) => (
            <div
              id={day.weekday + day.hour}
              className={classnames(
                props.classes.block,
                props.classes['color-block-' + getPercentageOpen(day)],
              )}
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
          title={index}
          className={classnames(
            props.classes.legendColorBlock,
            props.classes['color-block-' + index],
          )}
        />
      ))}
    </div>
  </div>
);

HistoryChart.propTypes = {
  history: PropTypes.array,
};

HistoryChart.default = {
  history: [],
};

export default injectSheet(style)(HistoryChart);