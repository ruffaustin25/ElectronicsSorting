package partsdatabase

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"database/sql"

	// Driver does not need import
	_ "github.com/go-sql-driver/mysql"
	"github.com/ruffaustin25/ElectronicsSorting/buildconfig"
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

	DSN := user + ":" + password + "@tcp(" + buildconfig.DatabaseURL + ")/electronics"
	parts.db, err = sql.Open("mysql", DSN)
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
func (database PartsDatabase) CreatePart(key string, name string) error {
	ctx, stop := context.WithTimeout(context.Background(), time.Second)
	defer stop()

	_, err := database.db.ExecContext(ctx, "INSERT INTO parts (`key`, `name`) VALUES ('"+key+"', '"+name+"')")
	if err != nil {
		return fmt.Errorf("Error on create part, %s", err)
	}
	return nil
}

// ArchivePart : Drops a part from the parts table (TODO: move to an archived parts table to make parts recoverable)
func (database PartsDatabase) ArchivePart(key string) {
	ctx, stop := context.WithTimeout(context.Background(), time.Second)
	defer stop()

	_, err := database.db.ExecContext(ctx, "DELETE FROM parts WHERE `key`='"+key+"'")
	if err != nil {
		log.Fatalf("Error on create part, %s", err)
	}
}

// UpdatePart :
func (database PartsDatabase) UpdatePart(part *partdata.PartData) {
	ctx, stop := context.WithTimeout(context.Background(), time.Second)
	defer stop()

	query := "UPDATE parts SET"
	if part.Name.Valid {
		query = query + " `name`='" + part.Name.String + "'"
	}
	if part.Description.Valid {
		query = query + ", `description`='" + part.Description.String + "'"
	}
	if part.Container.Valid {
		query = query + ", `container`='" + part.Container.String + "'"
	}
	if part.Row.Valid {
		query = query + ", `row`=" + strconv.FormatInt(int64(part.Row.Int32), 10)
	}
	if part.Column.Valid {
		query = query + ", `column`=" + strconv.FormatInt(int64(part.Column.Int32), 10)
	}
	if part.Depth.Valid {
		query = query + ", `depth`=" + strconv.FormatInt(int64(part.Depth.Int32), 10)
	}
	query = query + " WHERE `key`='" + part.Key.String + "'"

	_, err := database.db.ExecContext(ctx, query)
	if err != nil {
		log.Fatalf("Error on update part, %s", err)
	}
}
