## javascript基础语法学习

对于没有赋值的变量，javascript会将其值设定为undefined

但是常量对象的属性可以进行修改。

```javascript
// 您可以创建 const 对象：
const car = {type:"porsche", model:"911", color:"Black"};

// 您可以更改属性：
car.color = "White";

// 您可以添加属性：
car.owner = "Bill";
```

同时，常量数组也可以进行修改。

JavaScript 拥有动态类型。这意味着相同变量可用作不同类型。最终类型会与最后赋值的一个相同。

### let和const

let用于创建一个具有块作用域的变量。

而const必须在创建时赋初值，而且不能改变该对象指向的内存地址。

let和const修饰的变量的作用域都只在一个块内，不同的是，const定义的变量不能够再次修改。

### Undefined 与 Null 的区别：

Undefined 与 null 的值相等，但类型不相等：

```javascript
typeof undefined              // undefined
typeof null                   // object
null === undefined            // false
null == undefined             // true
```

请不要把字符串创建为对象。它会拖慢执行速度。例如： var y = new String("hanbo")

new 关键字使代码复杂化。也可能产生一些意想不到的结果。

```javascript
== 运算符和 === 运算符的区别

值相等，==运算符就会返回true
但是只有值和类型都相等，===运算符才会返回true

注意：在javascript中，两个对象之间永远不能相互比较。将始终返回false
```

javascript提供了很多字符串的处理方法，可以在使用时进行查阅。

### 字符串模板

在字符串模板内进行字符串替换的操作叫做字符串插值，语法为：&{...}

```javascript
let firstName = "John";
let lastName = "Doe";

let text = `Welcome ${firstName}, ${lastName}!`;
```

模板字符串有很多重要的用途，具体的以后再详细了解。

### js函数

js中，函数也不例外是一种对象。

* 参数规则：

JavaScript 函数定义不会为参数（parameter）规定数据类型。

JavaScript 函数不会对所传递的参数（argument）实行类型检查。

JavaScript 函数不会检查所接收参数（argument）的数量。

语法：

```javascript
function name(参数1, 参数2, ...){
	//代码块
}
```

* 匿名函数

类似于Go语言中的匿名函数，javascript中也可以不用给函数名称，直接定义函数体，然后赋值给一个变量。

```javascript
hello = function(参数1, 参数2){
	//...
}
```

#### 箭头函数

**箭头函数表达式**的语法比[函数表达式](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Operators/function)更简洁，并且没有自己的`this`，`arguments`，`super`或`new.target`。箭头函数表达式更适用于那些本来需要匿名函数的地方，并且它不能用作构造函数。

```javascript
hello = () => {
	return "Hello World!";
}
```

对于箭头函数，this关键字始终表示定义箭头函数的对象。（这与普通函数不尽相同）

例如：

```javascript
// 常规函数：
hello = function() {
  document.getElementById("demo").innerHTML += this;
}

// window 对象调用该函数：this表示调用该函数的对象
window.addEventListener("load", hello);

// button 对象调用该函数：
document.getElementById("btn").addEventListener("click", hello);

<--!-->
    
//箭头函数:this表示函数的拥有者
hello = () => {
    document.getElementById("demo").innerHTML += this;
}
window.addEventListener("load", hello);
document.getElementById("btn").addEventListener("click", hello);
```

箭头函数有很多用处，但是同时也有很多地方需要仔细理解，具体参考链接：https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Functions/Arrow_functions

* 基础语法

```javascript
(param1, param2, …, paramN) => { statements }
(param1, param2, …, paramN) => expression
//相当于：(param1, param2, …, paramN) =>{ return expression; }

// 当只有一个参数时，圆括号是可选的：
(singleParam) => { statements }
singleParam => { statements }

// 没有参数的函数应该写成一对圆括号。
() => { statements }
```

* 高级语法

```javascript
//加括号的函数体返回对象字面量表达式：
params => ({foo: bar})

//支持剩余参数和默认参数
(param1, param2, ...rest) => { statements }
(param1 = defaultValue1, param2, …, paramN = defaultValueN) => {
statements }

//同样支持参数列表解构
let f = ([a, b] = [1, 2], {x: c} = {x: a + b}) => a + b + c;
f();  // 6
```

#### js回调函数

回调函数是作为参数传递给另一个函数的函数。

示例如下：

```javascript
function myDisplayer(some) {
  document.getElementById("demo").innerHTML = some;
}

function myCalculator(num1, num2, myCallback) {
  let sum = num1 + num2;
  myCallback(sum);
}

myCalculator(5, 5, myDisplayer);
```

### 数组

可以通过数组名直接访问数组中的全部内容！也可以采用数组提供的Foreach(callfunction)方法进行调用

数组和对象的区别
在 JavaScript 中，数组使用数字索引。

在 JavaScript 中，对象使用命名索引。

数组是特殊类型的对象，具有数字索引。

数组和字符串一样，也有很多详细的方法，具体可以通过vscode插件进行详细查看。

具体常用方法列出几个如下：

```javascript
Array.forEach(func);
Array.map(func);
Array.filter(func);//过滤函数，通过回调函数进行过滤操作。
```

* 用const定义的数组，具有块作用域。同时不能对其进行赋值操作。

注意：但是可以对const数组中的元素进行修改和增添操作。（只有数组这个对象本身是const）

### JS数学操作

javascript中Math中提供了很多方法，可以提供常见的数学操作。

```javascript
//使用方法：(其余都类似)
Math.pow(x, y);
```

### js循环

![image-20211017110710739](C:\Users\lingwu\AppData\Roaming\Typora\typora-user-images\image-20211017110710739.png)

### js正则表达式

正则表达式是构成搜索模式的字符序列。

当我需要搜索文本中的数据时，可以使用搜索模式来描述需要搜索的内容。

例如：

```javascript
var x = /w3school/i;
```

x是一个正则表达式，其中w3school是模式，i是修饰符（表示对大小写不敏感。）

正则表达式常常与字符串方法search()和replace()配合使用。例如：

```javascript
var str = "Visit W3School";
var n = str.search(/w3school/i);
```

常用的正则模式如下：

![image-20211017111544511](C:\Users\lingwu\AppData\Roaming\Typora\typora-user-images\image-20211017111544511.png)

### this

call()和apply()方法是预定义的JavaScript方法。

他们都可以用于将另一个对象作为参数调用对象方法。

```javascript
var person1 = {
  fullName: function() {
    return this.firstName + " " + this.lastName;
  }
}
var person2 = {
  firstName:"Bill",
  lastName: "Gates",
}
console.log(person1.fullName.call(person2));  // 会返回 "Bill Gates"
```

### 类

与c++类似，类不是对象，类是对象的模板，用来创建对象。

```javascript
class Car{
	constructor(name, year){
		this.name = name;
		this.year = year;
	}
}

let myCar1 = new Car("Ford", 2014);
```

在创建新对象时会自动调用constructor方法。

在类中创建方法

```javascript
class ClassName{
	constructor() {...}
	method_1() {...}
	method_2() {...}
}
```

### (JSON)javaScript Object Notation

由于TCP通道只能传输二进制数据，所以需要传输的数据需要进行序列化和反序列化（文件本身格式和二进制格式之间转换），而描述两者转换之间的规范就是协议。**JSON就是一种常用的转换协议**。

常见用法见下图：

![image-20211017213919373](C:\Users\lingwu\AppData\Roaming\Typora\typora-user-images\image-20211017213919373.png)

### js对象

* 创建对象的方法

1. 使用对象字面量

```javascript
var person = {
    firstName:"Bill",
    lastName:"Gates",
    age:62,
    eyeColor:"blue"
};
```

2. 使用关键词new

```javascript
var person = new Object();
person.firstName = "Bill";
person.lastName = "Gates";
person.age = 50;
person.eyeColor = "blue"; 
```

注意：javascript中对象是通过引用来寻址的，而不是值。（对象被存储在堆中，而不是栈中。）

```javascript
var person = {firstName:"bill"};

var x = person;
x.firstName = "Gate";//这将同时改变x和person中的成员。
```

* 对象方法：

![image-20211018162206188](C:\Users\lingwu\AppData\Roaming\Typora\typora-user-images\image-20211018162206188.png)

* js对象访问器

下面的类中有一个get访问器，用于得到全称。

```javascript
var person = {
    firstName: "Bill",
    lastName: "Gates",
    get fullName() {
        return this.firstName + " " + this.lastName;
    }
};
//以属性形式访问fullName:person.fullName
//这样就不用带括号了。
```

* 为什么使用Getter和Setter？

  

  它提供了更简洁的语法

  它允许属性和方法的语法相同

  它可以确保更好的数据质量

  有利于后台工作

![image-20211018164426466](C:\Users\lingwu\AppData\Roaming\Typora\typora-user-images\image-20211018164426466.png)

* 类型构造器

与关键字new配合使用。用于创建一个需要的对象。

js提供用于原始对象的构造器，如下：

```javascript
var x1 = new Object();    // 一个新的 Object 对象
var x2 = new String();    // 一个新的 String 对象
var x3 = new Number();    // 一个新的 Number 对象
var x4 = new Boolean();   // 一个新的 Boolean 对象
var x5 = new Array();     // 一个新的 Array 对象
var x6 = new RegExp();    // 一个新的 RegExp 对象
var x7 = new Function();  // 一个新的 Function 对象
var x8 = new Date();      // 一个新的 Date 对象
```

基本类型不建议使用new＋构造函数来创建对象，这样会产生一些不必要的问题。

* 使用prototype属性

js prototype属性允许为对象构造器添加新属性。

```javascript
function Person(first, last, age, eyecolor){
    this.firstName = first;
    this.lastName = last;
    this.age = age;
    this.eyeColor = eyecolor;
}
Person.prototype.nationality = "English";
//也可以添加新方法
Person.prototype.name = function(){
    return this.firstName + " " + this.lastName;
};//分号别忘了
```

> 请只修改您自己的原型。绝不要修改标准 JavaScript 对象的原型。

* ECMAScript 5 (2009) 向 JavaScript 添加了许多新的对象方法。

语法：

```javascript
//以更改属性值为例：
Object.defineProperty(object, property, {value: value};)
```

#### promise对象

promise对象主要是为了处理js中回调地狱的问题。

promise对象在创建时需要加一个 resolve，其可以**将异步数据传递出来**

```javascript
let p = new Promise(function(resolve){
    resolve("hello World")
})
//通过then方法拿到异步数据
p.then(function(data){
    console.log(data)
})
//通过.then方法就可以拿到数据
//将回调地狱问题变为链式结构，得以优化
```

### js事件

每一个可用的事件都会有一个可用的事件处理器，，也就是事件触发时会运行的代码块。当我们定义了一个用来回应事件被激发的代码块的时候，我们说我们**注册了一个事件处理器**。注意事件处理器有时候被叫做**事件监听器**

* data事件

`data` 事件是一种被动消费的方式，就是说一旦有数据传输过来就会吐给监听函数，监听函数必须马上处理或者把数据缓存下来，不然数据就丢失了。

* readable事件

和 `data` 事件不同，`readable` 事件不会直接返回数据给业务，它只会通知业务有新数据来了，业务可以根据需求调用 `socket.read` 去主动读取数据。read 方法接收一个可选参数 size，若当前收到字节数小于传入的 size，那 read 方法返回 null，如果不传 size 那就读取所有数据。

### DOM简介

全称：Document Object Model:文档对象模型

作用：方便js代码快速获取到需要的对象，并对其进行操作，其本身是一个树结构。共有四种节点。

1. 文档结点
2. 元素结点
3. 属性结点
4. 文本结点

![](D:\lingwu\DianGroupProjectTask\学习图片\DOM.png)

可以通过`document.getElementByld("btn")`进行对象的获取，然后再进行后续的操作。