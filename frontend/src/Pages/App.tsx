import React, { Component} from 'react'
import {PrimaryButton, Label, imgProperties} from 'office-ui-fabric-react'
import AppHeaderBar from '../components/AppHeaderBar'


interface IProps{
    
}

interface IState{
    second : number
}

class App extends Component<IProps,IState> {
    public readonly state : IState = {
        second: 0
    }

    constructor(props : IProps){
        super(props)
    }

    render(): JSX.Element {
        
        return (
            <div>
                <AppHeaderBar/>
                <h1 className='red'>this is a ts react page!</h1>
                <h2>this.state.second={this.state.second}</h2>
            </div>
        )
    }
}

export default App