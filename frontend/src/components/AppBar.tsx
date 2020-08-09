import React, { Component } from 'react'
import { Icon, SearchBox } from 'office-ui-fabric-react'

import './AppBar.scss'

export default class AppBar extends Component {
    render() {
        return (
            <div id='AppBar'>
                <div id="iconWarpper">
                    {/* <Icon iconName="OfficeAddinsLogo" id='icon' /> */}
                    <div id='icon'>田</div>
                </div>
                <div className='Title'>Easy Post!</div>
            </div>
        )
    }
}
