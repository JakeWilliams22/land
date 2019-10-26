import React from 'react';
import './App.css';
import logo from './assets/logo.svg';
import JoinEmailList from './components/JoinEmailList/JoinEmailList';
import QuestionCard from './components/Questions/QuestionCard';
import {ApolloConsumer} from "./index";
import gql from 'graphql-tag'
import AdminApp from './admin/AdminApp.js'

class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            "title": "",
        };
    }

  fetchData = (client) => {
        client.query({query:
            gql`
                {
                  landingPages {
                    title
                  }
                }`}
        ).then(response => console.log(response));
  };

  render = () => (
    <div className="App">
        <ApolloConsumer>
            {apolloClient => this.fetchData(apolloClient)}
        </ApolloConsumer>
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
