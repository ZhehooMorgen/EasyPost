import React, { Component } from 'react'
import AppBar from './components/AppBar'
import './App.scss'
import Navigator from './components/Navigator/index';
import { DefaultLayout } from './components/Navigator';
import Container from './components/Container/Container';

export default class App extends Component {
    render(): JSX.Element {
        return (
            <div id="fullPage">
                <AppBar />
                <div id="functionArea">
                    <Navigator UXFuncModules={DefaultLayout}></Navigator>
                    <Container></Container>
                </div>
            </div>
        )
    }
}
