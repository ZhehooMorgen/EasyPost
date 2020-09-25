import React, { Component } from 'react'
import './Navigator.scss'
import { IconNames, Icon } from 'office-ui-fabric-react';
import { FuncItem, Docs } from './FuncItems';

interface props {
    UXFuncModules: FuncItem[]
}

export default class Navigator extends Component<props> {
    state: {
        chosen: number
    } = {
            chosen: 0,
        }
    render() {
        this.props.UXFuncModules[this.state.chosen].OnShow()
        let fList = this.props.UXFuncModules.map((func, index) => {
            return <Icon iconName={func.Icon} className="FuncIcon" key={index} onClick={() => {
                this.switchFuncUX(index)
            }} />
        })
        return <div id='navigator'>
            <div id='funcNav'>
                {fList}
                <div id="flexHolder"></div>
                <Icon iconName={IconNames.Signin} className="FuncIcon" />
                <Icon iconName={IconNames.Settings} className="FuncIcon" />
            </div>
            <div id='UxArea'>
                {this.props.UXFuncModules[this.state.chosen].Render(this)}
            </div>
        </div>
    }

    switchFuncUX(index: number) {
        if (!(index >= 0 && index < this.props.UXFuncModules.length)) {
            throw "selected"
        }
        this.props.UXFuncModules[this.state.chosen].OnHide()
        this.setState({
            chosen: index
        })
    }
}



export const DefaultLayout = [new Docs(), new Docs(), new Docs()]
export * from './FuncItems/def'