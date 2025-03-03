package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

var Location = "internal/storage"
var Format = "txt"
var FileLocation = fmt.Sprintf(`%s/storage.%s`, Location, Format)

type Data map[string][]byte

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
	}

	data[k] = v

	file, err = os.Create(FileLocation)

	if err != nil {
		return
	}

	_, err = file.WriteString(fmt.Sprintf(`%s->%s`, k, string(v)))

	if err != nil {
		return
	}

	defer file.Close()
}

func (s *Storage) Delete(k string) {
	data, err := s.GetAllStorageData()

	if err != nil {
		panic(ErrUnableToGetAllStorageItems)
	}

	delete(data, k)

	s.storeAllData(data)
}

func (s *Storage) GetAllStorageData() (Data, error) {
	data := make(Data)
	file, err := os.Open(FileLocation)

	if err != nil {
		panic(ErrUnableToReadFile)

		return nil, err
	}

	scanner := bufio.NewScanner(file)

	content := ""
	for scanner.Scan() {
		content += scanner.Text()
		result := strings.Split(content, "->")

		log.Println(len(result))
		k := result[0]
		v := []byte(result[1])

		data[k] = v
	}

	return data, nil
}

func (s *Storage) createFile() {
	file, err := os.Create(FileLocation)

	if err != nil {
		panic(ErrUnableToCreateNewStorageFile)
	}

	defer file.Close()
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
