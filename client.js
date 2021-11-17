const net = require('net');
//创建一个套接字Socket
const socket = new net.Socket({});
//与对应ip地址以及端口号建立连接。
socket.connect({
    host: '127.0.0.1',
    port: 4000
});

const lessonids = [
    "136797",
    "136798",
    "136799",
    "136800",
    "136801",
    "136803",
    "136804",
    "136806",
    "136807",
    "136808",
    "136809",
    "141994",
    "143517",
    "143557",
    "143564",
    "143644",
    "146470",
    "146569",
    "146582"
]

let id = Math.floor(Math.random() * lessonids.length);

//向Socket对象中写入内容
socket.write(encode(id));


function encode(index){
    buffer = Buffer.alloc(4);
    buffer.writeInt32BE(//向buffer里面写内容的函数
        lessonids[index]
    );
    return buffer;//返回一个buffer
}

//向服务器发送数据
socket.on('data', (buffer) => {
    console.log(buffer.toString())
    id = Math.floor(Math.random()*lessonids.length);
    socket.write(encode(id));
})