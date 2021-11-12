/*
 * @Descripttion:
 * @version:
 * @Author: xiaoshuyui
 * @email: guchengxi1994@qq.com
 * @Date: 2021-11-12 18:57:31
 * @LastEditors: xiaoshuyui
 * @LastEditTime: 2021-11-12 20:06:47
 */
package goPathlib

var sepLinux = CustomString{
	S: "/",
}
var sepWindows = CustomString{
	S: "\\",
}
var extsep = "."

// find the last pattern in the path s
func rFind(s, pattern string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == pattern {
			return i
		}
	}
	return -1
}

// remove the pattern on the right of a string
func rSplit(s, pattern string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == pattern {
			continue
		} else {
			return s[:i]
		}
	}
	return ""
}
