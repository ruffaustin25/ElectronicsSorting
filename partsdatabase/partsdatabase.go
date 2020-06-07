package partsdatabase

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"

	"github.com/ruffaustin25/ElectronicsSorting/partdata"
)

// PartsDatabase :
type PartsDatabase struct {
	Records [][]string
}

const dbFilePath string = "./partsDatabase/parts.csv"

// NewPartsDatabase :
func NewPartsDatabase() *PartsDatabase {
	file, err := os.Open(dbFilePath)
	if err != nil {
		log.Fatal(err)
	}

	db := PartsDatabase{}

	reader := csv.NewReader(bufio.NewReader(file))
	dbRecords, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	if len(dbRecords) == 0 {
		log.Fatal("No lines in csv database")
	}
	db.Records = dbRecords[1:len(dbRecords)]
	return &db
}

// GetPartsList : Gets a list of parts in the database
func (db PartsDatabase) GetPartsList() []partdata.PartData {
	partsList := []partdata.PartData{}

	for i := 0; i < len(db.Records); i++ {
		part := partdata.NewPartData(db.Records[i])
		partsList = append(partsList, *part)
	}
	return partsList
}
