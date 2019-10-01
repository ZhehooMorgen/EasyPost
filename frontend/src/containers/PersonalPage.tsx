import React, { Component } from 'react'
import { PrimaryButton } from 'office-ui-fabric-react'

class PersonalPage extends Component {
    render() {
        return (
            <div>
                <div>This is your personal page</div>
                <PrimaryButton onClick={()=>this.testReq()}>
                    test
                </PrimaryButton>
            </div>
        )
    }

    testReq() {
        let req = new XMLHttpRequest()
        req.onreadystatechange = () => {
            if (req.readyState == 4) {
                if (req.status == 200)
                    alert("great!")
            }
        }
        req.open("POST","http://localhost:86/api/login")
        req.send(`{"userName":"aa","password":"ss"}`)
    }
}

export default PersonalPage
