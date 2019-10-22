import React, { Component } from 'react'
import {IconButton, SearchBox} from 'office-ui-fabric-react'
import Avatar from './Avatar'


import './AppBar.css'


export default class AppBar extends Component {
    render() {
        return (
            <div className='AppBar'>
                
                <IconButton
                className = 'MenuButton' 
                iconProps={{iconName:'GlobalNavButton'}}
                />

                <div className = 'Title'>Easy Post!博客</div>

                <SearchBox className = 'SearchBox'
                 placeholder = '搜索你感兴趣的文章、用户、内容……'
                 underlined={true}
                />

                <Avatar>Avatar</Avatar>
            </div>
        )
    }
}
