# NodeJs学习

16、17号任务：初步了解nodejs，并尝试写一些小的demo，感受一下如何使用与运行代码。
后期任务：学习一下别人写的个人博客代码，大致了解后再尝试完成自己的任务。

## 创建node.js的目的

用js写一个高性能的Web服务器

由于io读写的速度限制，传统的多线程语言在处理高并发时往往会造成大量的空间浪费。而node.js在服务器端采用单线程模式，同时采用高性能的V8引擎，优化Web服务器的处理速度。

## Node.js基础学习

每个js文件就是一个模块，同时每个js文件中的代码都是运行在一个函数中的。（只是这个函数我们看不见，可以理解为一个文件为一个闭包）

可以使用require()函数引入外部模块，用exports关键字暴露变量给其他js文件。

只需要将需要暴露给外部的变量或者方法设置为exports的属性即可。

```javascript
exports.x = "我是需要导出的变量";
```

#### 模块

在Node.js的CommonJS标准中，每个js文件都是一个模块。

当node在执行模块中的代码时，它会首先在代码的最顶端，添加如下代码：

```javascript
function(exports, require, module, _filename, _dirname){
    //这里面才是我们自己写的代码！！！
}
```

![image-20211018205245074](C:\Users\lingwu\AppData\Roaming\Typora\typora-user-images\image-20211018205245074.png)

#### 包

包可以理解为规范的模块。包有公认的包结构，可以被开发者调用。

![image-20211018213136936](C:\Users\lingwu\AppData\Roaming\Typora\typora-user-images\image-20211018213136936.png)

其中，只有package.json是必须的！

![image-20211018213524080](C:\Users\lingwu\AppData\Roaming\Typora\typora-user-images\image-20211018213524080.png)

name：包标识（require()的参数）

#### buffer

由于数组中不能存储二进制文件，而buffer就是专门用来存储二进制文件数据的。（例如：图片，mp3文件等）

但是显示的时候都是以十六进制来显示的。

使用buffer不需要引入模块，直接使用即可。

buffer一个元素占用一个字节！

创建方式：

```javascript
Buffer.from(str)
Buffer.allco(size)
Buffer.alloUnsafe(size)

//用buf.toString()将缓冲区中的数据转换成字符串
var buf4 = Buffer.from("我是一段数据")
console.log(buf4.toString())
```

#### events（事件触发器）

Node.js中有些对象可以触发事件，他们都是`EventEmitter`类的实例。这些对象暴露了 `eventEmitter.on()` 函数，允许将一个或多个函数绑定到对象触发的命名事件。 通常，事件名称是驼峰式字符串，但也可以使用任何有效的 JavaScript 属性键。

当 `EventEmitter` 对象触发事件时，所有绑定到该特定事件的函数都会被同步地调用。 被调用的监听器返回的任何值都将被忽略和丢弃。

`eventEmitter.emit()` 方法允许将任意一组参数传给监听器函数。 记住，当调用普通的监听器函数时，标准的 `this` 关键字会被有意地设置为引用监听器绑定到的 `EventEmitter` 实例。

```javascript
const myEmitter = new MyEmitter();
myEmitter.on('event', function(a, b){
    console.log(a, b, this, this === myEmitter);
   //打印：
   //   a b MyEmitter {
   //     domain: null,
   //     _events: { event: [Function] },
   //     _eventsCount: 1,
   //     _maxListeners: undefined } true
});
myEmitter.emit('event', 'a', 'b');
```

#### 文件系统

需要先引入核心模块——fs

```javascript
var fs = require("fs");
```

支持同步和异步两种执行方式

常见的步骤：

1. 打开文件

```javascript
fs.open(path[,flags[,mode]], callback);//会返回一个文件描述符（数字）
```

2. 写入内容

```javascript
fs.writev(fd, buffers[,position], callback);
```

3. 关闭文件

```javascript
fs.close(fd[,callback]);
```

* 回调函数有两个参数（err, df）

```javascript
fs.open("hello.txt", "w", function(err, fd){
    if(!err){
        console.log(fd);
    }else{
        console.log(err);
    }
});
```

异步操作会将IO部分交给IO回调线程池调用，执行完毕之后将结果通过回调函数进行返回给主程序，并不会将主程序阻塞。

回调函数相对于主程序来说，会放在较后的时候执行！

* 更快捷的写法

```javascript
fs.writeFile("hello3.txt", "这是一条内容", {flag: "r+"},function(err){
    if(!err){
        console.log("写入成功");
    }else{
        console.log(err);
    }
})
```

* 流式输入

```javascript
var ws = fs.createWriteStream("hello3.txt");
//通过文件流操作进行输入
ws.write("通过可写流写入文件的内容");
ws.end();//关闭本地的接口
ws.close();//关闭另一个端口的接口
```

#### 流（基本知识）

* 流的四种基本类型

![](D:\lingwu\DianGroupProjectTask\学习图片\流的四种基本类型.png)

##### 可读流

可读流的作用是作为上游，提供数据给下游。

创建一个Readable实例后，一般需要提供一个_read方法，在这个方法中调用push产生数据。

```javascript
var Stream = require('stream')

var readable = Stream.Readable()

var source = ['a', 'b', 'c']

readable._read = function () {
  this.push(source.shift() || null)
}

```

或者：

```javascript
var Stream = require('stream')

var source = ['a', 'b', 'c']
var readable = Stream.Readable({
  read: function () {
    this.push(source.shift() || null)
  },
})
```

end事件

在**数据被消耗完**时，会出发end事件。下面两个条件都满足，表示数据被消耗完。

* 已经调用push(null),声明不会再有任何新的数据产生
* 缓存中的数据也被读取完

data事件

![](D:\lingwu\DianGroupProjectTask\学习图片\data事件.png)

总结：

* 使用push来产生数据
* 必须调用push(null)来结束流，否则下游会一直等待。
* push可以同步调用，也可以异步调用。
* end事件表示可读流中的数据已被完全消耗。
* 通常在flowing模式下使用可读流

### RPC通信

* 原理图：

![](D:\lingwu\DianGroupProjectTask\学习图片\RPC通信原理.jpg)

全称：Remote Procedure Call：远程方法调用。意为：调用其他进程或者机器上的函数。

* 分类

主流的RPC框架分为基于HTTP和基于TCP两种。

由于TCP通道只能传输二进制数据，所以需要传输的数据需要进行序列化和反序列化（文件本身格式和二进制格式之间转换），而描述两者转换之间的规范就是协议。JSON就是一种常用的转换协议。

一个完整的 RPC 框架主要有三部分组成：通信框架、通信协议、序列化和反序列化格式。

* 作用

可以在**本地**调用**位于服务器**中处理逻辑的函数。从而实现一种似乎在本地调用一样的效果。

#### Buffer使用规则

![](D:\lingwu\DianGroupProjectTask\学习图片\buffer替换规则.png)

```javascript
const fs = require('fs');
const inputStream = fs.createReadStream('input.txt'); // 创建可读流
const outputStream = fs.createWriteStream('output.txt'); // 创建可写流
inputStream.pipe(outputStream); // 管道读写
```

在Stream中我们是不需要手动去创建自己的缓冲区，在Node.js的流中会自动创建。

#### 实现

一个简单的案例

整体通过node.js提供的net模块方法来完成

* 服务器端

1. 使用`net.createServer(function(socket))`来创建一个tcp服务器。通过socket向用户端发送数据。
2. 通过`socket.on`为用户端设置一个data事件的监听器。监听器向buffer写入数据
3. 开启监听对应端口号

* 用户端

1. 通过`const socket = new net.Socket({})`创建一个套接字。
2. 通过`socket.connet()`进行对应ip地址和端口号连接。
3. `socket.write()`向socket中写入数据
4. 通过`socket.on()`向服务器端发送数据

