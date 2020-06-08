package partdata

import (
	"log"
	"strconv"

	"github.com/ruffaustin25/ElectronicsSorting/buildconfig"
)

// PartData : template data describing a part
type PartData struct {
	Name      string
	Key       string
	Container string
	Row       int
	Column    int
	Depth     int
	URL       string
}

// Serialize : converts a part data to a string slice
func (data PartData) Serialize() []string {
	record := []string{}
	record = append(record, data.Name)
	record = append(record, data.Key)
	record = append(record, data.Container)
	record = append(record, strconv.Itoa(data.Row))
	record = append(record, strconv.Itoa(data.Column))
	record = append(record, strconv.Itoa(data.Depth))
	return record
}

// NewPartData : populates a part data based on a string slice
func NewPartData(record []string) *PartData {
	data := PartData{}
	fieldCount := 6 // Number of serialized fields
	if len(record) != fieldCount {
		log.Printf("Not enough fields in PartData record")
	}
	data.Name = record[0]
	data.Key = record[1]
	data.Container = record[2]

	rowNum, err := strconv.Atoi(record[3])
	if err != nil {
		rowNum = 0
	}
	data.Row = rowNum

	colNum, err := strconv.Atoi(record[4])
	if err != nil {
		colNum = 0
	}
	data.Column = colNum

	depthNum, err := strconv.Atoi(record[5])
	if err != nil {
		depthNum = 0
	}
	data.Depth = depthNum

	data.URL = buildconfig.BaseURL + "/part?part=" + data.Key
	return &data
}
