export interface FileNode {
    name: string
    content: FileNode[] | string
}

let FileSys: FileNode = {
    name: "root",
    content: [
        {
            name: "code",
            content: [
                {
                    name: "abc.txt",
                    content: "this is abc.txt\n"
                },
                {
                    name: "empty.tsx",
                    content: "this is empty.tsx"
                }
            ]
        },
        {
            name: "empty",
            content: []
        },
        {
            name: "fileMarkDown.md",
            content: "# Title\
# 中文标题\
## Sub tile\
中文内容\
\*斜体\*\
\*\*粗体\*\*\
end"
        }
    ]
}

export default FileSys