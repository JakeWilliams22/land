import React from 'react';
import './QuestionCard.css';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import MultipleChoice from './MultipleChoice';
import OpenEnded from "./OpenEnded";

let equal = require('deep-equal');

class QuestionCard extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            'question_idx': 0,
            'questions': ['Loading...'],
        };
    };

    componentDidUpdate(prevProps, prevState) {
        if(!equal(this.props.json, prevProps.json)) {
            this.getQuestionsFromJson(this.props.json);
        }
    }

    getQuestionsFromJson = (questionList) => {
        let questions = [];
        if (!questionList) {
            return;
        }
        questionList.forEach((questionJson, idx) => {
            switch (parseInt(questionJson.questionType)) {
                case 0:
                    questions.push(<MultipleChoice
                        json={questionJson} action={this.nextQuestion} ID={idx}/>);
                    break;
                case 1:
                    questions.push(<OpenEnded
                        json={questionJson} action={this.nextQuestion} ID={idx}/>);
                    break;
                default:
                    console.log("Question not supported");
                    break;
            }
        });
        this.setState({'questions': questions});
    };

    getQuestions = () => {
        let q0 = <OpenEnded
            action={this.nextQuestion}
            question={"What sucks for you"}/>
        let q1 = <MultipleChoice
            action={this.nextQuestion}
            question={"Gender"}
            options={["M","F","IDFK"]}
            ID={"0"}/>;
        let q2 = <MultipleChoice
            action={this.nextQuestion}
            question={"Problems"}
            options={["99 of em but not bitches","Bitches"]}
            ID={"1"}/>;
        this.setState({'questions': [q0, q1, q2]});
    };

    nextQuestion = () => {
        this.setState({'question_idx': this.state.question_idx + 1});
    };

    getCurrentQuestion = () => {
        if (this.state.question_idx >= this.state.questions.length) {
            return "Thanks!";
            // TODO(F) Consider adding a fade out here.
        }
        return this.state.questions[this.state.question_idx];
    };

    render = () =>
        (
            <div className="QuestionCard">
                <Card>
                    <CardContent>
                        {this.getCurrentQuestion()}
                    </CardContent>
                </Card>
            </div>
        )
}

export default QuestionCard;
