## 说明
fixedlist 提供定长列表，初始化时定义好长度，add()负责向尾部添加数据,当数据达到指定长度时，fixedlist会自动删除头部数据。

该场景常用于统计最近X时间的数据，比如最近60秒内每秒的请求量，只需把每秒的统计数据顺序add即可。
## 安装
```
go get github.com/go-basic/fixedlist
```
## 使用
```
package main

import (
	"fmt"
	"github.com/go-basic/fixedlist"
)

func main() {
	f := fixedlist.NewFixedList(2)
	f.Add("a")
	f.Add("b")
	fmt.Println(f.Data())
	f.Add("c")
	fmt.Println(f.Data())
	f.Add("d")
	fmt.Println(f.Data())
	f.Add("e")
	fmt.Println(f.Data())
}
/**
[a b]
[b c]
[c d]
[d e]
 */
```