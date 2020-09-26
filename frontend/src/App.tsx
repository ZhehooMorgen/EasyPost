import React, { Component } from 'react'
import AppBar from './components/AppBar'
import './App.scss'
import Navigator from './components/Navigator/index';
import { DefaultLayout, FuncItem } from './components/Navigator';

class App extends Component {

    render(): JSX.Element {

        return (
            <div id="fullPage">
                <AppBar />
                <div id="functionArea">
                    <Navigator UXFuncModules={DefaultLayout}></Navigator>
                </div>

            </div>
        )
    }
}

export default App