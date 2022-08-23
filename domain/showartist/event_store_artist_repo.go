package showartist

import "bandolier/infrastructure"

type EventStoreArtistRepository struct {
	aggregateStore infrastructure.AggregateStore
}

func NewEventStoreArtistRepository(store infrastructure.AggregateStore) *EventStoreArtistRepository {
	return &EventStoreArtistRepository{
		aggregateStore: store,
	}
}

func (r *EventStoreArtistRepository) Save(bank *Artist) {
	r.aggregateStore.Save(bank, infrastructure.CommandMetadata{})
}

func (r *EventStoreArtistRepository) Get(artistID string) (*Artist, error) {
	artist := NewArtist()
	err := r.aggregateStore.Load(artistID, artist)
	if err != nil {
		return nil, err
	}

	return artist, nil
}
