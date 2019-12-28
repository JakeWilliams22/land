import React from 'react';
import './QuestionCard.css';
import Radio from '@material-ui/core/Radio';
import RadioGroup from '@material-ui/core/RadioGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';

let equal = require('deep-equal');

// Props:   Question
//          options
//          Next Action
//          ID
//          Json
class MultipleChoice extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            "options": this.props.options,
            "question": this.props.question,
            "ID": this.props.ID,
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
          "options": json.answers,
          "question": json.questionPrompt,
          "ID": this.props.ID,
      };
    };

    handleChange = (e) => {
      this.setState({'value': e.target.value})
      this.submitAndNext();
    };

    submitAndNext = () => {
        this.props.action();
        this.submitAnswer();
    };

    submitAnswer = () => {
        // TODO
    };

    getOptions = () => {
      let options = [];
      if (!this.state.options) {
          return null
      }
      this.state.options.forEach((answerText, idx) => options.push(
          <FormControlLabel
              key={idx}
              control={<Radio/>}
              label={answerText}
              value={answerText}
              className={"MultipleChoiceOption"}/>
      ));
      return options;
    };

    render = () =>
        (
            <div className="MultipleChoiceQuestion">
                {this.state.question !== undefined && this.state.options !== undefined ? (
                        <div>
                            <p className={"QuestionTitle"}>{this.state.question}</p>
                            <RadioGroup
                                aria-label={this.state.question}
                                name={"MC" + this.state.ID}
                                value={this.state.value}
                                onChange={this.handleChange}
                            >
                                {this.getOptions()}
                            </RadioGroup>
                        </div>)
                    : null}
            </div>
        );
}

export default MultipleChoice;
