import { createAction, handleActions } from 'redux-actions';

const CONFIG_SET = 'CONFIG_SET';

export const setConfig = createAction(CONFIG_SET, result => result);

export const actions = {
  setConfig,
};

export default handleActions({
  [CONFIG_SET]: (state, { payload }) => ({ ...state, ...payload }),
}, {
  apiUrl: '',
  displayHistory: true,
  displayHistoryChart: true,
  displayFilter: true,
  displayStatus: true,
  chartGradient: ["#D9534E","#D47551","#D09453","#CCB054","#C6C856","#A8C458","#8CC059","#72BC5A","#5CB85C"],
});