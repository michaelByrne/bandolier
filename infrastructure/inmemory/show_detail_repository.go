package inmemory

//
//import (
//	"bandolier/domain/readmodel"
//	"time"
//)
//
//type ShowDetailRepository struct {
//	readmodel.ShowDetailRepository
//
//	shows map[string]*readmodel.ShowDetail
//}
//
//func NewShowDetailRepository() *ShowDetailRepository {
//	return &ShowDetailRepository{
//		shows: make(map[string]*readmodel.ShowDetail),
//	}
//}
//
//func (r *ShowDetailRepository) AddShow(show readmodel.ShowDetail) {
//	r.shows[show.ShowID] = &show
//}
//
//func (r *ShowDetailRepository) AddSlotToShow(slotID, showID string, start time.Time, duration time.Duration) error {
//	show := r.shows[showID]
//	if show == nil {
//		return ShowNotFoundError{showID}
//	}
//
//	show.Slots = append(show.Slots, *readmodel.NewSlot(slotID, start, duration))
//	r.shows[showID] = show
//
//	return nil
//}
//
//func (r *ShowDetailRepository) SetHeadliner(showID string, headliner string) error {
//	show := r.shows[showID]
//	if show == nil {
//		return ShowNotFoundError{showID}
//	}
//	if show.Headliner == "" {
//		return HeadlinerAlreadySetError{showID}
//	}
//
//	show.Headliner = headliner
//	return nil
//}
//
//func (r *ShowDetailRepository) AddSlotBookingToShow(showID string, booking readmodel.Booking) error {
//	show := r.shows[showID]
//	if show == nil {
//		return ShowNotFoundError{showID}
//	}
//
//	var slotToModify *readmodel.Slot
//	for _, slot := range show.Slots {
//		if slot.ID == booking.SlotID {
//			slotToModify = &slot
//			break
//		}
//	}
//
//	slotToModify.Booking = &booking
//	for dex, slot := range show.Slots {
//		if slot.ID == booking.SlotID {
//			show.Slots[dex] = *slotToModify
//			break
//		}
//	}
//
//	r.shows[showID] = show
//	return nil
//}
//
//func (r *ShowDetailRepository) GetShowDetail(showID string) (*readmodel.ShowDetail, error) {
//	show := r.shows[showID]
//	if show == nil {
//		return nil, ShowNotFoundError{showID}
//	}
//	return show, nil
//}
//
//func (r *ShowDetailRepository) Clear() {
//	r.shows = make(map[string]*readmodel.ShowDetail)
//}
//
//func (r *ShowDetailRepository) GetSlot(showID string, slotID string) (*readmodel.Slot, error) {
//	show := r.shows[showID]
//	if show == nil {
//		return nil, ShowNotFoundError{showID}
//	}
//	for _, slot := range show.Slots {
//		if slot.ID == slotID {
//			return &slot, nil
//		}
//	}
//	return nil, SlotNotFoundError{showID, slotID}
//}
