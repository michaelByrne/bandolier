package inmemory

import "bandolier/eventsourcing"

type ColdStorage struct {
	eventsourcing.ColdStorage

	Events []interface{}
}

func NewColdStorage() *ColdStorage {
	return &ColdStorage{
		Events: make([]interface{}, 0),
	}
}

func (s *ColdStorage) SaveAll(events []interface{}) {
	for _, e := range events {
		s.Events = append(s.Events, e)
	}
}
