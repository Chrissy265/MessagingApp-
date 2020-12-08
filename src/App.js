import React, {Component} from "react"; 
import "./App.css";
import {connect, sendMsg} from "./api";
import Header from './Components/Header/Header'; 


class App extends Component {
  constructor(props){
    super(props); 
    connect();
  } 

constructor(props) {
  super(props);
  this.state ={
    ChatHistory: []
  }
} 

componentDidMount() {
   connect((msg) =>{
     console.log("New Message")
     this.setState(prevState => ({
       ChatHistory:[...this.state.ChatHistory,msg]
     }))
     console.log(this.state);
   });
}
  send() {
    console.log("Whats up")
    sendMsg("Whats up");
  }

  render() {
    return (
      <div className = "App">
        <Header />
        <ChatHistory ChatHistory={this.state.ChatHistory}/>
        <button onClick={this.send}>HIT</button>
      </div>
    );
  }


}

export default App;
