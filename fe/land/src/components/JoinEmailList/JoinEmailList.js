import React from 'react';
import './JoinEmailList.css';
import Input from '@material-ui/core/Input'
import Button from '@material-ui/core/Button'

class JoinEmailList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {"email":"","valid":true};
    };

    handleChange = (event) => {
        let email = event.target.value;
        let isValid = this.validate(email);
        this.setState({'email': email, 'valid':isValid});
    };

    validate = (email) => {
        return (email === "") || (email.includes("@") && email.includes("."));
    };

    joinList = () => {
        console.log("Join Email List with email " + this.state.email);
    };

    showThanks = () => {
      this.setState({"email":this.props.thanks || "Thanks!"});
    };

    submit = (event) => {
        event.preventDefault();
        this.joinList(event);
        this.showThanks();
    };

    render = () =>
        (
            <div>
                <p>{this.props.joinPrompt}</p>
                <div className="JoinEmailList">
                    <form autoComplete="on" onSubmit={this.submit}>
                        <Input
                            name={"email"}
                            value={this.state.email}
                            onChange={this.handleChange}
                            error={!this.state.valid}
                            placeholder={this.props.placeHolder || "Enter Your Email"}
                            autoComplete={"JEL email"}
                            type="email"/>
                        <Button type="submit">{this.props.joinText || "Join"}</Button>
                    </form>
                </div>
            </div>
        )
}

export default JoinEmailList;
