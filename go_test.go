package gopathlib_test

import (
	"fmt"
	"path"
	"testing"
)

func Test_split_ext(t *testing.T) {
	s := "c:\\docs\\a.txt"
	ext := path.Ext(s)
	fmt.Println(ext)

	s2 := "c:\\docs\\a.tx\\t"
	ext = path.Ext(s2)
	fmt.Println(ext)
}
