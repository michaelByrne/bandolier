package inmemory

import (
	"bandolier/domain/readmodel"
	"time"
)

type AvailableSlotsRepository struct {
	readmodel.AvailableSlotsRepository

	slots map[string]*readmodel.AvailableSlot
}

func NewAvailableSlotsRepository() *AvailableSlotsRepository {
	return &AvailableSlotsRepository{
		slots: make(map[string]*readmodel.AvailableSlot),
	}
}

func (r *AvailableSlotsRepository) Add(slot readmodel.AvailableSlot) {
	r.slots[slot.SlotID] = &slot
}

func (r *AvailableSlotsRepository) Delete(slotId string) {
	delete(r.slots, slotId)
}

func (r *AvailableSlotsRepository) MarkAsUnavailable(slotId string) {
	slot := r.slots[slotId]
	if slot != nil {
		slot.Available = false
	}
}

func (r *AvailableSlotsRepository) MarkAsAvailable(slotId string) {
	slot := r.slots[slotId]
	if slot != nil {
		slot.Available = true
	}
}

func (r *AvailableSlotsRepository) GetSlotsAvailableOn(time time.Time) []*readmodel.AvailableSlot {
	availabilityYear := time.Year()
	availabilityYearDay := time.YearDay()
	availableOnDate := make([]*readmodel.AvailableSlot, 0)
	for _, slot := range r.slots {
		scheduledStart := slot.Start
		if availabilityYear == scheduledStart.Year() && availabilityYearDay == scheduledStart.YearDay() {
			availableOnDate = append(availableOnDate, slot)
		}
	}
	return availableOnDate
}

func (r *AvailableSlotsRepository) Clear() {
	r.slots = make(map[string]*readmodel.AvailableSlot)
}
