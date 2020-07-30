import React, { Component } from 'react'
import { IconButton, SearchBox } from 'office-ui-fabric-react'
import Avatar from './Avatar'

import './AppBar.css'

export default class AppBar extends Component<{}, { hide: boolean }> {
    constructor(props: Readonly<{}>) {
        super(props)
        this.state = {
            hide: false
        }
    }
    componentDidMount() {
        this.ticker = setInterval(() => {
            let placeHolder = document.getElementById("AppbarPlaceHolder")
            if (placeHolder) {
                let toTop = placeHolder.getBoundingClientRect().top
                //console.log(toTop)
                if (toTop < -50) {
                    this.setState({
                        hide: true
                    })
                } else {
                    this.setState({
                        hide: false
                    })
                }
            }
        }, 100)
    }
    componentWillUnmount() {
        clearInterval(this.ticker)
    }
    ticker: NodeJS.Timeout
    render() {
        return (
            <div>
                <div id='AppBar'
                    style={this.state.hide ? { top: "-50px" } : {}}
                >
                    <IconButton
                        id='AppBarMenuButton'
                        iconProps={{
                            iconName: 'GlobalNavButton',
                            style: {
                                fontSize:24
                            },
                        }}
                    />

                    <h2 className='Title'>Easy Post!</h2>

                    <SearchBox className='SearchBox'
                        placeholder='搜索你感兴趣的文章、用户、内容……'
                        underlined={true}
                        style={{
                            color:"white"
                        }}
                        onSearch={(text)=>{
                            window.open("http://www.google.com/search?q="+text)
                        }}
                    />

                    <Avatar>Avatar</Avatar>
                </div>
                <div id='AppbarPlaceHolder'>
                        You should not see this!
                </div>
            </div>
        )
    }
}
