import React from 'react';
import './App.css';
import logo from './assets/logo.svg';
import JoinEmailList from './components/JoinEmailList/JoinEmailList.js';
import QuestionCard from "./components/Questions/QuestionCard";

function App() {
  return (
    <div className="App">
        <header className="App-header">
            <img src={logo} alt="logo" width="200px"/>
            <h1>Hummingbird</h1>
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
