package venueshow

import (
	"bandolier/domain/venueshow/commands"
	"bandolier/domain/venueshow/events"
	"bandolier/eventsourcing"
	"bandolier/infrastructure"
	"reflect"
	"time"
)

const prefix = "venueshow"

func RegisterTypes(tm *eventsourcing.TypeMapper) error {
	mustParseTime := func(s string) time.Time {
		t, _ := time.Parse(time.RFC3339, s)
		return t
	}
	mustParseDuration := func(s string) time.Duration {
		d, _ := time.ParseDuration(s)
		return d
	}

	err := tm.MapEvent(reflect.TypeOf(events.ShowScheduled{}), prefix+"-show-scheduled", func(data map[string]interface{}) interface{} {
		return events.ShowScheduled{
			ID:      data["id"].(string),
			VenueID: data["venue_id"].(string),
			Date:    mustParseTime(data["date"].(string)),
		}
	}, func(t interface{}) map[string]interface{} {
		event := t.(events.ShowScheduled)
		return map[string]interface{}{
			"id":       event.ID,
			"venue_id": event.VenueID,
			"date":     event.Date,
		}
	})
	if err != nil {
		return err
	}

	err = tm.MapEvent(reflect.TypeOf(events.ShowScheduleCancelled{}), prefix+"-show-cancelled", func(data map[string]interface{}) interface{} {
		return events.ShowScheduleCancelled{
			ShowID: data["id"].(string),
		}
	}, func(t interface{}) map[string]interface{} {
		event := t.(events.ShowScheduleCancelled)
		return map[string]interface{}{
			"id": event.ShowID,
		}
	})
	if err != nil {
		return err
	}

	err = tm.MapEvent(reflect.TypeOf(events.ShowArchived{}), prefix+"-show-archived", func(data map[string]interface{}) interface{} {
		return events.ShowArchived{
			ShowID:            data["id"].(string),
			DoorAmountInCents: int(data["door_amount_in_cents"].(float64)),
		}
	}, func(t interface{}) map[string]interface{} {
		event := t.(events.ShowArchived)
		return map[string]interface{}{
			"id":                   event.ShowID,
			"door_amount_in_cents": event.DoorAmountInCents,
		}
	})
	if err != nil {
		return err
	}

	err = tm.MapEvent(reflect.TypeOf(events.SlotScheduled{}), prefix+"-slot-scheduled", func(data map[string]interface{}) interface{} {
		return events.SlotScheduled{
			ID:        data["id"].(string),
			ShowID:    data["show_id"].(string),
			StartTime: mustParseTime(data["date"].(string)),
			Duration:  mustParseDuration(data["duration"].(string)),
		}
	}, func(t interface{}) map[string]interface{} {
		event := t.(events.SlotScheduled)
		return map[string]interface{}{
			"id":       event.ID,
			"show_id":  event.ShowID,
			"date":     event.StartTime,
			"duration": event.Duration.String(),
		}
	})
	if err != nil {
		return err
	}

	err = tm.MapEvent(reflect.TypeOf(events.SlotBooked{}), prefix+"-slot-booked", func(data map[string]interface{}) interface{} {
		return events.SlotBooked{
			ID:         data["id"].(string),
			ShowID:     data["show_id"].(string),
			ArtistName: data["artist_name"].(string),
			ArtistID:   data["artist_id"].(string),
			Headliner:  data["headliner"].(bool),
		}
	}, func(t interface{}) map[string]interface{} {
		event := t.(events.SlotBooked)
		return map[string]interface{}{
			"id":          event.ID,
			"show_id":     event.ShowID,
			"artist_name": event.ArtistName,
			"artist_id":   event.ArtistID,
			"headliner":   event.Headliner,
		}
	})
	if err != nil {
		return err
	}

	err = tm.MapEvent(reflect.TypeOf(events.SlotBookingCancelled{}), prefix+"-slot-booking-cancelled", func(data map[string]interface{}) interface{} {
		return events.SlotBookingCancelled{
			ID: data["id"].(string),
		}
	}, func(t interface{}) map[string]interface{} {
		event := t.(events.SlotBookingCancelled)
		return map[string]interface{}{
			"id": event.ID,
		}
	})
	if err != nil {
		return err
	}

	err = tm.MapEvent(reflect.TypeOf(events.SlotScheduleCancelled{}), prefix+"-slot-schedule-cancelled", func(data map[string]interface{}) interface{} {
		return events.SlotScheduleCancelled{
			ID: data["id"].(string),
		}
	}, func(t interface{}) map[string]interface{} {
		event := t.(events.SlotScheduleCancelled)
		return map[string]interface{}{
			"id": event.ID,
		}
	})
	if err != nil {
		return err
	}

	registerType := func(t interface{}, typeName string) {
		tm.RegisterType(infrastructure.GetValueType(t), typeName, func() interface{} {
			return t
		})
	}

	registerType(commands.ScheduleShow{}, prefix+"-schedule-show")
	registerType(commands.ScheduleSlot{}, prefix+"-schedule-slot")
	registerType(commands.CancelSlotBooking{}, prefix+"-cancel-slot-booking")
	registerType(commands.BookSlot{}, prefix+"-book-slot")
	registerType(commands.ArchiveShow{}, prefix+"-archive-show")
	registerType(commands.Cancel{}, prefix+"-cancel")

	return nil
}
