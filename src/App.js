import React, { Component } from "react";
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import Conversations from "./components/Conversations/Conversations";
import Contacts from "./components/Contacts/Contacts";
import Sidebar from "./components/Sidebar/Sidebar";

class App extends Component {
  constructor(){
    super();
    this.state = {
      name: 'React'
    };
  }

  render() {
    return (
      <Router>
        <div>
          <Sidebar />,
          <Switch>
            <Route exact path="/" component={Conversations} />
            <Route path="/Contacts/Contacts" component={Contacts} />
            <Route path="/Conversations/Conversations" component={Conversations} />
          </Switch>
        </div>
      </Router>
    );
  }
}


export default App;