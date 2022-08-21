package infrastructure

import (
	"fmt"

	"bandolier/eventsourcing"
)

type AggregateStore interface {
	Save(a eventsourcing.AggregateRoot, m CommandMetadata) error
	Load(aggregateId string, a eventsourcing.AggregateRoot) error
}

type AggregateNotFoundError struct{}

func (e AggregateNotFoundError) Error() string {
	return fmt.Sprintf("aggregate not found error")
}
