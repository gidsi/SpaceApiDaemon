import { createAction, handleActions } from 'redux-actions';
import request from 'superagent';

const STATUS_FETCHED = 'STATUS_FETCHED';

export const fetched = createAction(STATUS_FETCHED, result => result);

export const fetchStatus = () => (dispatch) => {
  request
    .get(`https://newstatus.chaospott.de/api/`)
    .set('Content-Type', 'application/json')
    .end(
      (err, res) => {
        if (!err) {
          dispatch(fetched(res.body));
        }
      }
    );
};

export const actions = {
  fetchStatus,
};

export default handleActions({
  [STATUS_FETCHED]: (state, { payload }) => payload,
}, null);