import React from 'react';
import {render} from 'react-dom';
import 'whatwg-fetch';

class Form extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      sourceString: 'some data, hey some more dat',
      targetString: 'some datam, hey some more data that you like'
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    this.setState({
      sourceString: event.target.value,
      targetString: event.target.value
    });
  }

  handleSubmit(event) {


     fetch('/api/v1/distance',{
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      source: this.state.sourceString,
      target: this.state.targetString,
    })
  }) .then(function fetchDistancesResponse(response) {
    return response.json()
  }).then(function(json) {
     render(
      <Table tableData={json}/>,
      document.getElementById('table')
    );
  }).catch(function(ex) {
    console.log('parsing failed', ex)
  })
    event.preventDefault();
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
  <div className="form-group">
    <textarea className="form-control" 
      value={this.state.sourceString} 
             onChange={(e) => this.setState({ sourceString: e.target.value })}
      id="exampleTextarea" rows="3"></textarea>
    
        <small id="emailHelp" className="form-text text-muted">Enter a comma delimited list of source strings</small>
  </div>
          <div className="form-group">
    <textarea className="form-control" value={this.state.targetString} 
             onChange={(e) => this.setState({ targetString: e.target.value })}
      id="exampleTextarea2" rows="3"></textarea>
            
        <small id="emailHelp" className="form-text text-muted">Enter a comma delimited list of mapping target strings</small>
  </div>
  <button type="submit" className="btn btn-primary">Submit</button>
</form>
    );
  }
}

class Table extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      tableData: props.tableData.results
    };
  }

  render() {
    var rows = []
    this.state.tableData.forEach(function(result, i) {
      rows.push(<Row result={result} key={i} />);
    });
    
    return (
      <table className="table">
              <thead>
                <tr>
                  <th>Source</th>
                  <th>Mapping Target</th>
                  <th>String Distance</th>
                </tr>
              </thead>
              <tbody>
                  {rows}
              </tbody>
            </table>
    );
  }
}

class Row extends React.Component {
  render() {
    return (
      <tr>
        <td>{this.props.result.Source}</td>
        <td>{this.props.result.Target}</td>
        <td>{this.props.result.Distance}</td>
      </tr>
    );
  }
}

// ========================================

function fetchDistances(formData) {

  fetch('/api/v1/distance',{
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      source: formData.sourceString,
      target: formData.targetString,
    })
  }) .then(function fetchDistancesResponse(response) {
    console.log()
    return response.json()
  }).then(function(json) {
    console.log('parsed json', json)
  }).catch(function(ex) {
    console.log('parsing failed', ex)
  })
}

render(
	 <Form />,
  		document.getElementById('root')
	);