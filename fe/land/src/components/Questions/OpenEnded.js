import React from 'react';
import './QuestionCard.css';
import TextField from '@material-ui/core/TextField';
import FormHelperText from "@material-ui/core/FormHelperText";

let equal = require('deep-equal');

// Props:   Json
//          Question
//          ID
class OpenEnded extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            "value": "",
            "question": this.props.question,
        };
        if (this.props.json) {
            this.state = this.stateFromJson(this.props.json);
        }
    }

    componentDidUpdate(prevProps, prevState) {
        if(!equal(this.props.json, prevProps.json)) {
            this.setState(this.stateFromJson(this.props.json));
        }
    }

    stateFromJson = (json) => {
        return {
            ...this.state,
            ...{
                "question": json.questionPrompt,
                "ID": this.props.ID,
            }
        };
    };

    submitAndNext = () => {
        this.props.action();
        this.submitAnswer();
    };

    submitAnswer = () => {
        // TODO
    };

    handleChange = (e) => {
        this.setState({'value': e.target.value})
    };

    render = () =>
        (
            <div className="OpenEndedQuestion">
                {this.state.question !== undefined ? (
                    <form onSubmit={this.submitAndNext}>
                        <p className={"QuestionTitle"}>{this.state.question}</p>
                        <div className={"OpenEndedQuestionContent"}>
                            <TextField
                                variant="outlined"
                                className='OpenEndedInput'
                                onChange={this.handleChange}
                                value={this.state.value}/>
                            <FormHelperText id="component-helper-text">Press enter to submit</FormHelperText>
                        </div>
                    </form>)
                    : null}
            </div>
        )
}

export default OpenEnded;
