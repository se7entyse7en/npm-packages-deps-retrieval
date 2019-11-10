import React from 'react';
import './App.css';
import Form from './form.js';

class App extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      response: {
        dependencies: null,
        err: null
      }
    };
  }

  render = () => {
    return (
      <div className="App container">
        <Form onResponseHandler={this.onResponseHandler}></Form>
      </div>
    );
  };

  onResponseHandler = (err, response) => {
    let s;
    if (err !== null) {
      s = {response: {dependencies: null, err: err}};
    } else {
      s = {response: {dependencies: response.getDependencies(), err: null}};
    }

    this.setState(s);
  };

}

export default App;
