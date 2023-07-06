package containerdata

import (
	"database/sql"
	"strconv"

	"github.com/ruffaustin25/ElectronicsSorting/buildconfig"
)

// ContainerData : template data describing a container
type ContainerData struct {
	Key         sql.NullString
	Name        sql.NullString
	Description sql.NullString
	Height      sql.NullInt32
	Width       sql.NullInt32
	Depth       sql.NullInt32
}

// FromMap :
func FromMap(record map[string]string) *ContainerData {
	data := ContainerData{}

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

	heightStr, found := record["height"]
	if found {
		height, err := strconv.Atoi(heightStr)
		if err == nil {
			data.Height = sql.NullInt32{Int32: int32(height), Valid: true}
		}
	}

	widthStr, found := record["width"]
	if found {
		width, err := strconv.Atoi(widthStr)
		if err == nil {
			data.Width = sql.NullInt32{Int32: int32(width), Valid: true}
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
func FromDatabaseRow(rows *sql.Rows) *ContainerData {
	data := ContainerData{}
	err := rows.Scan(&data.Key, &data.Name, &data.Description, &data.Height, &data.Width, &data.Depth)
	if err != nil {
		return nil
	}
	return &data
}

// GetURL :
func (data ContainerData) GetURL() string {
	return "http://" + buildconfig.BaseURL + ":" + strconv.Itoa(buildconfig.AppPort) + "/container?container=" + data.Key.String
}
