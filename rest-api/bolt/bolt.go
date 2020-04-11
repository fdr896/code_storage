package bolt

import (
	"log"

	"github.com/boltdb/bolt"
	"github.com/fdr896/code_storage/rest-api/core"
)

// Bolt contains all storages.
type Bolt struct {
	CodeStorage core.CodeStorage
	UserStorage core.UserStorage
}

// New initializes all fields of Bold instance.
func New(path string, bucketsNames [][]byte) *Bolt {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	codeStorage, err := NewCodeStorage(bucketsNames[0], path, db)
	if err != nil {
		log.Fatal(err)
	}

	userStorage, err := NewUserStorage(bucketsNames[1], path, db)
	if err != nil {
		log.Fatal(err)
	}

	return &Bolt{
		CodeStorage: codeStorage,
		UserStorage: userStorage,
	}
}
