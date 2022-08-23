package showvenue

import (
	"bandolier/infrastructure"
)

type EventStoreShowRepository struct {
	aggregateStore infrastructure.AggregateStore
}

func NewEventStoreShowRepository(store infrastructure.AggregateStore) *EventStoreShowRepository {
	return &EventStoreShowRepository{
		aggregateStore: store,
	}
}

func (r *EventStoreShowRepository) Save(show *Show) {
	r.aggregateStore.Save(show, infrastructure.CommandMetadata{})
}

func (r *EventStoreShowRepository) Get(id ShowID) (*Show, error) {
	show := NewShow()
	err := r.aggregateStore.Load(id.Value, show)
	if err != nil {
		return nil, err
	}

	return show, nil
}
