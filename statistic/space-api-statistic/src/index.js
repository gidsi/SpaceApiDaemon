import React from 'react';
import ReactDOM from 'react-dom';
import {
  BrowserRouter as Router,
  Route
} from 'react-router-dom'
import { Provider } from 'react-redux'

import store from './redux/store';
import registerServiceWorker from './registerServiceWorker';

import App from './App';

ReactDOM.render(
  <Provider store={store}>
    <Router>
      <div>
        <Route exact path="/" component={App}/>
      </div>
    </Router>
  </Provider>,
  document.getElementById('root')
);
registerServiceWorker();
