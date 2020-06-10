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
func (db PartsDatabase) GetPartsList() []partdata.PartData {
	ctx, stop := context.WithTimeout(context.Background(), time.Second)
	defer stop()

	rows, err := db.db.QueryContext(ctx, "SELECT * FROM part")
	if err != nil {
		log.Fatalf("Error on get parts, %s", err)
	}

	parts := []partdata.PartData{}

	for rows.Next() {
		parts = append(parts, *partdata.FromDatabaseRow(rows))
	}
	return parts
}

// GetPart : Gets the part with the corresponding url-friendly key name
func (db PartsDatabase) GetPart(key string) *partdata.PartData {
	ctx, stop := context.WithTimeout(context.Background(), time.Second)
	defer stop()

	rows, err := db.db.QueryContext(ctx, "SELECT * FROM part WHERE `key`='"+key+"'")
	if err != nil {
		log.Fatalf("Error on get parts, %s", err)
	}
	rows.Next()
	return partdata.FromDatabaseRow(rows)
}
