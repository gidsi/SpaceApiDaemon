import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux'
import { actions as configActions } from './redux/reducers/config'
import jss from 'jss'
import preset from 'jss-preset-default'

import store from './redux/store';

import App from './App';

jss.setup(preset())

export const SpaceApiStatistic = (id, config) => {
  store.dispatch(configActions.setConfig(config));
  ReactDOM.render(
    <Provider store={store}>
      <div>
        <App />
      </div>
    </Provider>,
    document.getElementById(id)
  );
};

window.SpaceApiStatistic = SpaceApiStatistic;


SpaceApiStatistic('root', {
  apiUrl: 'https://newstatus.chaospott.de/api',
  displayFilter: true,
  displayStatus: true,
  chartGradient: ['#012D41', '#1BA5B8', '#DAECF3', '#FF404E', '#1CA5B8'],
});