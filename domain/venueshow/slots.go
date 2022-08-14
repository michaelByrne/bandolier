package venueshow

import (
	"time"
)

type Slots struct {
	slots []*Slot
}

func (s *Slots) Add(id string, start time.Time, duration time.Duration, booked bool) {
	s.slots = append(s.slots, NewSlot(id, start, booked, duration))
}

func (s *Slots) Remove(id string) {
	slots := make([]*Slot, 0)
	for _, slot := range s.slots {
		if slot.ID != id {
			slots = append(slots, slot)
		}
	}

	s.slots = slots
}

func (s *Slots) MarkAsBooked(id string) {
	slot := s.getSlot(id)
	if slot != nil {
		slot.Book()
	}
}

func (s *Slots) MarkAsAvailable(id string) {
	slot := s.getSlot(id)
	if slot != nil {
		slot.Cancel()
	}
}

func (s *Slots) HasBookedSlot(id string) bool {
	slot := s.getSlot(id)
	if slot == nil {
		return false
	}

	return slot.Booked
}

func (s *Slots) GetBookedSlots() []*Slot {
	bookedSlots := make([]*Slot, 0)
	for _, slot := range s.slots {
		if slot.Booked {
			bookedSlots = append(bookedSlots, slot)
		}
	}

	return bookedSlots
}

func (s *Slots) GetAllSlots() []*Slot {
	return s.slots
}

func (s *Slots) GetStatus(id string) SlotStatus {
	slot := s.getSlot(id)
	if slot == nil {
		return SlotDoesNotExist
	}

	if slot.Booked {
		return SlotBooked
	}

	return SlotAvailable
}

func (s *Slots) getSlot(id string) *Slot {
	for _, slot := range s.slots {
		if slot.ID == id {
			return slot
		}
	}

	return nil
}

func (s *Slots) Overlaps(start time.Time, duration time.Duration) bool {
	for _, slot := range s.slots {
		if slot.Overlaps(start, duration) {
			return true
		}
	}

	return false
}
