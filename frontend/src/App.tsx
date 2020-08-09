import React, { Component } from 'react'
import AppBar from './components/AppBar'
import './App.scss'

class App extends Component {

    render(): JSX.Element {

        return (
            <div id="fullPage">
                <AppBar />
                <div>
                </div>

            </div>
        )
    }
}

export default App