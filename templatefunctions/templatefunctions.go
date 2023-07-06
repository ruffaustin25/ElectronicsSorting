package templatefunctions

import (
	"strconv"

	htmlTemplate "html/template"
	textTemplate "text/template"

	"github.com/ruffaustin25/ElectronicsSorting/buildconfig"
)

// GetHTMLFuncMap : get the library of common template functions
func GetHTMLFuncMap() htmlTemplate.FuncMap {
	funcMap := htmlTemplate.FuncMap{
		"isEven":             isEven,
		"getRowLetter":       getRowLetter,
		"getDepthLetter":     getDepthLetter,
		"formatPosition":     formatPosition,
		"formatPartURL":      formatPartURL,
		"formatContainerURL": formatContainerURL,
	}
	return funcMap
}

// GetTextFuncMap : get the library of common template functions
func GetTextFuncMap() textTemplate.FuncMap {
	funcMap := textTemplate.FuncMap{
		"isEven":             isEven,
		"getRowLetter":       getRowLetter,
		"getDepthLetter":     getDepthLetter,
		"formatPosition":     formatPosition,
		"formatPartURL":      formatPartURL,
		"formatContainerURL": formatContainerURL,
	}
	return funcMap
}

func isEven(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}

func getRowLetter(row int) string {
	return string('A' + row)
}

func getDepthLetter(depth int) string {
	return string('a' + depth)
}

func formatPosition(row int32, column int32, depth int32) string {
	return getRowLetter(int(row-1)) + strconv.Itoa(int(column)) + getDepthLetter(int(depth-1))
}

func formatPartURL(key string) string {
	return "http://" + buildconfig.BaseURL + ":" + strconv.Itoa(buildconfig.AppPort) + "/part?part=" + key
}

func formatContainerURL(key string) string {
	return "http://" + buildconfig.BaseURL + ":" + strconv.Itoa(buildconfig.AppPort) + "/container?container=" + key
}
