package partdata

import (
	"log"

	"github.com/ruffaustin25/ElectronicsSorting/buildconfig"
)

// PartData : template data describing a part
type PartData struct {
	Name      string
	Key       string
	Container string
	URL       string
}

// Serialize : converts a part data to a string slice
func (data PartData) Serialize() []string {
	record := []string{}
	record = append(record, data.Name)
	record = append(record, data.Key)
	record = append(record, data.Container)
	return record
}

// NewPartData : populates a part data based on a string slice
func NewPartData(record []string) *PartData {
	data := PartData{}
	fieldCount := 3 // Number of serialized fields
	if len(record) != fieldCount {
		log.Printf("Not enough fields in PartData record")
	}
	data.Name = record[0]
	data.Key = record[1]
	data.Container = record[2]
	data.URL = buildconfig.BaseURL + "/part?part=" + data.Key
	return &data
}
