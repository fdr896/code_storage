package memory

import (
	"github.com/rest-api/core"
)

// CodeStorage - structer which stores all code snippets
type CodeStorage struct {
	codes map[string]core.Code
}

// Get returns code object by its id if it exists otherwise it returns DoesNotExist error
func (cs *CodeStorage) Get(id string) (core.Code, error) {
	code := core.Code{}
	if _, ok := cs.codes[id]; !ok {
		return code, core.CodeDoesNotExist
	}

	return code, nil
}

// GetAll returns slice of all codes if storage has more than zero objects otherwise it returns ListEmpty error
func (cs *CodeStorage) GetAll() ([]core.Code, error) {
	result := make([]core.Code, len(cs.codes))

	for _, code := range cs.codes {
		result = append(result, code)
	}

	if len(result) == 0 {
		return result, core.CodeListIsEmpty
	}
	return result, nil
}

// AddCode adds new code snippet to database
func (cs *CodeStorage) AddCode(code *core.Code) error {
	cs.codes[code.ID] = *code
	return nil
}
