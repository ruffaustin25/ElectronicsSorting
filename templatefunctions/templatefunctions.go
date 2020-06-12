package templatefunctions

import (
	htmlTemplate "html/template"
	"strconv"
	textTemplate "text/template"

	"github.com/ruffaustin25/ElectronicsSorting/buildconfig"
)

// GetHTMLFuncMap : get the library of common template functions
func GetHTMLFuncMap() htmlTemplate.FuncMap {
	funcMap := htmlTemplate.FuncMap{
		"isEven":         isEven,
		"formatPosition": formatPosition,
		"formatURL":      formatURL,
	}
	return funcMap
}

// GetTextFuncMap : get the library of common template functions
func GetTextFuncMap() textTemplate.FuncMap {
	funcMap := textTemplate.FuncMap{
		"isEven":         isEven,
		"formatPosition": formatPosition,
		"formatURL":      formatURL,
	}
	return funcMap
}

func isEven(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}

func formatPosition(row int32, column int32, depth int32) string {
	return string('A'-1+row) + strconv.Itoa(int(column)) + string('a'-1+depth)
}

func formatURL(key string) string {
	return buildconfig.BaseURL + "/part?part=" + key
}
