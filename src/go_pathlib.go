package goPathlib

func SplitExt(p string) PathModel {
	var sepIndex int
	pathModel := PathModel{}
	dotIndex := rFind(p, extsep)
	if -1 == dotIndex {
		pathModel.Root = p
		pathModel.Ext = ""
		return pathModel
	}
	sepIndexLinux := rFind(p, sepLinux)
	sepIndexWindows := rFind(p, sepWindows)

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
			}
		}
	}
	pathModel.Root = p
	pathModel.Ext = ""
	return pathModel
}
