import React from 'react';
import './QuestionCard.css';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';

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
        let q1 = <p onClick={this.nextQuestion}>Q1</p>;
        let q2 = <p onClick={this.nextQuestion}>Q2</p>;
        this.setState({'questions': [q1, q2]});
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
