package storage

import (
	"fmt"
	"log"
	"os"
)

var Location = "internal/storage"
var Format = "json"

type Storage struct {
}

func CreateNewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Save(v []byte, k string) {
	exist := s.storageExists()

	if !exist {
		s.createFile()
	}

	os.WriteFile("storage.json", v, 0644)

	log.Println(v, k)
}

func (s *Storage) createFile() {
	f, err := os.Create(fmt.Sprintf(`%s/storage.%s`, Location, Format))

	if err != nil {
		panic(ErrUnableToCreateNewStorageFile)
		log.Fatal(err)
	}

	defer f.Close()
}

func (s *Storage) storageExists() bool {
	if _, err := os.Stat(fmt.Sprintf(`%s/storage.%s`, Location, Format)); err == nil {
		return true
	} else {
		return false
	}
}
