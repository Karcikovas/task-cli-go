package task

import "errors"

var ErrUnableToCreateNewTask = errors.New("unable to create new task")
var ErrUnableToDeleteTask = errors.New("unable to delete task")
var ErrUnableToGetStorageData = errors.New("unable to get storage data")
