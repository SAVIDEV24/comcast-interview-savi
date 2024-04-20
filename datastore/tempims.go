package datastore

import "stringinator-go/interfaces"

type tempIms struct {
	SeenStrings map[string]int
}

var NewTempIms = func(stringsMap map[string]int) interfaces.Store {
	return &tempIms{
		SeenStrings: stringsMap,
	}
}

func (t *tempIms) SaveStrings(input string) error {
	if t.SeenStrings[input] == 0 {
		t.SeenStrings[input] = 1
	} else {
		t.SeenStrings[input] += 1
	}

	return nil
}

func (t *tempIms) GetStrings() (map[string]int, error) {

	return t.SeenStrings, nil

}
