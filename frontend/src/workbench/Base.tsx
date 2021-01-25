import React, { Component } from 'react'
import './Base.scss'

export default abstract class Base<T> extends Component<{data:T}> {}

export class BaseAny extends Base<any>{
    render(){
        return RenderSingle(1,20,this.props.data)
    }
}

function RenderSingle(index: number, count: number, data:any): JSX.Element {
    let viewArray: JSX.Element[] = []
    for (let i = 0; i < count; i++) {
        viewArray.push(
            <div key={i}>
                <div>this is {index}, repeat for {count} times</div>
                <div>
                    {data}
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