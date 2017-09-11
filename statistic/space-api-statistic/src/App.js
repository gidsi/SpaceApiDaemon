import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import moment from 'moment';
import { actions as statusActions } from './redux/reducers/status';
import { actions as historyActions } from './redux/reducers/history';
import Status from "./container/Status";
import HistoryChart from "./container/HistoryChart";
import History from "./container/History";
import Filter from "./container/Filter";

const mapStateToProps = (store) => ({
  config: store.config,
});

const mapDispatchToProps = {
  ...statusActions,
  ...historyActions,
};

class App extends Component {
  componentWillMount() {
    this.props.fetchStatus();
    this.props.fetchHistory();
    setInterval(() => {
      this.props.setTime(moment().unix());
    }, 1000);
  }

  render() {
    return (
      <div>
        {this.props.config.displayStatus && <Status />}
        {this.props.config.displayFilter && <Filter />}
        {this.props.config.displayHistory && <History />}
        {this.props.config.displayHistoryChart && <HistoryChart />}
      </div>
    );
  }
}

App.propTypes = {
  current: PropTypes.object,
  history: PropTypes.object,
  config: PropTypes.object,
};

export default connect(mapStateToProps, mapDispatchToProps)(App);
