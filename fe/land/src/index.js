import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import * as serviceWorker from './serviceWorker';
import GqlClient from "./workers/gqlClient";

const gqlContext = React.createContext(null);
const gqlClient = new GqlClient();

let app = (
    <gqlContext.Provider value={gqlClient}>
        <App gqlClient={gqlClient}/>
    </gqlContext.Provider>
);


ReactDOM.render(app, document.getElementById('root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();

export {gqlContext};