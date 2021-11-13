/*
 * @Descripttion:
 * @version:
 * @Author: xiaoshuyui
 * @email: guchengxi1994@qq.com
 * @Date: 2021-11-12 18:57:31
 * @LastEditors: xiaoshuyui
 * @LastEditTime: 2021-11-12 20:58:05
 */
package goPathlib

type BaseModel interface {
	ToStrings() (string, string)
}

// get the path root and extension
type PathExtModel struct {
	Root string
	Ext  string
}

func (m *PathExtModel) ToStrings() (string, string) {
	return m.Root, m.Ext
}

// get the path head and tail
type PathModel struct {
	Head string
	Tail string
}

func (m *PathModel) ToStrings() (string, string) {
	return m.Head, m.Tail
}

// custom string
type CustomString struct {
	S string
}

func (c *CustomString) Multiple(i int) string {
	var s = ""
	for j := 0; j < i; j++ {
		s = s + c.S
	}
	return s
}
