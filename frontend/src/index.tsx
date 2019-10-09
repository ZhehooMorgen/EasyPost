import Print from './lib/print'
import React from 'react'
import ReactDOM from 'react-dom'
import App from './Pages/App'

let str : string = "success"

ReactDOM.render(
    <div>{Print(str)} <App/></div>,
    document.getElementById('root')
)