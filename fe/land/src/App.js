import React from 'react';
import './App.css';
import AdminApp from './admin/AdminApp.js'
import LandingPage from "./components/LandingPage/LandingPage";

const rootURL = "localhost:3000";

class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            "landingPageJson": "",
        };
    }

    componentDidMount() {
        this.fetchData();
    }

    fetchData = () => {
        const routeInfo = this.routeInfoFromURL();
        const pageId = routeInfo[0];
        const subPage = routeInfo[1];
        console.log(pageId);
        const queryText = `
            {
              landingPageById(id: "` + pageId  + `") {
                title
                //TODO(CUR)
              }
            }`;
        this.props.gqlClient.query({gqlQuery: queryText})
            .then(response => response.json())
            .then(text =>
                this.setState({"landingPageJson": text}));
    };

    routeInfoFromURL = () => {
      const splitURL = window.location.href.split("/");
      const rootURLIndex = splitURL.indexOf(rootURL);
      return [splitURL[rootURLIndex + 1], splitURL[rootURLIndex + 2]]
    };

    render = () => (
        <div className="App">
            {/*<AdminApp />*/}
            <LandingPage landingPageJson={this.state.landingPageJson} />
        </div>
    );
}

export default App;
