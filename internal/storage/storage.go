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

type RecordMap = map[string]string

type Data struct {
	Total   int       `json:"total"`
	Records RecordMap `json:"records"`
}

type Repository interface {
	GetOneBy(id int) (string, error)
	GetAll() (Data, error)
	InsertOrUpdate(v []byte) (string, error)
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

func (s *Storage) InsertOrUpdate(v []byte) (string, error) {
	data, err := s.readFile()

	log.Println("ERRORRR ", err.Error(), "data ", data)

	if data.Total == 0 {
		s.createFile()
	}

	data.Total += 1

	id := string(data.Total)
	data.Records[id] = string(v)

	s.writeFile(data)

	return string(v), nil
}

func (s *Storage) Delete(id string) (bool, error) {
	data, err := s.readFile()

	if err != nil {
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
		}, err
	}

	storage := Data{}

	err = json.Unmarshal(content, &storage)
	if err != nil {
		return storage, err
	}
	return storage, nil
}

func (s *Storage) writeFile(data Data) error {
	bytes, err := json.Marshal(data)

	if err != nil {
		return err
	}

	err = os.WriteFile(FileLocation, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) createFile() {
	file, err := os.Create(FileLocation)

	if err != nil {
		panic(ErrUnableToCreateNewStorageFile)
	}

	defer file.Close()
}
