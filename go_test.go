/*
 * @Descripttion:
 * @version:
 * @Author: xiaoshuyui
 * @email: guchengxi1994@qq.com
 * @Date: 2021-11-12 18:57:31
 * @LastEditors: xiaoshuyui
 * @LastEditTime: 2021-11-13 08:26:57
 */
package gopathlib_test

import (
	"fmt"
	// lib "github.com/guchengxi1994/go_pathlib"
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

func Test_split_0(t *testing.T) {
	s := "c:\\docs\\a.txt"
	fmt.Println(s[:0])
}

// func Test_spilt(t *testing.T) {
// 	s := "c:\\docs\\a.txt"
// 	s2 := "c:\\docs\\a.tx\\t"
// 	s3 := "a.txt"
// 	s4 := "a..txt"
// 	fmt.Println(lib.Split(s))
// 	fmt.Println(lib.SplitExt(s))
// 	fmt.Println(lib.Split(s2))
// 	fmt.Println(lib.SplitExt(s2))

// 	fmt.Println("==============================")
// 	fmt.Println(lib.Split(s3))
// 	fmt.Println(lib.SplitExt(s3))
// 	fmt.Println(lib.Split(s4))
// 	fmt.Println(lib.SplitExt(s4))
// }
