# fooutils
[![Build Status](https://travis-ci.org/importcjj/fooutils.svg?branch=master)](https://travis-ci.org/importcjj/fooutils)

乱七八糟的工具包

## 安装
```go get github.com/importcjj/fooutils```

#### 1.检测身份证号码是否有效

```go
package main

import "github.com/importcjj/fooutils"
import "fmt"

func main() {
     fmt.Println(fooutils.IsValidIDCode("身份证号码")
}
```
