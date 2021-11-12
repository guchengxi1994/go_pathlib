# pathlib in go

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