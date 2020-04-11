package bolt

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/fdr896/code_storage/rest-api/core"
)

// CodeStorage stores all code snippets.
type codeStorage struct {
	bucketName []byte
	Codes      *bolt.DB
}

// NewCodeStorage return a void instance of CodeStorage type.
func NewCodeStorage(bucketName []byte, path string, db *bolt.DB) (core.CodeStorage, error) {
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	}); err != nil {
		return nil, err
	}

	return &codeStorage{bucketName, db}, nil
}

func (cs *codeStorage) Close() {
	cs.Codes.Close()
}

// Get returns code object by its id if it exists otherwise it returns DoesNotExist error.
func (cs *codeStorage) Get(id string) (*core.Code, error) {
	code := &core.Code{}
	if err := cs.Codes.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(cs.bucketName)

		tmp := b.Get([]byte(id))
		if tmp == nil {
			return core.ErrNotFound
		}

		return json.Unmarshal(tmp, &code)
	}); err != nil {
		return nil, err
	}

	return code, nil
}

// GetAll returns slice of all codes if storage has more than zero objects otherwise it returns ListEmpty error.
func (cs *codeStorage) GetAll() ([]*core.Code, error) {
	result := make([]*core.Code, 0)

	if err := cs.Codes.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(cs.bucketName)

		err := b.ForEach(func(key []byte, value []byte) error {
			var decoded core.Code
			if err := json.Unmarshal(value, &decoded); err != nil {
				return err
			}

			result = append(result, &decoded)

			return nil
		})

		return err
	}); err != nil {
		return nil, err
	}

	return result, nil
}

// Add adds new code snippet to database.
func (cs *codeStorage) Add(code *core.Code) error {
	return cs.Codes.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(cs.bucketName)

		encoded, err := json.Marshal(code)
		if err != nil {
			return err
		}

		return b.Put([]byte(code.ID), encoded)
	})
}

// Delete deletes code from database by its id.
func (cs *codeStorage) Delete(id string) error {
	return cs.Codes.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(cs.bucketName)

		return b.Delete([]byte(id))
	})
}
