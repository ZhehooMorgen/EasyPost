import React, { Component } from 'react'
import './Container.scss'

export default class Container extends Component<{}, {
    index: number
}> {
    state = { index: 0 }
    render() {
        let rd = []
        for(let i=0;i<10;i++){
            rd.push(<div style={i===this.state.index?{}:{display:"none"}} className="containerWarrper">{RenderSingle(i,i*3)}</div>)
        }
        return <div id="AppContainer"
            onClick={() => {
                if (this.state.index === 9) {
                    this.setState({ index: 0 })
                } else {
                    this.setState({ index: this.state.index + 1 })
                }
            }}>
            {rd}
        </div>
    }
}

function RenderSingle(index: number, count: number): JSX.Element {
    let viewArray: JSX.Element[] = []
    for (let i = 0; i < count; i++) {
        viewArray.push(
            <div key={i}>
                <div>this is {index}, repeat for {count} times</div>
                <div>
                    {index} is good
                </div>
                <div>
                    {Math.random().toString(36).slice(-8)}
                </div>
            </div>
        )
    }
    return <div key={index}  className="editor">
        {viewArray}
    </div>
}

