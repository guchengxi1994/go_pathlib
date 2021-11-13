/*
 * @Descripttion:
 * @version:
 * @Author: xiaoshuyui
 * @email: guchengxi1994@qq.com
 * @Date: 2021-11-12 18:57:31
 * @LastEditors: xiaoshuyui
 * @LastEditTime: 2021-11-12 20:58:24
 */
package goPathlib

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
		head = rSplit(head, string(p[i]))
	}
	pathModel.Head = head
	return pathModel
}
