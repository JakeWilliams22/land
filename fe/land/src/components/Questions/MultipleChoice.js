import React from 'react';
import './QuestionCard.css';
import Radio from '@material-ui/core/Radio';
import RadioGroup from '@material-ui/core/RadioGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';

// Props:   Question
//          Answers
//          Next Action
//          ID
class MultipleChoice extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            "value": null,
        };
    }

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
      this.props.options.forEach(answerText => options.push(
          <FormControlLabel
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
                <p className={"QuestionTitle"}>{this.props.question}</p>
                <RadioGroup
                    aria-label={this.props.question}
                    name={"MC" + this.props.ID}
                    value={this.state.value}
                    onChange={this.handleChange}
                >
                    {this.getOptions()}
                </RadioGroup>
            </div>
        )
}

export default MultipleChoice;
