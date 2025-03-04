package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var Location = "internal/storage"
var Format = "json"
var FileLocation = fmt.Sprintf(`%s/storage.%s`, Location, Format)
var red = "\033[31m"

type RecordMap = map[string]string

type Data struct {
	Total   int       `json:"total"`
	Records RecordMap `json:"records"`
}

type Repository interface {
	GetOneBy(id int) (string, error)
	GetAll() (Data, error)
	InsertOrUpdate(v []byte) (*string, error)
	Delete(id string) (bool, error)
}

type Storage struct {
}

func CreateNewStorage() Repository {
	return &Storage{}
}

func (s *Storage) GetOneBy(id int) (string, error) {

	return "asdas", nil
}

func (s *Storage) InsertOrUpdate(v []byte) (*string, error) {
	data, err := s.readFile()

	if err != nil {
		log.Println(red + err.Error())
		return nil, ErrUnableToInsertOrUpdate
	}

	if data.Total == 0 {
		created, err := s.createFile()

		if !created && err != nil {
			log.Println(red + err.Error())
			return nil, ErrUnableToInsertOrUpdate
		}
	}

	data.Total += 1

	id := string(data.Total)
	data.Records[id] = string(v)

	update, err := s.writeFile(data)

	if !update && err != nil {
		log.Println(red + err.Error())
		return nil, ErrUnableToInsertOrUpdate
	}

	value := string(v)

	return &value, nil
}

func (s *Storage) Delete(id string) (bool, error) {
	data, err := s.readFile()

	if err != nil {
		log.Println(red + err.Error())
		return false, nil
	}

	delete(data.Records, id)

	return true, nil
}

func (s *Storage) GetAll() (Data, error) {
	return s.readFile()
}

func (s *Storage) readFile() (Data, error) {
	content, err := os.ReadFile(FileLocation)

	if err != nil {
		return Data{
			Total:   0,
			Records: make(RecordMap),
		}, nil
	}

	storage := Data{}

	err = json.Unmarshal(content, &storage)
	if err != nil {
		log.Println(red + err.Error())
		return storage, ErrUnableToUnmarshalStorage
	}
	return storage, nil
}

func (s *Storage) writeFile(data Data) (bool, error) {
	bytes, err := json.Marshal(data)

	if err != nil {
		log.Println(red + err.Error())
		return false, ErrUnableToMarshalStorage
	}

	err = os.WriteFile(FileLocation, bytes, 0644)
	if err != nil {
		log.Println(red + err.Error())
		return false, ErrUnableWriteToFile
	}
	return true, nil
}

func (s *Storage) createFile() (bool, error) {
	file, err := os.Create(FileLocation)

	if err != nil {
		log.Println(red + err.Error())
		return false, ErrUnableToCreateNewStorageFile
	}

	defer file.Close()

	return true, nil
}
