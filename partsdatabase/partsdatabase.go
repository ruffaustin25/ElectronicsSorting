package partsdatabase

import (
	"context"
	"log"
	"time"

	"database/sql"

	// Driver does not need import
	_ "github.com/go-sql-driver/mysql"
	"github.com/ruffaustin25/ElectronicsSorting/partdata"
)

// PartsDatabase :
type PartsDatabase struct {
	db *sql.DB
}

const dbFilePath string = "./partsdatabase/parts.csv"
const user string = "root"
const password string = "root"

// NewPartsDatabase :
func NewPartsDatabase() *PartsDatabase {
	var err error
	parts := PartsDatabase{}

	parts.db, err = sql.Open("mysql", user+":"+password+"@/electronics")
	if err != nil {
		log.Fatalf("Could not init SQL Database, %s", err)
	}

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	err = parts.db.PingContext(ctx)
	if err != nil {
		log.Fatalf("Could not ping SQL Database, %s", err)
	}

	return &parts
}

// GetPartsList : Gets all parts in the database
func (database PartsDatabase) GetPartsList() []partdata.PartData {
	ctx, stop := context.WithTimeout(context.Background(), time.Second)
	defer stop()

	rows, err := database.db.QueryContext(ctx, "SELECT * FROM parts")
	defer rows.Close()
	if err != nil {
		log.Fatalf("Error on get parts, %s", err)
	}

	parts := []partdata.PartData{}

	for rows.Next() {
		part := partdata.FromDatabaseRow(rows)
		if part == nil {
			break
		}
		parts = append(parts, *part)
	}
	return parts
}

// GetPart : Gets the part with the corresponding url-friendly key name
func (database PartsDatabase) GetPart(key string) *partdata.PartData {
	ctx, stop := context.WithTimeout(context.Background(), time.Second)
	defer stop()

	rows, err := database.db.QueryContext(ctx, "SELECT * FROM parts WHERE `key`='"+key+"'")
	defer rows.Close()
	if err != nil {
		log.Fatalf("Error on get parts, %s", err)
	}
	if !rows.Next() {
		return nil
	}
	return partdata.FromDatabaseRow(rows)
}

// CreatePart : Initializes a new part description
func (database PartsDatabase) CreatePart(key string, name string) {
	ctx, stop := context.WithTimeout(context.Background(), time.Second)
	defer stop()

	rows, err := database.db.QueryContext(ctx, "INSERT INTO parts (`key`, `name`) VALUES ('"+key+"', '"+name+"')")
	if rows != nil {
		rows.Close()
	}
	if err != nil {
		log.Fatalf("Error on create part, %s", err)
	}
}

// ArchivePart : Drops a part from the parts table (TODO: move to an archived parts table to make parts recoverable)
func (database PartsDatabase) ArchivePart(key string) {
	ctx, stop := context.WithTimeout(context.Background(), time.Second)
	defer stop()

	rows, err := database.db.QueryContext(ctx, "DELETE FROM parts WHERE `key`='"+key+"'")
	if rows != nil {
		rows.Close()
	}
	if err != nil {
		log.Fatalf("Error on create part, %s", err)
	}
}

// UpdatePart :
func (database PartsDatabase) UpdatePart(part *partdata.PartData) {
	ctx, stop := context.WithTimeout(context.Background(), time.Second)
	defer stop()

	rows, err := database.db.QueryContext(ctx, "UPDATE parts SET `name`='"+part.Name.String+"' WHERE `key`='"+part.Key.String+"'")
	if rows != nil {
		rows.Close()
	}
	if err != nil {
		log.Fatalf("Error on update part, %s", err)
	}
}
