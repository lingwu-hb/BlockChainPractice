## HTML基础知识学习

学习重点为：常用的HTML标签、理解什么是DOM树。

![image-20211016145019346](C:/Users/lingwu/AppData/Roaming/Typora/typora-user-images/image-20211016145019346.png)

### 注释

HTML注释语法：

```html
<!--注释文字-->
```

### 标签

html标签对字母大小写不敏感。但是还是建议使用小写字母。

在学习标签时，需要注意两个方向的内容：**标签的用途、标签在浏览器中的默认样式。**

所谓的**语义化**，就是指明白每个标签的用途（在什么情况下使用此标签更合理）

#### 常用的单闭合标签

![](../Image/单闭合标签.png)

#### 常见标签

* \<p\>段落

默认格式会在前面一行和后面一行出现空行，如果需要修改，可以通过css样式进行调整。

* \<span\>

用于设置单独的样式使用。（自定义）

```html
<p>我的母亲有 <span style="color:blue">蓝色</span> 的眼睛。</p>
```

* \<header\>

用于书写页面头部

* \<footer\>

用于书写底部部分

* \<section\>

用来定义一个区域

* \<aside>

代表侧边栏区域，但是作用等同于div，只是具备了语义化

* \<ul\> 和 \<\li> 

添加新闻信息列表

```html
<ul>
    <li>列表一</li>
    <li>列表二</li>
    <!--默认的情况下，每一项li前面都会自带一个圆点。-->
</ul>
```

* \<img>

![image-20211016201503569](C:/Users/lingwu/AppData/Roaming/Typora/typora-user-images/image-20211016201503569.png)

* \<a>
  * 表示超链接
  * target属性

![image-20211016201631412](C:/Users/lingwu/AppData/Roaming/Typora/typora-user-images/image-20211016201631412.png)

* 表格\<table>

* \<form>

![image-20211016202442846](C:/Users/lingwu/AppData/Roaming/Typora/typora-user-images/image-20211016202442846.png)

## CSS基础知识

css全称为：层叠样式表（Cascading Style Sheets)。它主要是用于定义HTML内容在浏览器内的显示样式，如文字大小、颜色、字体加粗等。

基本用法如下：

```html
p {
        font-size: 20px;
        /*设置文字字号*/
        color: red;
        /*设置文字颜色*/
        font-weight: bold;
        /*设置字体加粗*/
}
```

css样式由**选择符**和**声明**组成，而声明又是由属性和值组成。

**选择符：**又称选择器，指明网页中要应用样式规则的元素，如本例中是网页中所有的段（p）的文字将变成蓝色，而其他的元素（如ol）不会受到影响。

### css样式种类

css样式分为三种：内联式、嵌入式和外部式。

* 内联式

css样式代码要写在元素的开始标签里面，并且css样式代码要写在style=""双引号中，如果有多条css样式代码设置可以写在一起，中间用分号隔开。

```html
<p style = "color:red">
    这里文字是红色的。
</p>
```

* 嵌入式

css代码必须写在\<style>和\</style>之间。

```html
<style type = "text/css">
    span{
        color:red;
    }
</style>
```

* 外部式

外部式就是把css代码写在一个单独的外部文件中，这个css文件可以以```.css```为扩展名，在\<head>内使用\<link>标签可以将css样式文件连接到HTML文件内。

至于具体是什么样式，需要打开.css文件才可以知道。

```html
<!--此处的rel="stylesheet" type="text/css"是固定语法，不可以修改-->
<link href="style.css" rel="stylesheet" type="text/css" />
```

### css选择器

* 标签选择器

一般情况，标签选择器就是html中的标签。

* 类选择器

在html中运用得非常多。

语法：

```html
.类选择器名称{css样式代码}
```

使用时，直接在html标签中加一条属性calss = "类选择器名称"即可。例如：```<span class = "stress"> 勇气 </span>```

类选择器适用于会在文档中经常用到的样式。

* **ID选择器**

语法：

```html
<style type="text/css">
    #stress {
        color: red;
    }
    #green {
        color: green;
    }
</style>
```

* 类选择器和ID选择器的区别

前者可以在一个文档中**使用多次**，但是后者只能使用一次

前者可以采用类列表的方式，为同一个元素同时作用多个类选择器描述的样式，但是后者不可以。

* 子选择器与后代选择器

子选择器语法：

下面的代码表示：将名为food的class里面的一个子标签li里面的元素设置为某个样式。

```html
.food>li{border:1px solid red;}
```

后代选择器

```html
.food li{border: 1px solid red;}
```

总结：**>**作用于元素的第一代后代，**空格**作用于元素的所有后代。

* 通配符选择器

直接用一个(*)即可将html中所有的元素全部选中。

```html
<style type = "text/css">
    * {
        color: red;
        font-size: 20px;
    }
</style>
```

* 伪类选择器

下面代码表示：当鼠标滑过a标签指向的元素的地方时，元素颜色变为红色。

```html
a:hover {
        color: red;
    }
```

* 选择器的优先级顺序

内联样式 > id选择器 > 类选择器 > 标签选择器 > 通配符选择器

* 权值判断

当同一个元素被施加多种样式时，根据权值高进行调整。

计算权值时，根据不同类型和距离远近同时进行计算。

**标签的权值为1，类选择符的权值为10，ID选择符的权值最高为100。**

```html
p{color:red;} /*权值为1*/
p span{color:green;} /*权值为1+1=2*/
.warning{color:white;} /*权值为10*/
p span.warning{color:purple;} /*权值为1+1+10=12*/
#footer .note p{color:yellow;} /*权值为100+10+1=111*/
```

### CSS文本样式

* text-decoration可以设置添加到文本的修饰。

text-decoration默认值为none, 定义标准的文本。

text-decoration的值为underline为定义文本下的一条线。

text-decoration的值为overline为定义文本上的一条线。

text-decoration的值为line-through为定义穿过文本下的一条线，一般用于商品折扣价。

* text-align用于设置文本内容位置

text-align可以设置为left、center、right。

还有更多的东西，可以在网上查询。

### 元素分类

在CSS中，html中的标签元素大体被分为三种不同的类型：**块状元素**、**内联元素(又叫行内元素)**和**内联块状元素**。

* 常见的块状元素有：

```
<div>、<p>、<h1>、<ol>、<ul>、<dl>、<table>、<address>、<blockquote>、<form>
```

**块状元素的特点**：

1、每个块级元素都从新的一行开始，并且其后的元素也另起一行。（真霸道，一个块级元素独占一行）

2、元素的高度、宽度、行高以及顶和底边距都可设置。

3、元素宽度在不设置的情况下，是它本身父容器的100%（和父元素的宽度一致），除非设定一个宽度。

* 常见的内联元素

```
<a>、<span>、<br>、<i>、<em>、<strong>、<label>、<q>、<var>、<cite>、<code>
```

**内联元素的特点**：

1、和其他元素都在一行上；

2、元素的高度、宽度及顶部和底部边距**不可**设置；

3、元素的宽度就是它包含的文字或图片的宽度，不可改变。

* 常见的内联块状元素

```
<img>、<input>
```

由于内联元素无法设置元素的高度、宽度等，所以就有一种内联块状元素，同时具有内联元素和块状元素的特点。

**inline-block 元素特点：**

1、和其他元素都在一行上；

2、元素的高度、宽度、行高以及顶和底边距都可设置。

以上几种元素属性都可以通过设置元素的diaplay的属性值来改变。例如：

```html
p{
	display: inline-block;
}
```

### css盒子模型

![](../Image/css盒子模型.jpg)

* 代码示例：

```html
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>宽度和高度</title>
    <style type="text/css">
    li {
        border:1px solid red;
        margin:10px;
        border-bottom: 1px dotted #ccc;
        width:200px;
        height:30px;
        background-color:yellow;
    }
    </style>
</head>
<body>
    <ul>
        <li>别让不会说话害了你</li>
        <li>二十七八岁就应该有的见识</li>
        <li>别让不好意思害了你</li>
    </ul>
</body>
</html>
```

效果如下：

![](../Image/效果图.png)

* 边框border

盒子模型的边框就是围绕着内容及补白的线，这条线你可以设置它的粗细、样式和颜色(边框三个属性)。

```html
p{
	border: 1px, solid, red;/*缩写形式*/
}
/*全写形式*/
div{
    border-width:2px;
    border-style:solid;
    border-color:red;
}
```

可以单独设置下边沿

```html
div{border-bottom:1px solid red;}
```

### css布局模型

清楚了CSS3 盒模型的基本概念、 盒模型类型， 我们就可以深入探讨网页布局的基本模型了。布局模型与盒模型一样都是 CSS3 最基本、 最核心的概念。 但布局模型是建立在盒模型基础之上，又不同于我们常说的 CSS3 布局样式或 CSS3 布局模板。如果说布局模型是本，那么 CSS3 布局模板就是末了，是外在的表现形式。 
CSS3包含3种基本的布局模型，用英文概括为：```Flow、Layer 和 Float```。
在网页中，元素有三种布局模型：
1、流动模型（Flow）
2、浮动模型 (Float)
3、层模型（Layer）

* 流动模型

流动（Flow）是默认的网页布局模式。也就是说网页在默认状态下的 HTML 网页元素都是根据流动模型来分布网页内容的。

**流动布局模型具有2个比较典型的特征：**

第一点，**块状元素**都会在所处的**包含元素内**自上而下按顺序垂直延伸分布，因为在默认状态下，块状元素的宽度都为**100%**。实际上，块状元素都会以行的形式占据位置。

第二点，在流动模型下，**内联元素**都会在所处的包含元素内从左到右水平分布显示。（内联元素可不像块状元素这么霸道独占一行）

* 浮动模型

在css样式表中加入```float```即可将元素定义以浮动状态呈现。

```html
<style>
    div{
        width: 200px;
        height: 100px;
        border: 20px dotted #333;
    }
    #div1{
        float: left;/*相当于设置为向左对齐*/
    }
</style>
```

* 层布局模型

所谓的层布局模型，就是将html元素进行划分为不同的图层，然后再不同的图层上分别进行操作。其原理类似于photoShop的图层操作。CSS定义了一组定位（positioning）属性来支持层布局模型。

层模型有三种形式：

1、**绝对定位**(position: absolute)

设置了position: absolute之后，这条语句表示将元素从文档流中拖出来，然后可以使用left、right、top、bottom进行相对于其最接近的一个具有定位属性的父包含块进行绝对定位。如果不存在这样的包含块，则相对于body元素，即相对于**浏览器窗口**。

绝对定位会改变元素的性质，快的宽高被内容撑开。

绝对定位会使元素提升一个层级，同时，绝对定位是相对于其**包含块**进行定位的

> * 什么是包含块（containing block）？
>
> -- 正常情况下，包含块就是离当前元素最近的祖先块元素
>
> -- 绝对定位下的包含块：离它最近的开启了定位的祖先元素。如果所有的祖先元素都没有开启定位则根元素就是它的包含块。
>
> html（根元素、初始包含块）

2、**相对定位**(position: relative)

设置了position: relative之后，可以类似于绝对定位一样将元素相对于其原来的位置进行移动。不过，**元素原来的位置并没有清空**，其他元素仍然不能占据该元素的位置。

3、**固定定位**(position: fixed)

与绝对定位类似，只不过，固定定位是**相对于文档视图而言的**，不会随着文档流的流动而改变位置，除非浏览器显示发生变化。

可以参考一些小网站上的小广告，无论如何移动，广告在浏览器中的位置一直不会改变。

### Flex弹性模型

* Flex基本概念

![](../../学习图片/Flex基本概念1.png)

在 flex 容器中默认存在两条轴，水平主轴(main axis) 和垂直的交叉轴(cross axis)，这是默认的设置。

在容器中的每个单元块被称之为 flex item，每个项目占据的主轴空间为 (main size), 占据的交叉轴的空间为 (cross size)。

这里需要强调，不能先入为主认为宽度就是 main size，高度就是 cross size，这个还要取决于你主轴的方向，如果你垂直方向是主轴，那么项目的高度就是 main size。

* Flex容器

实现flex布局首先需要先指定一个容器，任何一个容器都可以被指定为flex布局，这样容器内部的元素就可以使用flex来进行布局。

```html
.container{
	display: flex | inline-flex;
}
```

分别生成一个块状或行内的 flex 容器盒子。简单说来，如果你使用块元素如 div，你就可以使用 flex，而如果你使用行内元素，你可以使用 inline-flex。

有六种属性可以设置在容器上，分别为：

1. flex-direction
2. flex-wrap
3. flex-flow
4. justify-content
5. align-items
6. align-content

## VSCode常用快捷键

文本编辑

* 行复制：ctrl + c。
* 行移动：ctrl + 向上/向下
* 行复制：shift + Alt + 向上/向下
* 代码缩进：ctrl + [ / ] 即可以完成整行的快速缩进

注释

* 行注释：ctrl + /
* **块注释**：alt + shift + A
* 格式化：alt + shift + F

