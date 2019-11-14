import React from 'react';
import './App.css';
import Form from './form.js';
import GraphViz from './graph.js';

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
    let result = null;
    if (this.state.response.err === null) {
        if (this.state.response.dependencies !== null) {
          result = <GraphViz dependencies={this.state.response.dependencies}></GraphViz>;
        }
      result = <div></div>;
    } else {
      result = <div>{this.state.response.err.message}</div>;
    }

    return (
      <div className="App container">
        <Form onResponseHandler={this.onResponseHandler}></Form>
        {result}
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
