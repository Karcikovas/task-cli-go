package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"task-cli-go/internal/logger"
)

var Location = "internal/storage"
var Format = "json"
var FileLocation = fmt.Sprintf(`%s/storage.%s`, Location, Format)

type RecordMap = map[string]string

type Data struct {
	Total             int       `json:"total"`
	AutoIncrementedID int       `json:"AutoIncrementedID"`
	Records           RecordMap `json:"records"`
}

type Repository interface {
	GetOneBy(id string) (*string, error)
	GetAll() (Data, error)
	Upsert(id string, v []byte) (*string, error)
	Delete(id string) (bool, error)
	GenerateID(data Data) int
}

type Storage struct {
	logger logger.Service
}

func CreateNewStorage(logger logger.Service) Repository {
	return &Storage{
		logger: logger,
	}
}

func (s *Storage) GetOneBy(id string) (*string, error) {
	data, err := s.readFile()

	if err != nil {
		s.logger.LogError(err.Error())
		return nil, ErrUnableToGetByID
	}

	if data.Total > 0 {
		item := data.Records[id]

		return &item, nil
	}

	return nil, ErrUnableToGetByID
}

func (s *Storage) Upsert(id string, v []byte) (*string, error) {
	data, e := s.readFile()

	if e != nil {
		s.logger.LogError(e.Error())
		return nil, ErrUnableToInsertOrUpdate
	}

	if data.Total == 0 {
		created, err := s.createFile()

		if !created && err != nil {
			s.logger.LogError(err.Error())
			return nil, ErrUnableToInsertOrUpdate
		}
	}

	if len(data.Records[id]) == 0 {
		data.AutoIncrementedID++
		data.Total++
	}

	data.Records[id] = string(v)

	update, err := s.writeFile(data)

	if !update || err != nil {
		s.logger.LogError(err.Error())
		return nil, ErrUnableToInsertOrUpdate
	}

	value := string(v)

	return &value, nil
}

func (s *Storage) Delete(id string) (bool, error) {
	data, err := s.readFile()

	if err != nil {
		s.logger.LogError(err.Error())
		return false, nil
	}

	delete(data.Records, id)
	data.Total -= 1

	update, err := s.writeFile(data)

	if !update || err != nil {
		s.logger.LogError(err.Error())
		return false, ErrUnableToDelete
	}

	return true, nil
}

func (s *Storage) GetAll() (Data, error) {
	return s.readFile()
}

func (s *Storage) GenerateID(data Data) int {
	return data.AutoIncrementedID + 1
}

func (s *Storage) readFile() (Data, error) {
	content, err := os.ReadFile(FileLocation)

	if err != nil {
		return Data{
			Total:             0,
			AutoIncrementedID: 0,
			Records:           make(RecordMap),
		}, nil
	}

	storage := Data{}

	err = json.Unmarshal(content, &storage)
	if err != nil {
		s.logger.LogError(err.Error())
		return storage, ErrUnableToUnmarshalStorage
	}
	return storage, nil
}

func (s *Storage) writeFile(data Data) (bool, error) {
	bytes, err := json.Marshal(data)

	if err != nil {
		s.logger.LogError(err.Error())
		return false, ErrUnableToMarshalStorage
	}

	err = os.WriteFile(FileLocation, bytes, 0600)
	if err != nil {
		s.logger.LogError(err.Error())
		return false, ErrUnableWriteToFile
	}
	return true, nil
}

func (s *Storage) createFile() (bool, error) {
	file, err := os.Create(FileLocation)

	if err != nil {
		s.logger.LogError(err.Error())
		return false, ErrUnableToCreateNewStorageFile
	}

	defer file.Close()

	return true, nil
}
