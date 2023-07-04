package partdata

import (
	"database/sql"
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

// FromMap :
func FromMap(record map[string]string) *PartData {
	data := PartData{}

	key, found := record["key"]
	if found {
		data.Key = sql.NullString{String: key, Valid: true}
	}

	name, found := record["name"]
	if found {
		data.Name = sql.NullString{String: name, Valid: true}
	}

	description, found := record["description"]
	if found {
		data.Description = sql.NullString{String: description, Valid: true}
	}

	container, found := record["container"]
	if found {
		data.Container = sql.NullString{String: container, Valid: true}
	}

	rowStr, found := record["row"]
	if found {
		row, err := strconv.Atoi(rowStr)
		if err == nil {
			data.Row = sql.NullInt32{Int32: int32(row), Valid: true}
		}
	}

	colStr, found := record["column"]
	if found {
		column, err := strconv.Atoi(colStr)
		if err == nil {
			data.Column = sql.NullInt32{Int32: int32(column), Valid: true}
		}
	}

	depthStr, found := record["depth"]
	if found {
		depth, err := strconv.Atoi(depthStr)
		if err == nil {
			data.Depth = sql.NullInt32{Int32: int32(depth), Valid: true}
		}
	}

	return &data
}

// FromDatabaseRow :
func FromDatabaseRow(rows *sql.Rows) *PartData {
	data := PartData{}
	err := rows.Scan(&data.Key, &data.Name, &data.Description, &data.Container, &data.Row, &data.Column, &data.Depth)
	if err != nil {
		return nil
	}
	return &data
}

// GetURL :
func (data PartData) GetURL() string {
	return "https://" + buildconfig.BaseURL + buildconfig.AppPort + buildconfig.BaseURL + "/part?part=" + data.Key.String
}
