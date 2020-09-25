import { Icon, IconNames } from 'office-ui-fabric-react';
import React, { Component } from 'react';
import './def.scss'


export class FuncItem {
    Render(component: Component): JSX.Element {
        return null
    }
    readonly Icon: IconNames
    OnShow() { }
    OnHide() { }
}

interface lineElementProps {
    indent?: number
    contents?: JSX.Element
    end?: JSX.Element
}

export class LineElement extends Component<lineElementProps>{
    render() {
        return <div className="lineElement">
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



