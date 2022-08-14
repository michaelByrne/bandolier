package venueshow

import "time"

type Slot struct {
	ID          string
	Order       int
	Booked      bool
	Duration    time.Duration
	StartTime   time.Time
	IsHeadliner bool
}

func NewSlot(id string, start time.Time, booked bool, duration time.Duration) *Slot {
	return &Slot{
		ID:        id,
		StartTime: start,
		Booked:    booked,
		Duration:  duration,
	}
}

func (s *Slot) Book() {
	s.Booked = true
}

func (s *Slot) Cancel() {
	s.Booked = false
}

func (s *Slot) SetHeadlineSlot(headliner bool) {
	s.IsHeadliner = headliner
}

func (s *Slot) SetOrder(order int) {
	s.Order = order
}

func (s *Slot) Overlaps(start time.Time, duration time.Duration) bool {
	thisStart := s.StartTime
	thisEnd := s.StartTime.Add(s.Duration)
	proposedStart := start
	proposedEnd := proposedStart.Add(duration)

	return thisStart.Before(proposedEnd) && thisEnd.After(proposedStart)
}
