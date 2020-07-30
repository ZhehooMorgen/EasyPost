import React, { Component } from 'react'
import AppBar from './components/AppBar'
import Editor from './components/Editor';
import './App.css'

class App extends Component {

    render(): JSX.Element {

        return (
            <div id="fullPage">
                <AppBar />
                <div>

                    <Editor></Editor>
                </div>

            </div>
        )
    }
}

export default App