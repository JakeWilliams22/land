import React from 'react';
import './AdminApp.css';
import Editor from "./components/Editor/editor";

class AdminApp extends React.Component {
    render = () =>
        (
            <div className="AdminApp">
               <Editor />
            </div>
        )
}

export default AdminApp;
