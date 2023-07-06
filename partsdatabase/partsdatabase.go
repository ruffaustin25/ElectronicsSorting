package partsdatabase

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ruffaustin25/ElectronicsSorting/containerdata"
	"github.com/ruffaustin25/ElectronicsSorting/partdata"
)

// PartsDatabase :
type PartsDatabase struct {
	db *sql.DB
}

const dbFilePath string = "./partsdatabase/parts.csv"
const user string = "root"
const password string = "root"
const maxConnectRetries = 100

// NewPartsDatabase :
func NewPartsDatabase() *PartsDatabase {
	var err error
	retryCount := 0
	parts := PartsDatabase{}

	for retryCount < maxConnectRetries {
		parts.db, err = sql.Open("sqlite3", "./sqlite.db")

		if err != nil {
			log.Printf("Could not init SQL Database, %s. Retrying in 10 seconds", err)
			retryCount++
			time.Sleep(10 * time.Second)
			continue
		}

		ctx, stop := context.WithCancel(context.Background())
		defer stop()

		err = parts.db.PingContext(ctx)
		if err != nil {
			log.Printf("Could not ping SQL Database, %s. Retrying in 10 seconds", err)
			retryCount++
			time.Sleep(10 * time.Second)
			continue
		}

		break
	}

	_, err = parts.db.Exec("CREATE TABLE IF NOT EXISTS containers (key CHAR(64), name CHAR(64), description VARCHAR(1024), height INT, width INT, depth INT, PRIMARY KEY (key))")
	if err != nil {
		log.Fatalf("Could not create default containers table, err: %s", err)
	}

	_, err = parts.db.Exec("CREATE TABLE IF NOT EXISTS parts (key CHAR(64), name CHAR(64), description VARCHAR(1024), container CHAR(32), row INT, column INT, depth INT, PRIMARY KEY (key), FOREIGN KEY (container) REFERENCES containers(container))")
	if err != nil {
		log.Fatalf("Could not create default parts table, err: %s", err)
	}

	if retryCount >= maxConnectRetries {
		log.Fatalf("Could not connect to SQL database after %d retries, stopping program", retryCount)
	} else {
		log.Printf("Successfully connected to SQL database!")
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
		log.Fatalf("Error on archive part, %s", err)
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

func (database PartsDatabase) GetContainersList() []containerdata.ContainerData {
	ctx, stop := context.WithTimeout(context.Background(), time.Second)
	defer stop()

	rows, err := database.db.QueryContext(ctx, "SELECT * FROM containers")
	defer rows.Close()
	if err != nil {
		log.Fatalf("Error on get containers, %s", err)
	}

	containers := []containerdata.ContainerData{}

	for rows.Next() {
		container := containerdata.FromDatabaseRow(rows)
		if container == nil {
			break
		}
		containers = append(containers, *container)
	}
	return containers
}

// GetContainer : Gets the container with the corresponding url-friendly key name
func (database PartsDatabase) GetContainer(key string) *containerdata.ContainerData {
	ctx, stop := context.WithTimeout(context.Background(), time.Second)
	defer stop()

	rows, err := database.db.QueryContext(ctx, "SELECT * FROM containers WHERE `key`='"+key+"'")
	defer rows.Close()
	if err != nil {
		log.Fatalf("Error on get container, %s", err)
	}
	if !rows.Next() {
		return nil
	}
	return containerdata.FromDatabaseRow(rows)
}

// GetPartsInContainer : Gets a 3d array of all parts with this one set as their container, the ordering will be depth->height->width
func (database PartsDatabase) GetPartsInContainer(container *containerdata.ContainerData) [][][]*partdata.PartData {
	if !container.Height.Valid || !container.Width.Valid || !container.Depth.Valid {
		return [][][]*partdata.PartData{}
	}

	ctx, stop := context.WithTimeout(context.Background(), time.Second)
	defer stop()

	rows, err := database.db.QueryContext(ctx, "SELECT * FROM parts WHERE `container`='"+container.Key.String+"'")
	defer rows.Close()
	if err != nil {
		log.Fatalf("Error on get parts, %s", err)
	}

	parts := [][][]*partdata.PartData{}
	for z := int32(0); z < container.Depth.Int32; z++ {
		yArr := [][]*partdata.PartData{}
		for y := int32(0); y < container.Height.Int32; y++ {
			xArr := []*partdata.PartData{}
			for x := int32(0); x < container.Width.Int32; x++ {
				xArr = append(xArr, nil)
			}
			yArr = append(yArr, xArr)
		}
		parts = append(parts, yArr)
	}

	for rows.Next() {
		part := partdata.FromDatabaseRow(rows)
		if part == nil {
			break
		}

		if !part.Row.Valid || !part.Column.Valid || !part.Depth.Valid {
			continue
		}

		row := part.Row.Int32 - int32(1)
		column := part.Column.Int32 - int32(1)
		depth := part.Depth.Int32 - int32(1)

		if row >= container.Height.Int32 || row < 0 {
			continue
		}
		if column >= container.Width.Int32 || column < 0 {
			continue
		}
		if depth >= container.Depth.Int32 || depth < 0 {
			continue
		}

		parts[depth][row][column] = part
	}
	return parts
}

// CreateContainer : Initializes a new container description
func (database PartsDatabase) CreateContainer(key string, name string) error {
	ctx, stop := context.WithTimeout(context.Background(), time.Second)
	defer stop()

	_, err := database.db.ExecContext(ctx, "INSERT INTO containers (`key`, `name`) VALUES ('"+key+"', '"+name+"')")
	if err != nil {
		return fmt.Errorf("Error on create container, %s", err)
	}
	return nil
}

// ArchiveContainer : Drops a container from the containers table (TODO: move to an archived containers table to make containers recoverable)
func (database PartsDatabase) ArchiveContainer(key string) {
	ctx, stop := context.WithTimeout(context.Background(), time.Second)
	defer stop()

	_, err := database.db.ExecContext(ctx, "DELETE FROM container WHERE `key`='"+key+"'")
	if err != nil {
		log.Fatalf("Error on archive container, %s", err)
	}
}

// UpdateContainer :
func (database PartsDatabase) UpdateContainer(container *containerdata.ContainerData) {
	ctx, stop := context.WithTimeout(context.Background(), time.Second)
	defer stop()

	query := "UPDATE containers SET"
	if container.Name.Valid {
		query = query + " `name`='" + container.Name.String + "'"
	}
	if container.Description.Valid {
		query = query + ", `description`='" + container.Description.String + "'"
	}
	if container.Height.Valid {
		query = query + ", `height`=" + strconv.FormatInt(int64(container.Height.Int32), 10)
	}
	if container.Width.Valid {
		query = query + ", `width`=" + strconv.FormatInt(int64(container.Width.Int32), 10)
	}
	if container.Depth.Valid {
		query = query + ", `depth`=" + strconv.FormatInt(int64(container.Depth.Int32), 10)
	}
	query = query + " WHERE `key`='" + container.Key.String + "'"

	_, err := database.db.ExecContext(ctx, query)
	if err != nil {
		log.Fatalf("Error on update container, %s", err)
	}
}
