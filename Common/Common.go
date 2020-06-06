package common

import "html/template"

// PartData : template data describing a part
type PartData struct {
	Name      string
	Container string
}

// GetCommonFuncMap : get the library of common template functions
func GetCommonFuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		"IsEven": IsEven,
	}
	return funcMap
}

func IsEven(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}
