import { Icon, IconNames } from 'office-ui-fabric-react';
import React, { Component } from 'react';
import FileSys, { FileNode } from '../../../lib/FileSys'
import { FuncItem, LineElement } from './def';

import './Docs.scss'

interface Node {
    Ref: FileNode
    Expand: boolean //cannot expand (do not have children) if null
    Children: Array<Node>   //need to load children (not no children) if null, no existing children if empty
}

export default class Docs extends FuncItem {
    state: {
        Node: Node
    }={
        Node:null
    }
    Render(component: Component): JSX.Element {
        let indexer = 0
        let genUX = (node: Node, depth: number): JSX.Element[] => {
            let elements = new Array<JSX.Element>()
            elements.push(
                <LineElement
                    key={indexer++}
                    indent={depth}
                    contents={
                        <div>
                            <Icon
                                iconName={node.Expand !== true ? node.Expand === false ? IconNames.FabricFolder : IconNames.TextDocument : IconNames.FabricOpenFolderHorizontal} style={{
                                    fontSize: "11px"
                                }} />
                            <div>{node.Ref.name}</div>
                        </div>
                    }
                    end={<Icon
                        iconName={IconNames.Cancel} style={{
                            fontSize: "11px"
                        }} />}
                />
            )
            if (node.Expand === true) {
                node.Children.forEach((node) => {
                    elements = elements.concat(genUX(node, depth + 1))
                })
            }
            return elements
        }

        return <div>
            {genUX(this.state.Node, 0)}
        </div>
    }
    Icon = IconNames.FabricDocLibrary
    OnShow() {
        console.log("doc show")
        this.state.Node = Docs.convertFromFileInfo(FileSys)
    }
    OnHide() {
        console.log("doc hide")
    }

    private static convertFromFileInfo(fileInfo: FileNode): Node {
        let node: Node
        node = {
            Ref: fileInfo,
            Expand: null,
            Children: null
        }
        if (fileInfo.content instanceof Array) {
            node.Expand = false
            node.Children = []
            fileInfo.content.forEach(child => {
                node.Children.push(this.convertFromFileInfo(child))
            })
        }
        return node;
    }
}

