import React from 'react';
import {render} from 'react-dom';

class App extends React.Component {
  render () {
    return <h5> Hello React!</h5>;
  }
}

render(<App/>, document.getElementById('root'));