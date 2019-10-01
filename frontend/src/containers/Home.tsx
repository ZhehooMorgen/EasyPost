import React, { Component } from 'react'

class Home extends Component {
    render() {
        return (
            <div>
                <div>welcome!</div>
                {this.props.children}
                <div>this is welcome page</div>
            </div>
        )
    }
}

export default Home