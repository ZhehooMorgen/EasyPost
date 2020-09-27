import React, { Component } from 'react'
import { Icon } from 'office-ui-fabric-react'
import Tab from './Tab'

import './AppBar.scss'

interface state {
    tabInfo: Array<string>,
    selection: number,
}

export default class AppBar extends Component<{}, state> {
    constructor(props: {}) {
        super(props)
        this.state = {
            tabInfo: ['fs', 'test-selection.md', 'close-test.tsx', 'demo',],
            selection: 0,
        }
    }

    tabClosed(index: number) {
        this.state.tabInfo.splice(index, 1)
        if (index <= this.state.selection) {
            if (this.state.selection != 0) {
                this.setState({
                    selection: this.state.selection - 1
                })
            } else {
                this.setState({
                    selection: 0
                })
            }
        }else{
            this.setState({
                selection:this.state.selection
            })
        }
    }

    tabSelected(index:number){
        this.setState({
            selection:index
        })
    }

    render() {
        let tabs = this.state.tabInfo.map((value, index) =>
            <Tab
                key={index.toString()}
                title={value}
                selected={index == this.state.selection ? true : false}
                onDismiss={() => this.tabClosed(index)}
                onClick={()=>this.tabSelected(index)}
            />
        )
        return (
            <div id='AppBar'>
                <div id="iconWarpper">
                    <Icon iconName="OfficeAddinsLogo" id='icon' />
                </div>
                <div>Easy Post!</div>
                <hr />
                {tabs}
            </div>
        )
    }
}
