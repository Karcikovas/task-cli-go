package storage

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

var Location = "internal/storage"
var Format = "json"
var FileLocation = fmt.Sprintf(`%s/storage.%s`, Location, Format)

type Data map[string]interface{}

type Storage struct {
}

func CreateNewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Save(v []byte, k string) {
	file, err := os.Open(FileLocation)
	data := make(Data)

	if err != nil {
		s.createFile()
	} else {
		decoder := json.NewDecoder(file)
		err := decoder.Decode(&data)

		if err != nil {
			panic(ErrUnableToDecodeFile)
		}
	}

	data[k] = v

	file, err = os.Create(FileLocation)

	if err != nil {
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(data)
}

func (s *Storage) Delete(k string) {
	data, err := s.GetAllStorageData()

	if err != nil {
		panic(ErrUnableToGetAllStorageItems)
	}

	delete(data, k)

	s.storeAllData(data)
}

func (s *Storage) createFile() {
	file, err := os.Create(FileLocation)

	if err != nil {
		panic(ErrUnableToCreateNewStorageFile)
	}

	defer file.Close()
}

func (s *Storage) GetAllStorageData() (Data, error) {
	data := make(Data)
	file, err := os.Open(FileLocation)

	if err != nil {
		panic(ErrUnableToReadFile)

		return nil, err
	}

	scanner := bufio.NewScanner(file)
	var jsonData map[string]string

	content := ""
	for scanner.Scan() {
		content += scanner.Text()
	}

	err = json.Unmarshal([]byte(content), &jsonData)
	if err != nil {
		panic(ErrUnableToUnmarshalJson)

		return nil, err
	}

	for key, encodedValue := range jsonData {
		decodedValue, err := base64.StdEncoding.DecodeString(encodedValue)
		if err != nil {
			panic(ErrUnableToDecodeFile)

			return nil, err
		}

		data[key] = string(decodedValue)
	}

	return data, nil
}

func (s *Storage) storeAllData(data Data) {
	file, err := os.Create(FileLocation)

	defer file.Close()

	if err != nil {
		return
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(data)
}
