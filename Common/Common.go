package common

import "html/template"

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
