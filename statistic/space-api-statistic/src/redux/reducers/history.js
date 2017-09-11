import { createAction, handleActions } from 'redux-actions';
import request from 'superagent';
import moment from 'moment';

const HISTORY_FETCHED = 'HISTORY_FETCHED';
const SET_FILTER = 'SET_FILTER';
const SET_TIME = 'SET_TIME';

export const fetched = createAction(HISTORY_FETCHED, result => result);
export const setFilter = createAction(SET_FILTER, result => result);
export const setTime = createAction(SET_TIME, result => result);

export const fetchHistory = () => (dispatch, getStore) => {
  const store = getStore();
  request
    .get(`${store.config.apiUrl}/history/state`)
    .set('Content-Type', 'application/json')
    .end(
      (err, res) => {
        if (!err) {
          dispatch(fetched(res.body));
        }
      },
    );
};

export const actions = {
  fetchHistory,
  setFilter,
  setTime,
};

export default handleActions({
  [HISTORY_FETCHED]: (state, { payload }) => ({
    ...state,
    items: payload.items
      .sort((a, b) => a.lastchange - b.lastchange)
      .map(
        (data, index, array) => {
          const nextElement = array[index + 1];
          return {
            difference: nextElement ? nextElement.lastchange - data.lastchange : null,
            from: data.lastchange,
            till: nextElement ? nextElement.lastchange : null,
            open: data.open,
          };
        },
      ),
  }),
  [SET_FILTER]: (state, { payload }) => ({
    ...state,
    filter: payload,
  }),
  [SET_TIME]: (state, { payload }) => ({
    ...state,
    time: payload,
  }),
}, { items: [], filter: moment().subtract(1, 'year').unix(), time: moment().unix() });
