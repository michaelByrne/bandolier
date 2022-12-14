package showartist

import (
	"bandolier/domain/showartist/commands"
	"bandolier/domain/showartist/events"
	"bandolier/eventsourcing"
	"bandolier/infrastructure"
	"reflect"
)

const prefix = "showartist"

func RegisterTypes(tm *eventsourcing.TypeMapper) error {
	//mustParseTime := func(s string) time.Time {
	//	t, _ := time.Parse(time.RFC3339, s)
	//	return t
	//}
	//mustParseDuration := func(s string) time.Duration {
	//	d, _ := time.ParseDuration(s)
	//	return d
	//}

	err := tm.MapEvent(reflect.TypeOf(events.ArtistCreated{}), prefix+"-artist-created", func(data map[string]interface{}) interface{} {
		return events.ArtistCreated{
			ID:   data["id"].(string),
			Name: data["name"].(string),
		}
	}, func(t interface{}) map[string]interface{} {
		event := t.(events.ArtistCreated)
		return map[string]interface{}{
			"id":   event.ID,
			"name": event.Name,
		}
	})
	if err != nil {
		return err
	}

	err = tm.MapEvent(reflect.TypeOf(events.Paid{}), prefix+"-artist-paid", func(data map[string]interface{}) interface{} {
		return events.Paid{
			ShowID:        data["show_id"].(string),
			ArtistID:      data["artist_id"].(string),
			AmountInCents: int(data["amount_in_cents"].(float64)),
		}
	}, func(t interface{}) map[string]interface{} {
		event := t.(events.Paid)
		return map[string]interface{}{
			"show_id":         event.ShowID,
			"artist_id":       event.ArtistID,
			"amount_in_cents": event.AmountInCents,
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

	registerType(commands.CreateArtist{}, prefix+"-create-artist")
	registerType(commands.PayArtist{}, prefix+"-pay-artist")

	return nil
}
