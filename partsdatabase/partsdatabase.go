package partsdatabase

import (
	"fmt"

	"github.com/ruffaustin25/ElectronicsSorting/common"
)

// Init :
func Init() {

}

// GetPartsList : Gets a list of parts in the database
// count : how many parts to retrieve
// startAt : which element to start at (for pagination)
func GetPartsList() []common.PartData {
	partsList := []common.PartData{}
	for i := 0; i < 500; i++ {
		partsList = append(partsList, common.PartData{Name: fmt.Sprintf("Part %d", i), Container: "0"})
	}
	return partsList
}
