package storage

import "errors"

var ErrUnableToCreateNewStorageFile = errors.New("unable to create new storage file")
var ErrUnableToGetAllStorageItems = errors.New("unable to get all storage items")
var ErrUnableToDecodeFile = errors.New("unable to decode file")
var ErrUnableToUnmarshalJson = errors.New("unable to unmarshal json")
var ErrUnableToReadFile = errors.New("unable to read file")
var ErrUnableToSaveInStorage = errors.New("unable to save in storage")
