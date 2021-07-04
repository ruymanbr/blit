import React from 'react';
import logo from './blit_logo.png';
import './App.css';
import axios from "axios";


class NameForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      status: "online",
      files: [],
      path: '',
      value:    '',
      totFiles: ' 0 files',
      totSize:  ' 0 Kb'
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);    
  }

  handleChange(event) {
    this.setState({value: event.target.value});
  }

  handleSubmit(event) {
    event.preventDefault();
    const data = {
        "path": this.state.value
    }
    axios.post("http://localhost:8080/api/v1/post", data,
          {
              headers:{
                  'Content-Type': 'application/json'
              }})    
    .then( response => {
      this.setState({files: response.data.data})
      this.setState({totFiles: response.data.totFiles})
      this.setState({totSize: response.data.totSize})
      this.setState({path: this.state.value})
      this.setState({status: "online"})
    })
    .catch( error => {
      if (!error.response) {
        this.setState({status: "offline"})
      } else {
        console.log(error)
        this.setState({files: []})
        this.setState({totFiles: '0 files'})
        this.setState({totSize: '0 Kb'})
        this.setState({path: this.state.value})
      }
      
    })
  }

  renderTableData() {
    return this.state.files.map((file, index) => {
      const { IsDir, LastM, FName, FSize } = file //destructuring
      return (
        <tr key={index}>
          <td className="w10">{IsDir}</td>
          <td className="w30">{LastM}</td>
          <td className="w45" dir={IsDir}>{FName}</td>
          <td className="w15">{FSize}</td>
        </tr>
      )
    })
  }
 
  getTotalFiles() {
    return this.state.totFiles
  }

  getTotalSize() {
    return this.state.totSize
  }

  componentDidMount() {
    
    
  }

  render() {
    return (
      <div className="App" onChange={this.UpdateTables}>
        <h1>Welcome to Blit Frontend demonstration</h1>
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
          Enter a <code>/full/path/</code> in your OS filesystem and click SUBMIT to list files and folders
          </p>          
          <a className="App-link" href="https://github.com/ruymanbr/blit" target="_blank" rel="noopener noreferrer">Github Repo</a>
          <br />
          <form className="App-form" onSubmit={this.handleSubmit}>
          <label>
          Path:
          </label>
          <input type="text" value={this.state.value} onChange={this.handleChange} />
          
          <input type="submit" value="Submit" />
          </form>
          <br />
          <div>
            <div className="backend-status-wrapper">Backend: { this.state.status === "online" ? <span className={this.state.status}>ONLINE</span> : <span className={this.state.status}>OFFLINE</span> }</div>
          </div>
        </header>
        <div className="App-response">          
          <table id='files'>
             <tbody>
                <tr>
                  <th></th>
                  <th></th>
                  <th className="counters">Total files: {this.getTotalFiles()}</th>
                  <th className="counters">Total Size: {this.getTotalSize()}</th>
                </tr>
                <tr className="header">
                  <th className="headers w10">Directory (y/n)</th>
                  <th className="headers w30">Last Modified</th>
                  <th className="headers w45">Name</th>
                  <th className="headers w15">Size</th>
                </tr>
                {this.renderTableData()}
             </tbody>
          </table>
        </div>
      </div>
    );
  }
}


NameForm.defaultProps = {
    action: 'http://localhost:8080/api/v1/post',
    method: 'post'
};

export default NameForm;

