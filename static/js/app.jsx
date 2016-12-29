import React from 'react';
import {render} from 'react-dom';

class Form extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      sourceString: 'test',
      targetString: 'another test'
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
    // alert('A name was submitted: ' + this.state.targetString);
    test(this.state)
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

// ========================================

function test(formData) {
  // console.log(formData)
  alert('the form was submitted: ' + formData.sourceString);
}

render(
	 <Form />,
  		document.getElementById('root')
	);