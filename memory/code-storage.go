package memory

import (
	"github.com/rest-api/core"
)

// CodeStorage - structer which stores all code snippets
type CodeStorage struct {
	codes map[string]core.Code
}

// NewCodeStorage return a void instance of CodeStorage type
func NewCodeStorage(codes map[string]core.Code) CodeStorage {
	return CodeStorage{codes}
}

// Get returns code object by its id if it exists otherwise it returns DoesNotExist error
func (cs *CodeStorage) Get(id string) (core.Code, error) {
	if _, ok := cs.codes[id]; !ok {
		return core.Code{}, core.CodeDoesNotExist
	}

	return cs.codes[id], nil
}

// GetAll returns slice of all codes if storage has more than zero objects otherwise it returns ListEmpty error
func (cs *CodeStorage) GetAll() ([]core.Code, error) {
	result := make([]core.Code, 0)

	for _, code := range cs.codes {
		result = append(result, code)
	}

	if len(result) == 0 {
		return result, core.CodeListIsEmpty
	}
	return result, nil
}

// Add adds new code snippet to database
func (cs *CodeStorage) Add(code *core.Code) error {
	cs.codes[code.ID] = core.NewCode(code)
	return nil
}

// Delete deletes code from database by its id
func (cs *CodeStorage) Delete(id string) error {
	delete(cs.codes, id)
	return nil
}
