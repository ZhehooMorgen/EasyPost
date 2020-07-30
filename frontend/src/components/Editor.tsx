import React, { Component } from 'react'
import { IconButton, SearchBox } from 'office-ui-fabric-react'
import './Editor.css'

import SimpleMDE from './SimpleMDE';
//import './SimpleMDE/easymde.css'
import './SimpleMDE/easymde.css'



export default class Editor extends Component{
    render(){
        return <SimpleMDE
        label="test"
        value={demoText}
        extraKeys={{
            "Ctrl-S":function(cm){
                console.log("pressed ctrl+s")
            }
        }}
        options={{
            spellChecker:false
        }}
        />
    }
}











const demoText = `
# test H1
## h2
### h3
#### h4
##### h5
<h5>test html h5</h5>
**test blod text **

*test 斜体*
中文
关闭spell chaaker

`