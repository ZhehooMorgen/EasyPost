import React, { Component } from "react";


class App extends Component {
    render() {
        console.log(this.props.children)
        return (
            <div>
                <h1>Title</h1>
                {this.props.children}
                <h2>Footer</h2>
            </div>
        )
    }
}

export default App