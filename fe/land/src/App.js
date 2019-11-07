import React from 'react';
import './App.css';
import logo from './assets/logo.svg';
import JoinEmailList from './components/JoinEmailList/JoinEmailList';
import QuestionCard from './components/Questions/QuestionCard';

import AdminApp from './admin/AdminApp.js'

class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            "title": "",
        };
    }

  fetchData = (gqlClient) => {
        gqlClient.query({query:
            `
                {
                  landingPages {
                    title
                  }
                }`}
        ).then(response => console.log(response));
      fetch('http://localhost:8080/graphql').then(r => r.text()).then(console.log)
  };

  render = () => (
    <div className="App">
        {/*<AdminApp />*/}
        <header className="App-header">
            <img src={logo} alt="logo" width="200px"/>
            <h1>{this.state.title}</h1>
            <h3>Sing like your favorite artists</h3>
        </header>
        <div className="App-content">
            <p>You found us while Hummingbird is still in the Lab!</p>
            <p>We'll let you know as soon as it's ready!</p>
            <JoinEmailList />
            <p>Help us build something better</p>
            <QuestionCard/>
        </div>
    </div>
  );
}

export default App;
