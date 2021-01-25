import React, { Component } from 'react'
import './Base.scss'

export default interface Base<T> {
    (data: T): JSX.Element
}

export function BaseAny(data: any){
    return RenderSingle(1, 20, data)
}

export function BaseString(data:string){
    return RenderSingle(2,19,data)
}

function RenderSingle(index: number, count: number, data: any): JSX.Element {
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
    return <div key={index} className="workbench">
        {viewArray}
    </div>
}