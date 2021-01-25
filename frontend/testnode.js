const net = require('net');

let server = net.createServer((c) => {
    console.log("connect")
})
server.on('error', (err) => {
    throw err;
})
server.listen(4000, () => {
    console.log('服务器已启动');
});