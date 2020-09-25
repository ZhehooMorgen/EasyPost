import { Icon, IconNames } from 'office-ui-fabric-react';
import React, { Component } from 'react';
import FileSys, { FileNode } from '../../../lib/FileSys'
import { FuncItem, LineElement } from './def';

import './Docs.scss'

interface Node {
    Ref: FileNode
    Expand: boolean
}

export default class Docs extends FuncItem {
    state: {
        Node: Node
    } = {
            Node: {
                Ref: FileSys,
                Expand: true
            }
        }
    Render(component: Component): JSX.Element {

        return <div>
            <LineElement
                contents={
                    <div>
                        <div>
                            <Icon iconName={IconNames.ChevronRight} className="chevron" style={{
                                transform: this.state.Node.Expand && "rotate(90deg)",
                                fontSize: "11px"
                            }} onClick={() => { this.state.Node.Expand = !this.state.Node.Expand; component.setState({}) }}
                            />
                        </div>
                        <div>test_test</div>
                    </div>
                }
                end={<div>!</div>}
            />
            <LineElement
                indent={1}
                contents={
                    <div>
                        <div>
                            <Icon
                                iconName={IconNames.ChevronRight} className="chevron" style={{
                                    transform: this.state.Node.Expand && "rotate(90deg)",
                                    fontSize: "11px"
                                }}
                                onClick={() => { this.state.Node.Expand = !this.state.Node.Expand; component.setState({}) }}
                            />
                        </div>
                        <div>test_test_________________abc</div>
                    </div>
                }
                end={<div>!</div>}
            />
        </div>
    }
    Icon = IconNames.FabricDocLibrary
    OnShow() {
        console.log("doc show")
    }
    OnHide() {
        console.log("doc hide")
    }
}

