import React from "react";
import {EditorInput} from "../EditorBase/Base";
import {Input} from "@material-ui/core";

class MCEditor extends React.Component {
    render = () => (
        <div>
            <EditorInput title={"Question Title"}/>
            <MCOption />
        </div>
    )
}

class MCOption extends React.Component {
    render = () => (
        <div>
            <Input /> Add / Remove
        </div>
    )
}

export default MCEditor;