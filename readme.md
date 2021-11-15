<!--
 * @Descripttion: 
 * @version: 
 * @Author: xiaoshuyui
 * @email: guchengxi1994@qq.com
 * @Date: 2021-11-12 18:57:31
 * @LastEditors: xiaoshuyui
 * @LastEditTime: 2021-11-12 18:58:30
-->
# pathlib in go

## golang

```go
import (
    "fmt"
    "path"
)

func main() {
    s := "c:\\docs\\a.txt"
    ext := path.Ext(s)
    fmt.Println(ext)

    s2 := "c:\\docs\\a.tx\\t"
    ext = path.Ext(s2)
    fmt.Println(ext)
}

results:

    .txt
    .tx\t   
```

## python

```python
import os

s = "c:\\docs\\a.txt"
s2 = "c:\\docs\\a.tx\\t"

print(os.path.splitext(s))
print(os.path.splitext(s2))

results:

    ('c:\\docs\\a', '.txt')
    ('c:\\docs\\a.tx\\t', '')

```

## I think python's results are better.

# How to use

## code:

```go

package main

import (
	"fmt"

	lib "github.com/guchengxi1994/go_pathlib/src"
)

func main() {
	s := "/home\\xiaoshuyui\\ax.txt"
    // get extension
	fmt.Printf("lib.SplitExt(s): %v\n", lib.SplitExt(s))
    // get dirname
	fmt.Printf("lib.SplitExt(s): %v\n", lib.Dirname(s))
    // ...
    // Split get head and tail
    // SplitDrive get drive
    // Basename get basename
    // IsAbs test if a path is absolute
}

```