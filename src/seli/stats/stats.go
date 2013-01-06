package stats

import "encoding/json"

var instantiated *selistats = nil

type selistats struct {
	TotalPackages int
	LastRepositoryModifiedTime int
}

func (seliStats *selistats) Update() {
	seliStats.TotalPackages += 1
}

func (seliStats *selistats) Resource() ([]byte, error) {
	return json.Marshal(instantiated)
}



func SeliStats() *selistats {
	if instantiated == nil {
		instantiated = new(selistats)
	}
	return instantiated
}
