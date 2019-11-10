import React from 'react';

const {DependenciesRequest} = require('./api_pb.js');
const {DependenciesServiceClient} = require('./api_grpc_web_pb.js');

const api = new DependenciesServiceClient('http://localhost:9090');

class Form extends React.Component {

  constructor(props) {
    super(props);
    this.state = {packageName: "", version: ""};
  }

  render = () => {
    return (
      <div className="row">
        <div className="col-8 offset-2">
          <form onSubmit={this.handleSubmit}>
            <div className="form-group">
              <label htmlFor="packageName">Package name</label>
              <input type="text" className="form-control" id="packageNameInput" value={this.state.packageName} onChange={this.handlePackageNameChange} placeholder="Enter package name"/>
            </div>
            <div className="form-group">
              <label htmlFor="version">Version</label>
              <input type="text" className="form-control" id="versionInput" value={this.state.version} onChange={this.handleVersionChange} placeholder="Enter version (default: latest)"/>
            </div>
            <button type="submit" className="btn btn-primary">Submit</button>
          </form>
        </div>
      </div>
    );
  }

  handlePackageNameChange = e => {
    this.setState({packageName: e.target.value});
  };

  handleVersionChange = e => {
    this.setState({version: e.target.value});
  };

  handleSubmit = e => {
    e.preventDefault();

    const packageName = this.state.packageName;
    const version = this.state.version;

    const request = new DependenciesRequest();
    request.setName(packageName);
    request.setVersion(version);

    api.getDependencies(request, {}, this.props.onResponseHandler);
  };

}

export default Form;
