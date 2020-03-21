package core

import "errors"

var CodeDoesNotExist = errors.New("Code with required id does not exist")
var CodeListIsEmpty = errors.New("List of codes is empty")
