package showbank

import "bandolier/infrastructure"

type EventStoreBankRepository struct {
	aggregateStore infrastructure.AggregateStore
}

func NewEventStoreBankRepository(store infrastructure.AggregateStore) *EventStoreBankRepository {
	return &EventStoreBankRepository{
		aggregateStore: store,
	}
}

func (r *EventStoreBankRepository) Save(bank *Bank) {
	r.aggregateStore.Save(bank, infrastructure.CommandMetadata{})
}

func (r *EventStoreBankRepository) Get(showID string) (*Bank, error) {
	bank := NewBank()
	err := r.aggregateStore.Load(showID, bank)
	if err != nil {
		return nil, err
	}

	return bank, nil
}
