import React from 'react';
import './editor.css';
import {EditorInput, EditorSection} from "../EditorBase/Base";
import QuestionEditor from "../QuestionEditor/QuestionEditor";

class Editor extends React.Component {
    render = () =>
        (
            <div className="AdminApp">
                <EditorInput title="Title" />
                <EditorInput title="Tagline" />
                <EditorInput title="Under Construction Text" />
                <EditorJoinEmailList />
                <QuestionEditor />
            </div>
        )
}

class EditorJoinEmailList extends React.Component {
    render = () => (
        <EditorSection
            title={"JoinEmailList"}>
            <EditorInput title={"Prompt Text"} />
            <EditorInput title={"Button Text"} />
        </EditorSection>
    )
}

export default Editor;
