package memory

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/rest-api/core"
)

// CodeStorage - structer which stores all code snippets
type codeStorage struct {
	bucketName []byte
	Codes      *bolt.DB
}

// NewCodeStorage return a void instance of CodeStorage type
func NewCodeStorage(bucketName []byte) codeStorage {
	return codeStorage{bucketName: bucketName}
}

// Get returns code object by its id if it exists otherwise it returns DoesNotExist error
func (cs *codeStorage) Get(id string) (core.Code, error) {
	code := core.Code{}
	err := cs.Codes.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(cs.bucketName)

		tmp := b.Get([]byte(id))

		if tmp == nil {
			return core.CodeDoesNotExist.ErrorMessage
		}

		return json.Unmarshal(tmp, &code)
	})

	if err != nil {
		return code, err
	}

	return code, nil
}

// GetAll returns slice of all codes if storage has more than zero objects otherwise it returns ListEmpty error
func (cs *codeStorage) GetAll() ([]core.Code, error) {
	result := make([]core.Code, 0)

	err := cs.Codes.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(cs.bucketName)

		err := b.ForEach(func(key []byte, value []byte) error {
			var tmp core.Code
			if err := json.Unmarshal(value, &tmp); err != nil {
				return err
			}

			result = append(result, tmp)

			return nil
		})

		return err
	})

	return result, err
}

// Add adds new code snippet to database
func (cs *codeStorage) Add(code *core.Code) error {
	if !core.CheckCode(code) {
		return core.UnsupportedJSON.ErrorMessage
	}

	return cs.Codes.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(cs.bucketName)

		code := core.NewCode(code)

		id := code.ID
		converted, err := json.Marshal(code)

		if err != nil {
			return err
		}

		return b.Put([]byte(id), converted)
	})
}

// Delete deletes code from database by its id
func (cs *codeStorage) Delete(id string) error {
	return cs.Codes.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(cs.bucketName)

		return b.Delete([]byte(id))
	})
}
