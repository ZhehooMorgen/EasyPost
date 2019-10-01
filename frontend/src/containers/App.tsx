import React, {Component} from 'react'

class App extends Component{
    render(){
        return (
            <div>
                <h1>head</h1>
                {this.props.children}
                <h2>foot</h2>
            </div>
        )
    }
}

export default App