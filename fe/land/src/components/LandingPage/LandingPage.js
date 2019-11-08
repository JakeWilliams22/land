import React from 'react';
import logo from '../../assets/logo.svg';
import JoinEmailList from '../../components/JoinEmailList/JoinEmailList';
import QuestionCard from '../../components/Questions/QuestionCard';

class LandingPage extends React.Component {
    constructor(props) {
        super(props);
        const landingPageData = LandingPageData.fromJson(this.props.landingPageJson);
        this.state = {
            "pageData": landingPageData,
        };
    }

    render = () => (
        <div className="LandingPage">
            <header className="LandingPage-header">
                <img src={logo} alt="logo" width="200px"/>
                <h1>{this.state.pageData.title}</h1>
                <h3>{this.state.pageData.subTitle}</h3>
            </header>
            <div className="App-content">
                <p>{this.state.pageData.bodyText}</p>
                {this.state.pageData.joinEmailList}
                <p>Help us build something better</p>
                {this.state.pageData.questionCard}
            </div>
        </div>
    );
}

class LandingPageData {
    constructor() {
        this.title = "Hummingbird";
        this.subTitle = "Sing like your favorite artists";
        this.bodyText = "You found us while Hummingbird is still in the Lab! " +
                        "We'll let you know as soon as it's ready!";
        this.joinEmailList = (<JoinEmailList />);
        this.questionCard = (<QuestionCard/>);
    }

    static fromJson = (rawJson) => {
        //TODO(CUR)
        return new LandingPageData();
    }
}

export {LandingPageData};
export default LandingPage;
