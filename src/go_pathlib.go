/*
 * @Descripttion:
 * @version:
 * @Author: xiaoshuyui
 * @email: guchengxi1994@qq.com
 * @Date: 2021-11-12 18:57:31
 * @LastEditors: xiaoshuyui
 * @LastEditTime: 2021-11-13 09:12:32
 */
package goPathlib

import (
	"strings"
)

func SplitExt(p string) PathExtModel {
	var sepIndex int
	pathModel := PathExtModel{}
	dotIndex := rFind(p, extsep)
	if -1 == dotIndex {
		pathModel.Root = p
		pathModel.Ext = ""
		return pathModel
	}
	sepIndexLinux := rFind(p, sepLinux.S)
	sepIndexWindows := rFind(p, sepWindows.S)

	if sepIndexLinux > sepIndexWindows {
		sepIndex = sepIndexLinux
	} else {
		sepIndex = sepIndexWindows
	}

	if dotIndex > sepIndex {
		filenameIndex := sepIndex + 1
		for {
			if filenameIndex < dotIndex {
				if p[filenameIndex:filenameIndex+1] != extsep {
					pathModel.Root = p[:dotIndex]
					pathModel.Ext = p[dotIndex:]
					return pathModel
				}
				filenameIndex += 1
			} else {
				break
			}
		}
	}
	pathModel.Root = p
	pathModel.Ext = ""
	return pathModel
}

func Split(p string) PathModel {
	var pathModel PathModel
	var sepIndex int
	sepIndexLinux := rFind(p, sepLinux.S)
	sepIndexWindows := rFind(p, sepWindows.S)
	if sepIndexLinux > sepIndexWindows {
		sepIndex = sepIndexLinux
	} else {
		sepIndex = sepIndexWindows
	}

	i := sepIndex + 1
	head, tail := p[:i], p[i:]
	pathModel.Tail = tail
	if head != "" && (head != sepWindows.Multiple(len(head)) || head != sepLinux.Multiple(len(head))) {
		head = rStrip(head, string(p[i]))
	}
	pathModel.Head = head
	return pathModel
}

func SplitDrive(p string) PathModel {
	var model PathModel
	index := lFind(p, ":")
	if index == -1 {
		model.Head = p[:0]
		model.Tail = p
	}
	model.Head = p[:index+1]
	model.Tail = p[index+1:]

	return model
}

func Dirname(p string) string {
	_p := strings.ReplaceAll(p, sepWindows.S, sepLinux.S)
	index := rFind(_p, sepLinux.S)
	if index == -1 {
		return p
	}
	index = index + 1
	head := p[:index]
	if head != sepLinux.Multiple(len(head)) {
		head = rStrip(head, sepLinux.S)
	}
	return head
}

func Basename(p string) string {
	_p := strings.ReplaceAll(p, sepWindows.S, sepLinux.S)
	index := rFind(_p, sepLinux.S)
	if index == -1 {
		return p
	}
	index = index + 1
	return p[:index]
}

// Test whether a path is absolute
func IsAbs(p string) bool {
	_p := strings.ReplaceAll(p, sepWindows.S, sepLinux.S)
	return string(_p[0]) == sepLinux.S
}

/*
Join two or more pathname components, inserting '/' as needed.
If any component is an absolute path, all previous path components
will be discarded.  An empty last part will result in a path that
ends with a separator.
*/
func Join(a string, args ...string) string {
	_a := strings.ReplaceAll(a, sepWindows.S, sepLinux.S)
	for _, v := range args {
		if string(v[0]) == sepLinux.S {
			_a = v
		} else if v == "" || string(v[len(v)-1]) == sepLinux.S {
			_a = _a + v
		} else {
			_a = _a + sepLinux.S + v
		}
	}
	return _a
}
