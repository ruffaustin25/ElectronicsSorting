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
	Parts []partdata.PartData
}

const dbFilePath string = "./partsdatabase/parts.csv"

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

	db.Parts = []partdata.PartData{}
	for i := 1; i < len(dbRecords); i++ {
		part := partdata.NewPartData(dbRecords[i])
		db.Parts = append(db.Parts, *part)
	}
	return &db
}

// GetPart : Gets the part with the corresponding url-friendly key name
func (db PartsDatabase) GetPart(key string) *partdata.PartData {
	for _, part := range db.Parts {
		if part.Key == key {
			return &part
		}
	}
	return nil
}
