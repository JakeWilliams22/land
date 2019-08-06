import React from 'react';
import './App.css';
import logo from './logo.svg';
import JoinEmailList from './components/JoinEmailList/JoinEmailList.js';

function App() {
  return (
    <div className="App">
        <header className="App-header">
            <img src={logo} alt="logo" width="300px"/>
            <h1>Hummingbird</h1>
            <h3>Sing like your favorite artists</h3>
        </header>
        <p>You found us while Hummingbird is still in the Lab!</p>
        <p>We'll let you know as soon as it's ready!</p>
        <JoinEmailList></JoinEmailList>
    </div>
  );
}

export default App;
