import React from 'react';
import './QuestionCard.css';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import MultipleChoice from './MultipleChoice';
import OpenEnded from "./OpenEnded";

class QuestionCard extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            'question_idx': 0,
            'questions': ['Loading...'],
        };
    };

    componentDidMount() {
        this.getQuestions();
    }

    getQuestions = () => {
        // TODO(W) Fetch questions from server
        // TODO(W) Format them into ReactNodes
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
            // TODO(F) Consider adding a fade out on the element here. Not worth displaying.
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
