package datastore

import (
	"encoding/json"
	"fmt"
	"os"
	"stringinator-go/interfaces"
	"stringinator-go/model"
	"sync"
)

type InMemoryStore struct {
	SeenStrings map[string]int
	FilePath    string
	Mutex       sync.Mutex
}

// NewInMemoryStore creates a new instance of InMemoryStore.
func NewInMemoryStore(filePath string) interfaces.Store {
	store := &InMemoryStore{
		SeenStrings: make(map[string]int),
		FilePath:    filePath,
	}

	file, err := os.Open(model.FilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}

	defer file.Close()

	//Read the file content
	content, err := os.ReadFile(file.Name())
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	//Unmarshal JSON data into the in-memory store
	err = json.Unmarshal(content, &store.SeenStrings)
	if err != nil {
		fmt.Println("Error occurred while unmarshaling JSON:", err)
		return nil
	}

	return store
}

// Add input to the in memory storage file
func (ims *InMemoryStore) SaveStrings(input string) error {
	ims.Mutex.Lock()
	defer ims.Mutex.Unlock()

	if ims.SeenStrings[input] == 0 {
		ims.SeenStrings[input] = 1
	} else {
		ims.SeenStrings[input] += 1
	}

	//Save data to file after each update.
	//Marshal the data to JSON
	jsonContent, err := json.Marshal(ims.SeenStrings)
	if err != nil {
		fmt.Println("Error occurred while marshaling accounts:", err)
		return err
	}

	//Write the JSON data to the file.

	err = os.WriteFile(ims.FilePath, jsonContent, 0644)
	if err != nil {
		fmt.Println("Error occurred while writing to file:", err)
		return err
	}
	return nil
}

// Get Seen strings as map from the in memory data storage file.
func (ims *InMemoryStore) GetStrings() (map[string]int, error) {
	ims.Mutex.Lock()
	defer ims.Mutex.Unlock()

	//Open the file
	file, err := os.Open(model.FilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}

	defer file.Close()

	//Read the file content
	content, err := os.ReadFile(file.Name())
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	//Unmarshal JSON data into the in-memory store
	err = json.Unmarshal(content, &ims.SeenStrings)
	if err != nil {
		fmt.Println("Error occurred while unmarshaling JSON:", err)
		return nil, err
	}

	return ims.SeenStrings, err

}
