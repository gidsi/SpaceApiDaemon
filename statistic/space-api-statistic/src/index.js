import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import jss from 'jss';
import preset from 'jss-preset-default';
import { actions as configActions } from './redux/reducers/config';

import store from './redux/store';

import App from './App';

jss.setup(preset());

export const SpaceApiStatistic = (element, config) => {
  store.dispatch(configActions.setConfig(config));
  ReactDOM.render(
    <Provider store={store}>
      <div>
        <App />
      </div>
    </Provider>,
    element,
  );
};

export default SpaceApiStatistic;

SpaceApiStatistic(document.getElementById('root'), {
  apiUrl: 'https://newstatus.chaospott.de/api',
  displayFilter: true,
  displayStatus: true,
});
