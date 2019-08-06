import React from 'react';
import Input from '@material-ui/core/Input'

class JoinEmailList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {"email":""};

    }

    handleChange = (event) =>
        this.setState({'email': event.target.value});


    render = () =>
        (
            <div className="JoinEmailList">
                <Input value={this.state.email} onChange={this.handleChange}>Test</Input>
            </div>
        )
}

export default JoinEmailList;
