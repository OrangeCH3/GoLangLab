# Basic Map

Go语言中提供的映射关系容器为`map`，其内部使用`散列表（hash）`实现。

map是一种无序的基于`key-value`的数据结构，Go语言中的`map`是引用类型，必须初始化才能使用。

## map定义

Go语言中 `map` 的定义语法如下：

```cgo
map[KeyType]ValueType
```

其中：

- KeyType:表示键的类型。
- ValueType:表示键对应的值的类型。

`map`类型的变量默认初始值为`nil`，需要使用`make()`函数来分配内存。语法为：

```cgo
make(map[KeyType]ValueType, [cap])
```

其中`cap`表示`map`的容量，该参数虽然不是必须的，但是我们应该在初始化`map`的时候就为其指定一个合适的容量。

## map基本使用

`map`中的数据都是成对出现的，`map`的基本使用示例代码如下：

```go
package main

import "fmt"

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)
}
```

输出：

```cgo
map[小明:100 张三:90]
100
type of a:map[string]int
```

`map`也支持在声明的时候填充元素，例如：

```go
package main

import "fmt"

func main() {
	userInfo := map[string]string{
		"username": "沙河小王子",
		"password": "123456",
	}
	fmt.Println(userInfo)
}
```

## 判断某个键是否存在














