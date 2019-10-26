import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import * as serviceWorker from './serviceWorker';

import { ApolloProvider } from 'react-apollo';
import { ApolloClient } from 'apollo-client';
import { createHttpLink } from 'apollo-link-http';
import { InMemoryCache } from 'apollo-cache-inmemory';

const httpLink = createHttpLink({
    uri: 'http://localhost:3030/graphql/'
});

const client = new ApolloClient({
    link: httpLink,
    cache: new InMemoryCache()
});

const Apollo = React.createContext(<ApolloProvider client={null}/>);
const ApolloConsumer = Apollo.Consumer;

let app = (
    <Apollo.Provider value={client}>
        <App />
    </Apollo.Provider>
);


ReactDOM.render(app, document.getElementById('root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();

export {ApolloConsumer};