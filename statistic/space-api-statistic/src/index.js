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
  /* eslint-disable react/jsx-filename-extension */
  ReactDOM.render(
    <Provider store={store}>
      <div>
        <App />
      </div>
    </Provider>,
    element,
  );
  /* eslint-enable react/jsx-filename-extension */
};

export default SpaceApiStatistic;
