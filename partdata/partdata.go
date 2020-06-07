package partdata

import (
	"log"
	"reflect"
)

// PartData : template data describing a part
type PartData struct {
	Name      string
	URLName   string
	URL       string
	Container string
}

// Serialize : converts a part data to a string slice
func (data PartData) Serialize() []string {
	record := []string{}
	record = append(record, data.Name)
	record = append(record, data.URLName)
	record = append(record, data.URL)
	record = append(record, data.Container)
	return record
}

// NewPartData : populates a part data based on a string slice
func NewPartData(record []string) *PartData {
	data := PartData{}
	fieldCount := reflect.TypeOf(data).NumField()
	if len(record) != fieldCount {
		log.Printf("Not enough fields in PartData record")
	}
	data.Name = record[0]
	data.URLName = record[1]
	data.URL = record[2]
	data.Container = record[3]
	return &data
}
