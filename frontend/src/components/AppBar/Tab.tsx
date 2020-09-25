import React, { Component } from 'react'
import { Icon } from 'office-ui-fabric-react'
import './Tab.scss'

interface Props {
    title?: string
    selected?: boolean
    onDismiss?: () => void
    onClick?: () => void
}

export default class Tab extends Component<Props>{
    render() {
        return (
            <div className={this.props.selected ? "TabSelected" : "Tab"} onClick={() => this.props.onClick()}>
                <div className="TabContent">{!!this.props.title ? this.props.title : "tab 标题"}</div>
                <div className="iconWarpper">
                    <Icon iconName="Clear" className="close" onClick={(event) => { this.props.onDismiss(); event.preventDefault() }} />
                </div>
            </div>
        )
    }
}
