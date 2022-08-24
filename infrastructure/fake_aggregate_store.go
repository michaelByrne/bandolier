package infrastructure

import (
	"bandolier/eventsourcing"
)

type FakeAggregateStore struct {
	AggregateStore

	aggregate     eventsourcing.AggregateRoot
	initialEvents []interface{}
}

func NewFakeAggregateStore() *FakeAggregateStore {
	return &FakeAggregateStore{}
}

func (f *FakeAggregateStore) SetInitialEvents(events []interface{}) {
	f.initialEvents = events
}

func (f *FakeAggregateStore) GetStoredChanges() []interface{} {
	if f.aggregate != nil {
		return f.aggregate.GetChanges()
	}

	return []interface{}{}
}

func (f *FakeAggregateStore) Save(a eventsourcing.AggregateRoot, m CommandMetadata) error {
	f.aggregate = a
	return nil
}

func (f *FakeAggregateStore) Load(aggregateId string, a eventsourcing.AggregateRoot) error {
	a.Load(f.initialEvents)
	a.ClearChanges()
	return nil
}
