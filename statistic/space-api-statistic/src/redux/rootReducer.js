import { combineReducers } from 'redux';
import status from './reducers/status';
import history from './reducers/history';
import config from './reducers/config';

export default combineReducers({
  status,
  history,
  config,
});