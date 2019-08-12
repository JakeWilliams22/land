import React from 'react';
import './QuestionCard.css';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button'
import Input from "@material-ui/core/Input";
import FormHelperText from "@material-ui/core/FormHelperText";
import FormControl from "@material-ui/core/FormControl";

// Props:   Question
//          Answers
//          Next Action
//          ID
class OpenEnded extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            "value": "",
        };
    }

    handleChange = (e) => {
        this.setState({'value': e.target.value})
    };

    submitAndNext = () => {
        this.props.action();
        this.submitAnswer();
    };

    submitAnswer = () => {
        // TODO
    };

    render = () =>
        (
            <div className="OpenEndedQuestion">
                <form onSubmit={this.submitAndNext}>
                    <p className={"QuestionTitle"}>{this.props.question}</p>
                    <div className={"OpenEndedQuestionContent"}>
                        <TextField
                            variant="outlined"
                            className='OpenEndedInput'
                            onChange={this.handleChange}
                            value={this.state.value}/>
                        <FormHelperText id="component-helper-text">Press enter to submit</FormHelperText>
                    </div>
                </form>
            </div>
        )
}

export default OpenEnded;
