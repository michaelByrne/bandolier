package showbank

import (
	"bandolier/domain/showbank/commands"
	"bandolier/domain/showbank/events"
	"bandolier/eventsourcing"
	"bandolier/infrastructure"
	"reflect"
)

const prefix = "showbank"

func RegisterTypes(tm *eventsourcing.TypeMapper) error {
	//mustParseTime := func(s string) time.Time {
	//	t, _ := time.Parse(time.RFC3339, s)
	//	return t
	//}
	//mustParseDuration := func(s string) time.Duration {
	//	d, _ := time.ParseDuration(s)
	//	return d
	//}

	err := tm.MapEvent(reflect.TypeOf(events.BankOpened{}), prefix+"-bank-opened", func(data map[string]interface{}) interface{} {
		return events.BankOpened{
			ShowID:         data["show_id"].(string),
			PresaleInCents: int(data["presale_in_cents"].(float64)),
		}
	}, func(t interface{}) map[string]interface{} {
		event := t.(events.BankOpened)
		return map[string]interface{}{
			"show_id":          event.ShowID,
			"presale_in_cents": event.PresaleInCents,
		}
	})
	if err != nil {
		return err
	}

	err = tm.MapEvent(reflect.TypeOf(events.CoversReceived{}), prefix+"-covers-received", func(data map[string]interface{}) interface{} {
		return events.CoversReceived{
			ShowID:        data["show_id"].(string),
			AmountInCents: int(data["amount_in_cents"].(float64)),
			NewBalance:    int(data["new_balance"].(float64)),
		}
	}, func(t interface{}) map[string]interface{} {
		event := t.(events.CoversReceived)
		return map[string]interface{}{
			"show_id":         event.ShowID,
			"amount_in_cents": event.AmountInCents,
			"new_balance":     event.NewBalance,
		}
	})
	if err != nil {
		return err
	}

	err = tm.MapEvent(reflect.TypeOf(events.DoorPaid{}), prefix+"-door-paid", func(data map[string]interface{}) interface{} {
		return events.DoorPaid{
			ShowID:         data["show_id"].(string),
			AmountInCents:  int(data["amount_in_cents"].(float64)),
			BalanceInCents: int(data["new_balance"].(float64)),
		}
	}, func(t interface{}) map[string]interface{} {
		event := t.(events.DoorPaid)
		return map[string]interface{}{
			"show_id":         event.ShowID,
			"amount_in_cents": event.AmountInCents,
			"new_balance":     event.BalanceInCents,
		}
	})
	if err != nil {
		return err
	}

	err = tm.MapEvent(reflect.TypeOf(events.CoversReceived{}), prefix+"-covers-received", func(data map[string]interface{}) interface{} {
		return events.CoversReceived{
			ShowID:        data["show_id"].(string),
			AmountInCents: int(data["amount_in_cents"].(float64)),
			NewBalance:    int(data["new_balance"].(float64)),
		}
	}, func(t interface{}) map[string]interface{} {
		event := t.(events.CoversReceived)
		return map[string]interface{}{
			"show_id":         event.ShowID,
			"amount_in_cents": event.AmountInCents,
			"new_balance":     event.NewBalance,
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

	registerType(commands.OpenBank{}, prefix+"-open-bank")
	registerType(commands.ReceiveCovers{}, prefix+"-receive-covers")
	registerType(commands.PayDoor{}, prefix+"-pay-door")

	return nil
}
