import React from 'react';
import './App.css';
import AdminApp from './admin/AdminApp.js'
import LandingPage, {LandingPageData} from "./components/LandingPage/LandingPage";
import {landingPageByIdQuery} from "./workers/gqlQueries";
import GqlClient from "./workers/gqlClient";

const rootURL = "localhost:3000";

class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            "landingPageData": new LandingPageData(),
        };
    }

    componentDidMount() {
        this.fetchData()
            .then(() => {
                this.state.landingPageData.pageview();
            });
    }

    fetchData = () => {
        const routeInfo = this.routeInfoFromURL();
        const pageId = routeInfo[0];
        const subPage = routeInfo[1];
        const queryText = landingPageByIdQuery(pageId);
        return GqlClient.query({gqlQuery: queryText})
            .then(response => response.json())
            .then(text =>
                this.setState({"landingPageData": LandingPageData.fromJson(text)}));
    };

    routeInfoFromURL = () => {
      const splitURL = window.location.href.split("/");
      const rootURLIndex = splitURL.indexOf(rootURL);
      return [splitURL[rootURLIndex + 1], splitURL[rootURLIndex + 2]]
    };

    render = () => (
        <div className="App">
            {/*<AdminApp />*/}
            <LandingPage landingPageData={this.state.landingPageData} />
        </div>
    );
}

export default App;
