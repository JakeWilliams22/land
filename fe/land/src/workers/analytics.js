import ReactGA from 'react-ga';

class Analytics {
    // Name should be unique across all landing pages.
    // Use landing page ID.
    constructor(id, name) {
        this.gaID = id;
        this.gaName = "asdf";
        ReactGA.initialize([{
            trackingId: 'UA-156318179-1',
            gaOptions: {
                name: this.gaName
            }
        }]);
    }

    pageView() {
        ReactGA.pageview(this.gaName, [this.gaName]);
    }
}

export default Analytics;