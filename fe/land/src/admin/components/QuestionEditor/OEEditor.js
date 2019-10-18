import React from "react";
import {EditorInput} from "../EditorBase/Base";
import {Input} from "@material-ui/core";

class OEEditor extends React.Component {
    render = () => (
        <div>
            <EditorInput title={"Prompt"}/>
        </div>
    )
}

export default OEEditor;