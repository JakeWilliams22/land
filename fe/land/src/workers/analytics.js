import ReactGA from 'react-ga';

class Analytics {
    // Name should be unique across all landing pages.
    // Use landing page ID.
    constructor(id, name) {
        this.gaID = id;
        this.gaName = name.replace("-","");
        ReactGA.initialize([{
            trackingId: this.gaID,
            gaOptions: {
                name: this.gaName
            }
        }]);
    }

    pageView() {
        if (this.gaName !== undefined && this.gaID !== undefined) {
            ReactGA.pageview(this.gaName, [this.gaName]);
        }
    }
}

export default Analytics;