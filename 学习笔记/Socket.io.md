### Socket.io学习

Socket.IO is composed of two parts:

- A server that integrates with (or mounts on) the Node.JS HTTP Server [socket.io](https://github.com/socketio/socket.io)
- A client library that loads on the browser side [socket.io-client](https://github.com/socketio/socket.io-client)(also can loads on Node.js)

#### Not use Socket.io

* 没有使用Socket.io的普通通信

```javascript
const express = require(`express`)
const app = express();//app可以理解为一个function handler，可以作为参数传递给http连接。
const http = require('http');
const server = http.createServer(app);

app.get('/', (req, res) => {
    res.sendFile(__dirname + `/index.html`);
});

server.listen(3000, () => {
    console.log('listening on *:3000');
});
```

* 使用Socket.io的通信过程

相当于在上面的普通http代码的层面上又加了一层封装

```javascript
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
//监听connection事件的发生。
io.on('connection', (socket) => {
    console.log(`a user connected`);
});
server.listen(3000, () => {
    console.log('listening on *:3000');
});
```

#### The main Idea

The main idea behind Socket.IO is that you can send and receive any events you want, with any data you want. Any objects that can be encoded as JSON will do, and binary data is supported too.

* 服务器可以通过.emit()方法主动发送消息。

```javascript
io.emit('some event', { someProperty: 'some value', otherProperty: 'other value' });
// This will emit the event to all connected sockets
```

* Socket对象在服务器和客户端都继承了EventEmitter类，所以可以执行以下操作：

sending an event is done with: `socket.emit()`

receiving an event is done by registering a listener: `socket.on(<event name>, <listener>)`

#### Homework

- Broadcast a message to connected users when someone connects or disconnects.(√)

- Add support for nicknames.
- Don’t send the same message to the user that sent it. Instead, append the message directly as soon as he/she presses enter.(√)
- Add “{user} is typing” functionality.——暂时不太熟悉html5，不知道怎么实现用户正在输入功能。
- Show who’s online.——有些问题，不知道怎么解决
- Add private messaging.——涉及到命名空间的概念
- Share your improvements!

