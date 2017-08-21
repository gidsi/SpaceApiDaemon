import { combineReducers } from 'redux';
import status from './reducers/status';
import history from './reducers/history';

export default combineReducers({
  status,
  history,
});