package gopathlib_test

import (
	"fmt"
	goPathlib "goPathlib/src"
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

	m := goPathlib.SplitExt(s)
	fmt.Printf("m: %v\n", m)
	m = goPathlib.SplitExt(s2)
	fmt.Printf("m: %v\n", m)

}
