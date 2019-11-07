import React from "react";
import InputLabel from "@material-ui/core/InputLabel";
import Input from "@material-ui/core/Input";
import {Checkbox} from "@material-ui/core";

class EditorInput extends React.Component {
    render = () =>
        (
            <div className="EditorInput">
                <InputLabel>{this.props.title}</InputLabel>
                <Input />
            </div>
        )
}

class EditorSection extends React.Component {
    constructor(props) {
        super(props);
        this.state = {"show":false};
    };

    render = () =>
        (
            <div className={"EditorSection"}>
                <Checkbox
                    checked={this.state.show}
                    onChange={() => this.setState({"show":!this.state.show})}/> {this.props.title}
                <div className={"EditorSectionContent"}>
                    {this.state.show ? this.props.children : ""}
                </div>
            </div>
        )
}

export {EditorSection, EditorInput};