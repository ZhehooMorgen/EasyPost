import React, { Component } from 'react'
import './Base.scss'

export default class Base extends Component {
    render(){
        return RenderSingle(1,20)
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
    return <div key={index}className="workbench">
        {viewArray}
    </div>
}