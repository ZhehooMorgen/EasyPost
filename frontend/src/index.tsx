import React from 'react';
import ReactDOM from 'react-dom';
import { Router, Route,hashHistory} from 'react-router';


import * as serviceWorker from './serviceWorker';

import App from './containers/App'
import Home from './containers/Home'

ReactDOM.render(
  <Router history={hashHistory}>
    <Route path="/" component={App}>
      <Route path="/home" component={Home} />
    </Route>
  </Router>, 
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
