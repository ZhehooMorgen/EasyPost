import React, { Component } from 'react'
import './Navigator.scss'
import { IconNames, Icon } from 'office-ui-fabric-react';
import { FuncItem, Docs } from './FuncItems';

interface props {
    UXFuncModules: FuncItem<any>[]
}

export default class Navigator extends Component<props> {
    state: {
        chosen: number
    } = {
            chosen: null,
        }
    componentDidMount() {
        this.switchFuncUX(0)
    }
    render() {
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
                {this.state.chosen != null && this.props.UXFuncModules[this.state.chosen].Render(this, (i) => {
                    console.log("addTab: " + i)
                })}
            </div>
        </div>
    }

    switchFuncUX(index: number) {
        if (!(index >= 0 && index < this.props.UXFuncModules.length)) {
            throw "invalid selection"
        }
        if (this.state.chosen != null) {
            this.props.UXFuncModules[this.state.chosen].OnHide()
        }
        this.props.UXFuncModules[index].OnShow()
        this.setState({
            chosen: index
        })
    }
}


export const DefaultLayout = [new Docs(), new Docs(), new Docs()]
export * from './FuncItems/def'