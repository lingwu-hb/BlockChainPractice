# Go 语言的特点
![在这里插入图片描述](https://img-blog.csdnimg.cn/2d8f79f8919b4764859668b71f40e25b.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAbGluZ3d1X2hi,size_20,color_FFFFFF,t_70,g_se,x_16)
![在这里插入图片描述](https://img-blog.csdnimg.cn/2b03e3ce283346ff9e7aebc5912139f4.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAbGluZ3d1X2hi,size_20,color_FFFFFF,t_70,g_se,x_16)
# Go 语言项目开发目录结构
![在这里插入图片描述](https://img-blog.csdnimg.cn/2627dab8142642b0b8a820c9ed856075.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAbGluZ3d1X2hi,size_20,color_FFFFFF,t_70,g_se,x_16)
# 基础语法
## 变量声明

Go支持声明，可以使用var name type = xxx 来进行声明并赋值。
也可以采用 name := xxx 来进行声明并赋值

由于函数外的语句必须以关键字开始，所以 := 运算符不能出现在函数外部！

>* 也就是说，用var就不要冒号，不用var就要用冒号

```go
var a string = "hanbo"
//或者
a := "hanbo"
```
注意：下面两种都是可行的

```go
var a = "hanbo"
var b string = "hanbo"
```
* 基础变量的零值

Go语言中的零值是变量没有做初始化时系统默认设置的值。

```go
var b bool // bool型零值是false
var s string // string的零值是""
var a *int // 指针的初始值为<nil>
var a []int
var a map[string] int
var a chan int
var a func(string) int
var a error // error是接口
```

以上六种类型零值常量都是nil

所有其他数值型的类型（包括complex64/128）零值都是0，可以用常量表达式代表数值0的任何形式表示出来。

* 其他数据类型

![在这里插入图片描述](https://img-blog.csdnimg.cn/7692633ab7f0462e84bd5d3715c9c38c.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAbGluZ3d1X2hi,size_20,color_FFFFFF,t_70,g_se,x_16)
## 枚举类型
Go语言的枚举类型与c++不同，一般采用一下方式实现。
```go
package main

import "fmt"

type CMState int

const ( //Go语言中枚举类型的声明
	Follower CMState = iota
	Candidate
	Leader
	Dead
)

//实现String()方法之后，就可以按照字符串的格式进行输出了（很方便）
func (s CMState) String() string {
	switch s {
	case Follower:
		return "Follower"
	case Candidate:
		return "Candidate"
	case Leader:
		return "Leader"
	case Dead:
		return "Dead"
	default:
		panic("unreachable")
	}
}

func main() {
	fmt.Println(Dead)
}

//输出结果为：Dead

```

## 结构体
* 结构体指针中支持隐式调用，相当于移除了c中的->运算符
```go
package main
import "fmt"
type Vertex struct {
	X int
	Y int
}
func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

```
* 结构体的初始化（可以直接按照顺序依次赋值，也可以采用键值对的方式进行赋值）

```cpp
type People struct {
    name  string
    child *People
}
relation := &People{
	//键值对冒号初始化
    name: "爷爷",//键值对方式进行初始化
    child: &People{
        name: "爸爸",
        child: &People{
                name: "我",
        },
    },
}
```
* 匿名结构体

与匿名函数类似，Go语言也支持匿名结构体的使用。重点关注下面的如何初始化匿名结构体部分

```cpp
// 打印消息类型, 传入匿名结构体
func printMsgType(msg *struct {
    id   int
    data string
}) {
    // 使用动词%T打印msg的类型
    fmt.Printf("%T\n", msg)
}
func main() {
    // 实例化一个匿名结构体
    msg := &struct {  // 定义部分
        id   int
        data string
    }{  // 值初始化部分
        1024,
        "hello",
    }
    printMsgType(msg)
}
```



## Slices
切片文法类似于没有长度的数组文法。

这是一个数组文法：

```go
[3]bool{true, true, false}
```

下面这样则会创建一个和上面相同的数组，然后构建一个引用了它的切片：

```go
[]bool{true, true, false}
```
* GoLang中数组与c语言中的数组的区别

有长度检查，同时数组名不再是地址了，要获得地址需要使用&运算符

```go
// 有长度检查, 也为地址传参
func use_array(args *[4]int) {
	args[1] = 100 //但是使用还是和C一致,不需要别加"*"操作符
}

func main() {
	var args = [4]int{1, 2, 3, 4}
	use_array(&args) // 数组名已经不是表示地址了, 需要使用"&"得到地址
	fmt.Println(args)
}
```

## 分支语句
* if语句

Go语言的 if 判断语句必须要将左花括号与 if 写在同一行。同时条件的内容不能带括号

```go
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
```

if语句和c里面的for循环一样，可以进行初始化，其变量的作用域包括其他的一些else语句

```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {//if语句的初始化循环
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

* switch语句

switch 是编写一连串 if - else 语句的简便方法。它运行第一个值等于条件表达式的 case 语句。与if语句类似，GoLang中的switch语句也支持在循环开始时进行初始化工作！

Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，不过 Go 只运行选定的 case，而非之后所有的 case。 实际上，**Go 自动提供了在这些语言中每个 case 后面所需的 break 语句**。 除非以 fallthrough 语句结束，否则分支会自动终止。 **Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。**

另外，switch语句支持没有条件，相当于条件永远为真，这种写法使if-else语句更轻便！

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```

## 循环语句
* c中的while在GoLang中就是for，所以for循环时可以不带分号

```go
var i int = 11
for i > 0 {
	i--
	tudou := i
]
```

for 循环的 range 形式可遍历切片或映射。

当使用 for 循环遍历切片时，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```
## 指针
### 内存模型
* 一个函数会分配一个栈帧（在stack内部）。栈帧空间内会记录以下三种内容：
1. 形参
2. 局部变量
3. 栈基指针和栈顶指针值 
![请添加图片描述](https://img-blog.csdnimg.cn/6310e89d5f73402b8dae8eedd7cf3512.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAbGluZ3d1X2hi,size_20,color_FFFFFF,t_70,g_se,x_16)
### 指针作为函数返回值
不能够放回局部变量的地址值，只能够返回局部变量的值。因为函数执行完毕后，操作系统会回收栈帧空间，导致传出的地址里面的值随时可能被其他的程序修改，无法保证安全。
## 函数
* 函数类型

与c语言中的函数指针相似但有不同，Go语言中将函数视为一种类型，可以创建一个函数类型的变量，然后将函数赋值给该变量。

```go
func test(){
	fmt.Println("hello")
}
var f func()
f = test
f()//相当于执行了tes函数
```

### 函数参数传递
Go语言中所有的情况都是传值处理，只是有些类型本身是引用类型，所以才可以在函数内部进行修改。具体理解参见[博客](https://www.flysnow.org/2018/02/24/golang-function-parameters-passed-by-value.html)

### 匿名函数

Go语言支持匿名函数，即在需要使用函数时再定义函数，匿名函数没有函数名只有函数体，函数可以作为一种类型被赋值给函数类型的变量，匿名函数也往往以变量方式传递，这与C语言的回调函数比较类似，不同的是，Go语言支持随时在代码里定义匿名函数。

匿名函数非常好用，可以在任何地方创建，并作为一种值赋值给函数类型的变量。比较常用的地方是将其作为参数传给函数，也就是相当于c语言里面的回调函数。

例如：

```go
// 遍历切片的每个元素, 通过给定函数进行元素访问
func visit(list []int, f func(int)) {
    for _, v := range list {
        f(v)
    }
}
func main() {
    // 使用匿名函数打印切片内容
    visit([]int{1, 2, 3, 4}, func(v int) {
        fmt.Println(v)
    })
}
```
* 深入理解匿名函数

可以把匿名函数理解为一个变量，将其整体作为一个变量进行使用即可。

1）将其作为一个值赋给一个变量

```go
funValue := func(v int){
	fmt.Println(v)
}
//然后执行函数
funValue()
```

2)）或者可以在其本身后面加上括号，直接对匿名函数进行调用即可

```go
func(v int) {
	fmt.Println(v)
} ()//表示声明并调用该匿名函数
```

### 闭包

>* 闭包 = 引用环境 + 函数

```go
// 准备一个字符串
str := "hello world"
// 创建一个匿名函数
foo := func() {
   
    // 匿名函数中访问str
    str = "hello dude"
}
// 调用匿名函数
foo()
```

也就是说，如果一个正常的函数调用了不是在其函数体内部定义的变量的话，那么该函数就捕获了该变量，同时自身也叫做闭包。该变量的生存周期与函数相同

* 下面是一个自己写的输出斐波那契的运用闭包的程序语句

```go
package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	i, a, b := -1, 0, 0
	return func() int{
		i = i+1
		if i == 0 {
			return 0
		}
		if i == 1 {
			a = 1
			return 1
		}
		if i == 2 {
			b = 1
			return 1
		} else {
			c := b
			b = a + b
			a = c
			return b
		}
	}
}

func main() {
	f := fibonacci()
	fmt.Println("这是斐波那契数列")
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```

### 方法

方法就是一类带特殊的 **接收者** 参数的函数。

方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。

在此例中，Abs 方法拥有一个名为 v，类型为 Vertex 的接收者。

>* 如果方法的接收者是值而不是指针，那么即使在方法内部修改其接收者得值，该值原来的内存空间仍然没有改变！因此，在设定方法接收者时，常常使用的是指针接收者而不是值接收者！

```cpp
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {//表示函数Abs()可以被结构体v所直接用.运算符进行调用
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

Go语言的指针重定向
```cpp
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```

>* 对于语句 v.Scale(5)，即便 v 是个值而非指针，带指针接收者的方法也能被直接调用。 也就是说，由于 Scale 方法有一个**指针接收者**，为方便起见，**Go 会将语句 v.Scale(5) 解释为 (&v).Scale(5)。** 同样的重定向也可以发生在相反的方向，也就是将指针 v.scale() 解释为 (*v).scale() 并进行调用！ 

### 标准化输入输出函数
#### 输出函数简介
* Printf()函数
```go
func Printf(format string, a ...interface{}) (n int, err error)
```
与c语言的格式化输出函数相同，根据格式化表将参数进行格式解析并且输出

示例：

```go
package main

import (
	"fmt"
)

func main() {
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age)

	// It is conventional not to worry about any
	// error returned by Printf.

}
```

* Println()函数

```go
func Println(a ...interface{}) (n int, err error)
```

按照默认的格式进行解析，然后将内容输出到标准输出流中，同时自动在每个参数之间加上一个空格，以及在末尾加上一个换行符

示例：

```go
package main

import (
	"fmt"
)

func main() {
	const name, age = "Kim", 22
	fmt.Println(name, "is", age, "years old.")

	// It is conventional not to worry about any
	// error returned by Println.

}

```

#### 输入函数简介
* 参阅官方文档（重点关注各个函数**如何处理换行符和空格符**）
![在这里插入图片描述](https://img-blog.csdnimg.cn/2e6f8f9e42dd4fbf90ec150615fc633d.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAbGluZ3d1X2hi,size_20,color_FFFFFF,t_70,g_se,x_16)

### 内存分配函数
* 主要是 new() 和 make() 两者的不同之处

1）new(type) 分配一块内存空间，然后返回指向该内存的指针

2）make(type, size) 分配一块空间，返回与类型相同的值，而不是指针。

详细区别可以参阅[此博客](https://www.cnblogs.com/chenpingzhao/p/9918062.html)
# 其他特性
## defer 
后面可以接一个函数，然后Go会在该函数快要退出的时候逆序执行被defer的函数。

>* 注意：原理是被defer的函数f()被保存到了一个调用栈中，被保存的也包括f()函数中调用的对象，所以即使后面该对象被改变了，在逆序调用时，该变量的值仍然保持不变！

```cpp
func main() {
	var a int = 1
	defer fmt.Println(a)
	
	a = 2
	fmt.Println(a)
}
```
## 接口
接口是一种自定义类型（**本质**），它包含了一些方法。可以声明一个接口类型变量，该变量可以保存任何实现了这些方法的值。下面是示例：

```go
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

### 隐式接口

观察下面的程序

```go
package main
import(
	"fmt"
	"math"
)

type Vertex struct{
	X float64
	Y float64
}

type I interface {
	Abs() float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main(){
	v := Vertex{3, 4}//v类型实现了Abs()方法
	var i I = v
	i.Abs()
}
```

观察main函数可以看到，显示说明了v之后，再用 i 保存 v 的操作有些累赘，所以可以采用隐式接口的方式简化

```go
	var i I = &Vertex{3,4}//此处采用隐式调用简化程序
	fmt.Println(i.Abs())
```

### 空接口
* 定义：

前面提到接口是一种签名了一些方法的类型。而空接口就是指没有签名任何方法的特殊接口类型。

```go
type i interface{}//定义一个空接口类型
```

* 使用方法

由于所有类型都至少实现了空方法，所以空接口可以承载任何类型。

1）用空接口初始化一个实例，该实例可以保存任何类型的值

2）函数参数设置为空接口类型，这样就可以接受任何类型的参数

但是承载后其动态类型需要用类型断言来判断

参考[博客](https://juejin.cn/post/6844904183770906638)
### 类型断言
类型断言 提供了访问接口值底层具体值的方式。

```go
t := i.(T)
```

该语句**断言接口值 i** 保存了**具体类型 T**，并将其底层类型为 T 的值赋予变量 t。

若 i 并未保存 T 类型的值，该语句就会触发一个panic。（报错）

为了 判断 一个接口值是否保存了一个特定的类型，类型断言可返回两个值：其底层值以及一个报告断言是否成功的布尔值。

```go
t, ok := i.(T)
```

若 i 保存了一个 T，那么 t 将会是其底层值，而 ok 为 true。

否则，ok 将为 false 而 t 将为 **T 类型的零值**（不同类型的零值不尽相同），程序并不会产生panic。

注意这种语法和读取一个映射时的相同之处。

### Stringer接口
fmt 包中定义的 Stringer 是最普遍的接口之一。

```go
type Stringer interface {
    String() string
}
```

Stringer 是一个可以用字符串描述自己的类型。fmt 包（还有很多包）都通过此接口来打印值。

使用示例：
```go
package main

import "fmt"

type IPAddr [4]byte

// TODO: 给 IPAddr 添加一个 "String() string" 方法
func (i IPAddr) String() string{
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		//fmt.Printf在格式化输出时，会调用该类型最底层的输出函数（个人目前理解）
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```

### error（错误）
GoLang 里面的 error 类型的底层原型就是一个接口。其内部签名了一个方法。

```go
type error interface{
	Error() string
}
```
任何实现了该方法的值都可以被保存为 error 类型

```go

package main
import (
    "errors"
    "fmt"
    "math"
)
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return -1, errors.New("math: square root of negative number")
    }
    return math.Sqrt(f), nil
}
func main() {
    result, err := Sqrt(-13)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(result)
    }
}
```

## Go语言调试
使用内置的delve进行调试。

命令行进入包所在目录，然后输入dlv debug 即可进行调试

```go
$ dlv debug
Type 'help' for list of commands.
(dlv)
```

键入help即可查看所有的命令。

### 常用命令
详细用法可以使用h  < comand > 命令调出命令手册

```cpp
    args ------------------------ Print function arguments.
    break (alias: b) ------------ Sets a breakpoint.
    breakpoints (alias: bp) ----- Print out info for active breakpoints.
    clear ----------------------- Deletes breakpoint.
    continue (alias: c) --------- Run until breakpoint or program termination.
    exit (alias: quit | q) ------ Exit the debugger.
    funcs ----------------------- Print list of functions.
    goroutine (alias: gr) ------- Shows or changes current goroutine
    goroutines (alias: grs) ----- List program goroutines.
    list (alias: ls | l) -------- Show source code.
    locals ---------------------- Print local variables.
    next (alias: n) ------------- Step over to next source line.
    on -------------------------- Executes a command when a breakpoint is hit.
    print (alias: p) ------------ Evaluate an expression.
    stack (alias: bt) ----------- Print stack trace.
    step (alias: s) ------------- Single step through program.
    step-instruction (alias: si)  Single step a single cpu instruction.
    stepout --------------------- Step out of the current function.
    vars ------------------------ Print package variables.
```

## 多线程
### 相关概念
* 进程：程序运行的最小单位，一般情况下，一个程序会分配一个进程用于执行
* 线程：操作系统用于资源分配与调度的最小单位。一个进程可以有多个线程
* 协程：由用户自己分配的比线程更小的调用单位，Go语言中的goroutine就是典型的协程

综上所述：可以简单认为：进程 > 线程 > 协程

>* 进程间具有独立性，不同进程之间的变量一般是不会共享的。
>* 进程间是可以同时访问同一个变量的（有时候这种特性会引发一些问题）

* 并发：**一个处理器**在**多个线程间**反复横跳进行处理，表现出来的就是多个线程同时在进行
* 并行：**多内核处理器**处理**多个进程**时，每个处理器可以独立处理一些进程，这些进程间就是并行关系
### Go程

下面的语句会启动一个轻量级线程，然后在新的线程中执行函数 f 
```go
go f(x, y, z)
```
并发的图解
![在这里插入图片描述](https://img-blog.csdnimg.cn/ad82aa35631b446a92bda1e7e9e62fd6.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAbGluZ3d1X2hi,size_20,color_FFFFFF,t_70,g_se,x_16)
### Go程间竞争和锁机制

由于不同Go程之间访问共享变量，从而引发内存读写错误。（逻辑错误，编译器不会报错）

```go
package main
import (
    "fmt"
    "runtime"
    "sync"
)
var (
    count int32
    wg    sync.WaitGroup
)
func main() {
    wg.Add(2)
    go incCount()
    go incCount()
    wg.Wait()
    fmt.Println(count)
}
func incCount() {
    defer wg.Done()
    for i := 0; i < 2; i++ {
        value := count
        runtime.Gosched()//可以让正在执行的Go程暂停，转而执行另外的Go程
        value++
        count = value
    }
}
```
执行上面的代码，可以看到cout的值为2，因为两个Go程之间会互相修改count变量，从而相互覆盖对方的修改行为，最好导致结果错误。

>* 可以使用 go build -race 来显示不同Go程之间的竞争情况

解决方法：

* 互斥锁

另一种同步访问共享资源的方式是使用互斥锁，互斥锁这个名字来自互斥的概念。互斥锁用于在代码上创建一个临界区，保证同一时间只有一个 goroutine 可以执行这个临界代码。

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
)

var (
    counter int64
    wg      sync.WaitGroup
    mutex   sync.Mutex
)

func main() {
    wg.Add(2)

    go incCounter(1)
    go incCounter(2)

    wg.Wait()
    fmt.Println(counter)
}

func incCounter(id int) {
    defer wg.Done()

    for count := 0; count < 2; count++ {
        //同一时刻只允许一个goroutine进入这个临界区
        mutex.Lock()
        {
            value := counter
            runtime.Gosched()
            value++
            counter = value
        }
        mutex.Unlock() //释放锁，允许其他正在等待的goroutine进入临界区
    }
}
```

在Gosched()函数强制让当前goruntine退出当前线程之后，调度器会再次分配这个goruntine再次执行！从而保证同一时间只有一个goruntine可以进入有lock分配的临界区

### channel（通道）

Go语言的chan类型用于不同Go程间通信。chan类型是类型相关的，并且只能用make()函数进行创建

由于工程上有无数的线程，不同线程间有无数的数据需要处理，如果只是采用共享内存的方式处理，必须要引入锁机制才能避免出现内存写入错误。同时会使得代码变得非常臃肿。

所以Go语言采用消息机制来传递信息，这有点类似于不同进程间的处理机制。这样每个Go程就只需要处理好自己需要做好的那部分内容，然后将处理结构放回即可。

Go程间的消息传递机制就是用chan实现的！

```go
ci := make(chan int)
cs := make(chan string)
cf := make(chan interface{})
```
![在这里插入图片描述](https://img-blog.csdnimg.cn/bdbbf064c47048eeb519c73c02522476.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAbGluZ3d1X2hi,size_20,color_FFFFFF,t_70,g_se,x_16)
* 通道的类型
1）只读通道

```go
func test(a <- chan int){
	//只能从a通道中读取数据，无法发送数据
	<- a
}
```
2）只写通道

```go
func test(a chan<- int){
	//只能向a通道中写入数据，无法读取数据
	a <- 1
}
```

* 通道的运用

创建通道后，可以采用 <- 运算符进行发送和接受与通道类型一致的数据。

```go
//发送通道
ch := make(chan int)
ch <- 0//直接舍弃掉接收到的数据

//在另外一个Go程内部进行接收
data := <-ch
```

没有缓冲的通道要求必须在一个Go程中发送数据，然后在另外一个Go程中进行接受。如果只有接受或者发送通道，则通道会出现杜塞，程序会报错。

```go
package main
func main() {
    // 创建一个整型通道
    ch := make(chan int)
    // 尝试将0通过通道发送
    ch <- 0
}
```

运行代码，报错信息如下：

```go
fatal error: all goroutines are asleep - deadlock!
```

* Go语言通道的正确利用案例：

```go
package main
import (
    "fmt"
    "time"
)
func main() {
    // 构建一个通道
    ch := make(chan int)
    // 开启一个并发匿名函数
    go func() {
        // 从3循环到0
        for i := 3; i >= 0; i-- {
            // 发送3到0之间的数值
            ch <- i
            // 每次发送完时等待
            time.Sleep(time.Second)
        }
    }()//由于是匿名函数调用，此处的括号一定不要掉了
    // 遍历接收通道数据
    for data := range ch {
        // 打印通道数据
        fmt.Println(data)
        // 当遇到数据0时, 退出接收循环
        if data == 0 {
                break
        }
    }
}
```

* select 使用

下面的select怎么理解？

```go
    ch := make(chan int)
    quit := make(chan bool)
    //新开一个协程
    go func() {
        for {
            select {
            case num := <-ch:
                fmt.Println("num = ", num)
            case <-time.After(3 * time.Second):
                fmt.Println("超时")
                quit <- true
            }
        }
    }() //别忘了()
```

select 语句就是会发生堵塞，直到其中一个通信可以进行。

如果有多个分支可以进行，则随机选取一个进行执行。

```go
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```

# 包学习