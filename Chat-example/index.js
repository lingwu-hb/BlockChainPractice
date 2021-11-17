const express = require(`express`)
const app = express();//app可以理解为一个function handler，可以作为参数传递给http连接。
const http = require('http');
const server = http.createServer(app);
const { Server } = require("socket.io");
const io = new Server(server);//用http的server来初始化socket.io的server

app.get('/', (req, res) => {
    res.sendFile(__dirname + `/index.html`);
});
//.on方法就是由socket.io模块所提供的。
io.on('connection', (socket) => {
    console.log(`a user connected:` + socket.id);
    //告诉其他已经连接的用户，someone连接了服务器
    socket.broadcast.emit(`someone has connected`, socket.id);
    io.emit(`Show Online`, io.sockets.sockets);
    socket.on('disconnect', () => {
        console.log('user disconnected');
    });
});

server.listen(3000, () => {
    console.log('listening on *:3000');
});