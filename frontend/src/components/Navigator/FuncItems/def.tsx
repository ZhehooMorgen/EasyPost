import { Icon, IconNames } from 'office-ui-fabric-react';
import React, { Component } from 'react';
import { ITabOpener } from '../../../workbench';
import './def.scss'

export abstract class FuncItem<T> {
    Render(component: Component,openTab: ITabOpener<T>): JSX.Element {
        return null
    }
    readonly Icon: IconNames
    OnShow() { }
    OnHide() { }
}

export class LineElement extends Component<{
    indent?: number
    contents?: JSX.Element[]
    end?: JSX.Element
    onClick?: () => void
}>{
    render() {
        return <div
            className="lineElement"
            onClick={() => {
                if (this.props.onClick != undefined && this.props.onClick != null) {
                    this.props.onClick()
                }
            }}
        >
            <div className="forIndent" style={{
                width: "" + 20 * (this.props.indent != null ? this.props.indent : 0) + "px"
            }} />
            <div className="content">
                {this.props.contents}
            </div>
            <div className="end">{this.props.end}</div>
        </div>
    }
}



