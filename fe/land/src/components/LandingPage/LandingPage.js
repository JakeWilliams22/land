import React from 'react';
import logo from '../../assets/logo.svg';
import JoinEmailList from '../../components/JoinEmailList/JoinEmailList';
import QuestionCard from '../../components/Questions/QuestionCard';

class LandingPage extends React.Component {
    constructor(props) {
        super(props);
    }

    render = () => (
        <div className="LandingPage">
            <header className="LandingPage-header">
                <img src={logo} alt="logo" width="200px"/>
                <h1>{this.props.landingPageData.title}</h1>
                <h3>{this.props.landingPageData.subTitle}</h3>
            </header>
            <div className="App-content">
                <p>{this.props.landingPageData.bodyText}</p>
                {this.props.landingPageData.joinEmailList}
                <p>Help us build something better</p>
                {this.props.landingPageData.questionCard}
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
        var landingPageData = new LandingPageData();
        var landingPageJson = rawJson.data.landingPage;
        landingPageData.title = landingPageJson.title;
        landingPageData.subTitle = landingPageJson.subTitle;
        landingPageData.bodyText = landingPageJson.bodyText;
        landingPageData.joinEmailList = LandingPageData.createJoinEmailListElem(
            landingPageJson.joinEmailList.joinPrompt,
            landingPageJson.joinEmailList.joinButtonText);
        landingPageData.questionCard =
            LandingPageData.createQuestionCardElem(landingPageJson.questions);
        console.log(landingPageData);
        return landingPageData;
    };

    static createJoinEmailListElem = (joinPrompt, joinButtonText) => {
        return (<JoinEmailList joinPrompt={joinPrompt} joinText={joinButtonText}/>)
    };

    static createQuestionCardElem = (questionJson) => {
        var questions = Array.from(questionJson);
        if (questions && questions.length) {
            return (<QuestionCard json={questions} />);
        }
        return null;
    }
}

export {LandingPageData};
export default LandingPage;
