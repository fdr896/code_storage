package memory

import (
	"github.com/rest-api/core"
)

// CodeStorage - structer which stores all code snippets
type codeStorage struct {
	codes map[string]core.Code
}

// NewCodeStorage return a void instance of CodeStorage type
func NewCodeStorage(codes map[string]core.Code) codeStorage {
	return codeStorage{codes}
}

// Get returns code object by its id if it exists otherwise it returns DoesNotExist error
func (cs *codeStorage) Get(id string) (core.Code, error) {
	if _, ok := cs.codes[id]; !ok {
		return core.Code{}, core.CodeDoesNotExist.ErrorMessage
	}

	return cs.codes[id], nil
}

// GetAll returns slice of all codes if storage has more than zero objects otherwise it returns ListEmpty error
func (cs *codeStorage) GetAll() ([]core.Code, error) {
	result := make([]core.Code, 0)

	for _, code := range cs.codes {
		result = append(result, code)
	}

	return result, nil
}

// Add adds new code snippet to database
func (cs *codeStorage) Add(code *core.Code) error {
	if !core.CheckCode(code) {
		return core.UnsupportedJSON.ErrorMessage
	}

	cs.codes[code.ID] = core.NewCode(code)
	return nil
}

// Delete deletes code from database by its id
func (cs *codeStorage) Delete(id string) error {
	if _, ok := cs.codes[id]; !ok {
		return core.CodeDoesNotExist.ErrorMessage
	}

	delete(cs.codes, id)
	return nil
}
