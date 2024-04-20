package interfaces

type Store interface {
	SaveStrings(input string) error
	GetStrings() (map[string]int, error)
}
