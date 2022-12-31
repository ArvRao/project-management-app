import React, {Component} from 'react';
import logo from './logo.svg';
import './App.css';

class App extends Component {
  state = {
    name: {firstName: "Arvind", lastName: "Rao"}
  }
  render(): React.ReactNode {
    return (
      <div className="App">
        Hello {this.state.name.firstName} {this.state.name.lastName}
        <button onClick={() => {
          this.setState({name: "SL"})
          console.log(this.state)
        }}>Click here</button>
      </div>
    );   
  }
}

export default App;
