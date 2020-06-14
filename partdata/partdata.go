package partdata

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/ruffaustin25/ElectronicsSorting/buildconfig"
)

// PartData : template data describing a part
type PartData struct {
	Key         sql.NullString
	Name        sql.NullString
	Description sql.NullString
	Container   sql.NullString
	Row         sql.NullInt32
	Column      sql.NullInt32
	Depth       sql.NullInt32
}

// NewPartData : populates a part data based on a string slice
func NewPartData(record []string) *PartData {
	data := PartData{}
	requiredFieldCount := 2 // Number of required serialized fields
	if len(record) < requiredFieldCount {
		log.Printf("Not enough fields in PartData record")
	}
	data.Key = sql.NullString{String: record[0], Valid: true}
	data.Name = sql.NullString{String: record[1], Valid: true}

	if len(record) < 3 {
		return &data
	}
	data.Container = sql.NullString{String: record[2], Valid: true}

	if len(record) < 4 {
		return &data
	}
	rowNum, err := strconv.Atoi(record[3])
	if err != nil {
		rowNum = 0
	}
	data.Row = sql.NullInt32{Int32: int32(rowNum), Valid: true}

	if len(record) < 5 {
		return &data
	}
	colNum, err := strconv.Atoi(record[4])
	if err != nil {
		colNum = 0
	}
	data.Column = sql.NullInt32{Int32: int32(colNum), Valid: true}

	if len(record) < 6 {
		return &data
	}
	depthNum, err := strconv.Atoi(record[5])
	if err != nil {
		depthNum = 0
	}
	data.Depth = sql.NullInt32{Int32: int32(depthNum), Valid: true}

	return &data
}

func FromDatabaseRow(rows *sql.Rows) *PartData {
	data := PartData{}
	err := rows.Scan(&data.Key, &data.Name, &data.Description, &data.Container, &data.Row, &data.Column, &data.Depth)
	if err != nil {
		return nil
	}
	return &data
}

func (data PartData) GetURL() string {
	return buildconfig.BaseURL + "/part?part=" + data.Key.String
}
