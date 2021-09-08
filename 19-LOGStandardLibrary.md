# LOG Standard Library

Go语言内置的`log包`实现了简单的日志服务。本文介绍了标准库`log`的基本使用。

## 使用 Logger

`log包`定义了`Logger类型`，该类型提供了一些格式化输出的方法。本包也提供了一个预定义的`“标准”logger`，可以通过调用函数`Print系列(Print|Printf|Println）`、`Fatal系列（Fatal|Fatalf|Fatalln）`、和`Panic系列（Panic|Panicf|Panicln）`来使用，比自行创建一个logger对象更容易使用。

例如，我们可以像下面的代码一样直接通过`log包`来调用上面提到的方法，默认它们会将日志信息打印到终端界面：

```go
package main

import "log"

func main() {
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发Fatal的日志。")
	//log.Panicln("这是一条会触发Panic的日志。")
}
```

编译并执行上面的代码会得到如下输出：

```cgo
2021/09/07 18:48:57 这是一条很普通的日志。
2021/09/07 18:48:57 这是一条很普通的日志。
2021/09/07 18:48:57 这是一条会触发fatal的日志。
```

`logger`会打印每条日志信息的日期、时间，默认输出到系统的标准错误。`Fatal`系列函数会在写入日志信息后调用`os.Exit(1)`。`Panic`系列函数会在写入日志信息后`panic`。

## 配置 logger

### 标准 logger 的配置

默认情况下的`logger`只会提供日志的时间信息，但是很多情况下我们希望得到更多信息，比如记录该日志的文件名和行号等。`log标准库`中为我们提供了定制这些设置的方法。

`log标准库`中的`Flags函数`会返回`标准logger`的输出配置，而`SetFlags函数`用来设置`标准logger`的输出配置。

```cgo
func Flags() int
func SetFlags(flag int)
```

### flag 选项

`log标准库`提供了如下的flag选项，它们是一系列定义好的常量。

```cgo
const (
    // 控制输出日志信息的细节，不能控制输出的顺序和格式。
    // 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
    Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
    LUTC                          // 使用UTC时间
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
)
```

下面我们在记录日志之前先设置一下标准logger的输出选项如下：

```go
package main

import "log"

func main() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
}
```

编译执行后得到的输出结果如下：

```cgo
2021/09/07 18:54:40.294602 
D:/GoProjects/Ditto/github.com/OrangeCH3/GoLangLab-Demo/github.com/OrangeCH3/learngo/3-senior/04-log/main.go:21: 
这是一条很普通的日志。
```

### 配置日志前缀

`log标准库`中还提供了关于日志信息前缀的两个方法：

```cgo
func Prefix() string
func SetPrefix(prefix string)
```

其中`Prefix函数`用来查看`标准logger`的输出前缀，`SetPrefix函数`用来设置输出前缀。

```go
package main

import "log"

func main() {
	log.SetFlags(log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[OrangeCH3]")
	log.Println("这是一条很普通的日志。")
}
```

上面的代码输出如下：

```cgo
[OrangeCH3]2021/09/07 18:58:28.212577 这是一条很普通的日志。
```

这样我们就能够在代码中为我们的日志信息添加指定的前缀，方便之后对日志信息进行检索和处理。

### 配置日志输出位置

```cgo
func SetOutput(w io.Writer)
```

`SetOutput函数`用来设置`标准logger`的输出目的地，默认是标准错误输出。

例如，下面的代码会把日志输出到同目录下的`log.log文件`中。

```go
package main

import (
	"fmt"
	"log"
	"os"
)
func main() {
	logFile, err := os.OpenFile("./github.com/OrangeCH3/learngo/3-senior/04-log/log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[OrangeCH3]")
	log.Println("这是一条很普通的日志。")
}
```

如果你要使用`标准的logger`，我们通常会把上面的配置操作写到`init函数`中。

```go
package main

import (
	"fmt"
	"log"
	"os"
)
func init() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
```

## 创建 logger

`log标准库`中还提供了一个创建`新logger对象`的构造函数–`New`，支持我们创建自己的logger示例。`New函数`的签名如下：

```cgo
func New(out io.Writer, prefix string, flag int) *Logger
```

`New`创建一个`Logger对象`。其中，参数out设置日志信息写入的目的地。参数prefix会添加到生成的每一条日志前面。参数flag定义日志的属性（时间、文件等等）。

举个例子：

```go
package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "[New-OrangeCH3]", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义的logger记录的日志。")
}
```

将上面的代码编译执行之后，得到结果如下：

```cgo
[New-OrangeCH3]2021/09/07 19:10:29 main.go:47: 这是自定义的logger记录的日志。
```

## 总结

Go内置的`log库`功能有限，例如无法满足记录不同级别日志的情况，我们在实际的项目中根据自己的需要选择使用第三方的日志库，如logrus、zap等。

logrus: [logrus](https://github.com/sirupsen/logrus)

zap: [zap](https://github.com/uber-go/zap)
