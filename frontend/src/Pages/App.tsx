import React, { Component } from 'react'
import {PrimaryButton, Label} from 'office-ui-fabric-react'
import './App.css'

let time: number = 0

setInterval(() => {
    time++
}, 1000)

class App extends Component {
    timer : NodeJS.Timeout
    componentDidMount(){
        this.timer = setInterval(()=>this.setState({}),1000)
    }
    componentWillUnmount(){
        clearInterval(this.timer)
    }
    render(): JSX.Element {
        return (
            <div>
                <h1 className='red'>this is a ts react page!</h1>
                <h2>run for {time} seconds</h2>
                <PrimaryButton className='red'>哈！</PrimaryButton>
                <Label>fsfsf</Label>
            </div>
        )
    }
}

export default App