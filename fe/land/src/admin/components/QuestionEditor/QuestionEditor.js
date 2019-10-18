import React from "react";
import {EditorSection} from "../EditorBase/Base";
import MCEditor from "./MCEditor";
import OEEditor from "./OEEditor";

class EditorQuestions extends React.Component {
    render = () => (
        <EditorSection title={"Questions"}>
            <div>Header with add, navigate, delete questions</div>
            <div>Question Fields</div>
            <div><MCEditor /></div>
            <div><OEEditor /></div>
        </EditorSection>
    )
}

export default EditorQuestions;