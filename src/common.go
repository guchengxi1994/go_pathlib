package goPathlib

var sepLinux = "/"
var sepWindows = "\\"
var extsep = "."

// find the last pattern in the path s
func rFind(s, pattern string) int {
	for i := len(s) - 1; i >= 0 && s[i] != '/'; i-- {
		if string(s[i]) == pattern {
			return i
		}
		if i == 0 {
			return 0
		}
	}
	return -1
}
