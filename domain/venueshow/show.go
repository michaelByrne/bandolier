package venueshow

import (
	"bandolier/domain/venueshow/commands"
	"bandolier/domain/venueshow/events"
	"bandolier/eventsourcing"
	"github.com/google/uuid"
	"time"
)

type Show struct {
	eventsourcing.AggregateBase

	isArchived  bool
	isCancelled bool
	isScheduled bool
	slots       Slots
}

func NewShow() *Show {
	a := &Show{
		AggregateBase: eventsourcing.NewAggregateRoot(),
	}

	a.Register(events.ShowScheduled{}, func(e interface{}) { a.ShowScheduled(e.(events.ShowScheduled)) })
	a.Register(events.SlotScheduled{}, func(e interface{}) { a.SlotScheduled(e.(events.SlotScheduled)) })
	a.Register(events.SlotBooked{}, func(e interface{}) { a.SlotBooked(e.(events.SlotBooked)) })
	a.Register(events.SlotBookingCancelled{}, func(e interface{}) { a.SlotBookingCancelled(e.(events.SlotBookingCancelled)) })
	a.Register(events.SlotScheduleCancelled{}, func(e interface{}) { a.SlotScheduleCancelled(e.(events.SlotScheduleCancelled)) })
	a.Register(events.ShowScheduleCancelled{}, func(e interface{}) { a.ShowScheduleCancelled(e.(events.ShowScheduleCancelled)) })
	a.Register(events.ShowArchived{}, func(e interface{}) { a.ShowArchived(e.(events.ShowArchived)) })

	return a
}

// COMMANDS

func (s *Show) ScheduleShow(venue Venue, time time.Time, slots []commands.ScheduledSlot) error {
	err := s.isShowCancelledOrArchived()
	if err != nil {
		return err
	}

	if s.isScheduled {
		return ShowAlreadyScheduledError{}
	}

	showID := NewShowID(venue.ID, time)
	s.Raise(events.NewShowScheduled(showID.Value, venue.ID, time))

	for _, slot := range slots {
		s.Raise(events.NewSlotScheduled(uuid.New().String(), showID.Value, slot.StartTime, slot.Duration))
	}
	return nil
}

func (s *Show) ScheduleSlot(id string, start time.Time, duration time.Duration) error {
	err := s.isShowCancelledOrArchived()
	if err != nil {
		return err
	}

	if !s.isScheduled {
		return ShowNotScheduledError{}
	}
	if s.slots.Overlaps(start, duration) {
		return SlotOverlapsError{}
	}

	s.Raise(events.NewSlotScheduled(id, s.Id, start, duration))
	return nil
}

func (s *Show) BookSlot(id, artistID, artistName string, headliner bool) error {
	err := s.isShowCancelledOrArchived()
	if err != nil {
		return err
	}

	if !s.isScheduled {
		return ShowNotScheduledError{}
	}
	slotStatus := s.slots.GetStatus(id)

	switch slotStatus {
	case SlotAvailable:
		s.Raise(events.NewSlotBooked(id, s.Id, artistID, artistName, headliner))
		return nil
	case SlotBooked:
		return SlotAlreadyBookedError{}
	case SlotDoesNotExist:
		return SlotDoesNotExistError{}
	default:
		return InvalidSlotStatusError{}
	}
}

func (s *Show) CancelSlotBooking(id string) error {
	err := s.isShowCancelledOrArchived()
	if err != nil {
		return err
	}

	if !s.isScheduled {
		return ShowNotScheduledError{}
	}

	if !s.slots.HasBookedSlot(id) {
		return &SlotNotBookedError{}
	}

	s.Raise(events.NewSlotBookingCancelled(id))
	return nil
}

func (s *Show) Cancel() error {
	err := s.isShowCancelledOrArchived()
	if err != nil {
		return err
	}

	if !s.isScheduled {
		return ShowNotScheduledError{}
	}

	for _, bookedSlot := range s.slots.GetBookedSlots() {
		s.Raise(events.NewSlotBookingCancelled(bookedSlot.ID))
	}

	for _, bookedSlot := range s.slots.GetAllSlots() {
		s.Raise(events.NewSlotScheduleCancelled(bookedSlot.ID))
	}

	s.Raise(events.NewShowScheduleCancelled(s.Id))
	return nil
}

func (s *Show) Archive(amount int) error {
	if !s.isScheduled {
		return ShowNotScheduledError{}
	}

	if s.isArchived {
		return ShowAlreadyArchivedError{}
	}

	s.Raise(events.NewShowArchived(s.Id, amount))
	return nil
}

// EVENTS

func (s *Show) ShowScheduled(e events.ShowScheduled) {
	s.Id = NewShowID(e.VenueID, e.Date).Value
	s.isScheduled = true
}

func (s *Show) SlotScheduled(e events.SlotScheduled) {
	s.slots.Add(e.ID, e.StartTime, e.Duration, false)
}

func (s *Show) SlotBooked(e events.SlotBooked) {
	s.slots.MarkAsBooked(e.ID)
}

func (s *Show) SlotBookingCancelled(e events.SlotBookingCancelled) {
	s.slots.MarkAsAvailable(e.ID)
}

func (s *Show) SlotScheduleCancelled(e events.SlotScheduleCancelled) {
	s.slots.Remove(e.ID)
}

func (s *Show) ShowScheduleCancelled(_ events.ShowScheduleCancelled) {
	s.isCancelled = true
}

func (s *Show) ShowArchived(_ events.ShowArchived) {
	s.isArchived = true
}

// HELPERS

func (s *Show) isShowCancelledOrArchived() error {
	if s.isArchived {
		return ShowScheduleAlreadyArchivedError{}
	}

	if s.isCancelled {
		return ShowScheduleAlreadyCancelledError{}
	}

	return nil
}
