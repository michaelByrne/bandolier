package application

import (
	"bandolier/domain/readmodel"
	"bandolier/domain/venueshow/events"
	"bandolier/infrastructure/inmemory"
	"bandolier/infrastructure/projections"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewShowDetailProjection(t *testing.T) {
	r := inmemory.NewShowDetailRepository()
	showID := "243d76d4-e25f-44ef-9fb1-3cf02d049a48"
	venueID := "ed0daf8f-1c90-48a2-b570-2280671dcad1"
	firstSlotID := "7231bf00-ba64-4324-aa0c-24c92b6cca03"
	secondSlotID := "192cc612-3084-45ff-ab76-8b999bd54d56"
	firstArtistID := "3601dfcb-5d84-4a04-8f38-76e7d24b358f"
	secondArtistID := "469fb0d3-38e0-4a31-b9eb-7e81ae0bff91"
	firstArtistName := "Gwar"
	secondArtistName := "Godflesh"
	firstStartTime := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)
	secondStartTime := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC).Add(time.Minute * 45)
	showDate := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)

	p := &ShowDetailTests{
		ProjectionTests: projections.NewProjectionTests(t, func() projections.Projection {
			return NewShowDetailProjection(r)
		}),
		repository:        r,
		showID:            showID,
		venueID:           venueID,
		firstSlotID:       firstSlotID,
		secondSlotID:      secondSlotID,
		firstArtistID:     firstArtistID,
		secondArtistID:    secondArtistID,
		firstArtistName:   firstArtistName,
		secondArtistName:  secondArtistName,
		firstArtistStart:  firstStartTime,
		secondArtistStart: secondStartTime,
		date:              showDate,
	}

	t.Run("ShouldAddShowDetailToTheList", p.ShouldAddShowDetailToTheList)
	t.Run("ShouldAddASlotBookingToTheShow", p.ShouldAddASlotBookingToTheShow)
	t.Run("ShouldAddHeadlinerToShow", p.ShouldAddHeadlinerToShow)
}

func (d ShowDetailTests) ShouldAddShowDetailToTheList(t *testing.T) {
	d.Given(events.NewShowScheduled(d.showID, d.venueID, d.date))

	detail, err := d.repository.GetShowDetail(d.showID)
	require.NoError(t, err)

	d.Then(readmodel.NewShowDetail(d.showID, "", d.date, d.venueID, []readmodel.Slot{}), detail)
}

func (d ShowDetailTests) ShouldAddASlotBookingToTheShow(t *testing.T) {
	d.Given(
		events.NewShowScheduled(d.showID, d.venueID, d.date),
		events.NewSlotScheduled(d.firstSlotID, d.showID, d.firstArtistStart, time.Minute*45),
		events.NewSlotBooked(d.firstSlotID, d.showID, d.firstArtistID, d.firstArtistName, false),
	)

	detail, err := d.repository.GetShowDetail(d.showID)
	require.NoError(t, err)

	d.Then(readmodel.NewBooking(d.firstSlotID, d.firstArtistID, d.firstArtistName, d.firstArtistStart, time.Minute*45), detail.Slots[0].Booking)
}

func (d ShowDetailTests) ShouldAddHeadlinerToShow(t *testing.T) {
	d.Given(
		events.NewShowScheduled(d.showID, d.venueID, d.date),
		events.NewSlotScheduled(d.firstSlotID, d.showID, d.firstArtistStart, time.Minute*45),
		events.NewSlotBooked(d.firstSlotID, d.showID, d.firstArtistID, d.firstArtistName, true),
	)

	detail, err := d.repository.GetShowDetail(d.showID)
	require.NoError(t, err)

	d.Then(d.firstArtistName, detail.Headliner)
}

type ShowDetailTests struct {
	projections.ProjectionTests

	repository        readmodel.ShowDetailRepository
	showID            string
	firstSlotID       string
	secondSlotID      string
	firstArtistID     string
	secondArtistID    string
	firstArtistName   string
	secondArtistName  string
	venueID           string
	date              time.Time
	firstArtistStart  time.Time
	secondArtistStart time.Time
}
