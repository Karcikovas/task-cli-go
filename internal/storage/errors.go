package storage

import "errors"

var ErrUnableToCreateNewStorageFile = errors.New("unable to create new storage file")
var ErrUnableToInsertOrUpdate = errors.New("unable to insert or update")
var ErrUnableToUnmarshalStorage = errors.New("unable to unmarshal storage")
var ErrUnableToMarshalStorage = errors.New("unable to marshal storage")
var ErrUnableWriteToFile = errors.New("unable write to file")
var ErrUnableToGetByID = errors.New("unable to get by id")
