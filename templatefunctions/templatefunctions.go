package templatefunctions

import (
	htmlTemplate "html/template"
	"strconv"
	textTemplate "text/template"
)

// GetHtmlFuncMap : get the library of common template functions
func GetHtmlFuncMap() htmlTemplate.FuncMap {
	funcMap := htmlTemplate.FuncMap{
		"IsEven":         IsEven,
		"FormatPosition": FormatPosition,
	}
	return funcMap
}

// GetTextFuncMap : get the library of common template functions
func GetTextFuncMap() textTemplate.FuncMap {
	funcMap := textTemplate.FuncMap{
		"IsEven":         IsEven,
		"FormatPosition": FormatPosition,
	}
	return funcMap
}

func IsEven(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}

func FormatPosition(row int, column int, depth int) string {
	return string('A'-1+row) + strconv.Itoa(column) + string('a'-1+depth)
}
