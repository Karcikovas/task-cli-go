package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

var Location = "internal/storage"
var Format = "json"
var FileLocation = fmt.Sprintf(`%s/storage.%s`, Location, Format)

type Data struct {
	Total   int      `json:"total"`
	Records []string `json:"records"`
}

type Repository interface {
	GetAll() (Data, error)
	InsertOrUpdate(v []byte)
	Delete(id int)
}

type Storage struct {
}

func CreateNewStorage() Repository {
	return &Storage{}
}

func (s *Storage) InsertOrUpdate(v []byte) {
	data, _ := s.readFile()

	if data.Total == 0 {
		s.createFile()
	}

	data.Total += 1
	data.Records = append(data.Records, string(v))

	s.writeFile(data)
}

func (s *Storage) Delete(id int) {

}

func (s *Storage) GetAll() (Data, error) {
	return s.readFile()
}

func (s *Storage) readFile() (Data, error) {
	content, err := os.ReadFile(FileLocation)

	if err != nil {
		return Data{
			Total:   0,
			Records: []string{},
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
