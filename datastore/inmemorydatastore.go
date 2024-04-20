package datastore

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type InMemoryStore struct {
	SeenStrings map[string]int
	FilePath    string
	Mutex       sync.Mutex
}

// NewInMemoryStore creates a new instance of InMemoryStore.
func NewInMemoryStore(filePath string) *InMemoryStore {
	store := &InMemoryStore{
		SeenStrings: make(map[string]int),
		FilePath:    filePath,
	}

	store.loadFromFile()

	return store
}

// Add input to the in memory storage file
func (ims *InMemoryStore) AddInput(input string) error {
	ims.Mutex.Lock()
	defer ims.Mutex.Unlock()

	if ims.SeenStrings[input] == 0 {
		ims.SeenStrings[input] = 1
	} else {
		ims.SeenStrings[input] += 1
	}

	//Save data to file after each update.
	err := ims.SaveToFile()

	if err != nil {
		fmt.Println("Error occurred while adding seen string:,", err)
		return err
	}
	return nil
}

// loadFromFile loads data from the storage file into the in- memory store.
func (ims *InMemoryStore) loadFromFile() {
	//Open the file
	file, err := os.Open(ims.FilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	//Read the file content
	content, err := os.ReadFile(file.Name())
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	stringmap := make(map[string]int)
	//Unmarshal JSON data into the in-memory store
	err = json.Unmarshal(content, &stringmap)
	if err != nil {
		fmt.Println("Error occurred while unmarshaling JSON:", err)
		return
	}

	ims.SeenStrings = stringmap
}

// SaveToFile saves the in-memory data to the persistent storage file.
func (ims *InMemoryStore) SaveToFile() error {

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
func (ims *InMemoryStore) GetSeenStrings() map[string]int {
	ims.Mutex.Lock()
	defer ims.Mutex.Unlock()

	ims.loadFromFile()

	return ims.SeenStrings

}
