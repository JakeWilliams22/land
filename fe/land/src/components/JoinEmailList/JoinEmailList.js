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
        console.log(email)
        let isValid = this.validate(email);
        console.log(isValid)
        this.setState({'email': email, 'valid':isValid});
    };

    validate = (email) => {
        return (email === "") || (email.includes("@") && email.includes("."));
    };

    joinList = () => {
        console.log("Join Email List with email " + this.state.email);
    };

    render = () =>
        (
            <div className="JoinEmailList">
                <Input value={this.state.email} onChange={this.handleChange} error={!this.state.valid}>Test</Input>
                <Button onClick={this.joinList}>{this.props.joinText || "Join"}</Button>
            </div>
        )
}

export default JoinEmailList;
