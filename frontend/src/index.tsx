import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { createStore } from 'redux';
import { mergeStyles } from 'office-ui-fabric-react';
import { Router, Route, IndexRedirect, hashHistory as history } from 'react-router';
import { initializeIcons } from 'office-ui-fabric-react/lib/Icons';
initializeIcons()

import App from './containers/App';
import Welcome from './containers/Welcome'

// Inject some global styles
mergeStyles({
  selectors: {
    ':global(body), :global(html), :global(#app)': {
      margin: 0,
      padding: 0,
      height: '100vh'
    }
  }
});

import reducers from './reducers'

let store = createStore(reducers);
setInterval(() => {
  store.dispatch({
    type: "add",
    num: 4,
  })
}, 1000);


ReactDOM.render(
  <Provider store={store}>
    <Router history={history}>
      <Route path="/" component={App}>
        <IndexRedirect to="welcome"/>
        <Route path="welcome" component={Welcome} />
      </Route>
    </Router>
  </Provider>
  , document.getElementById('app'));

