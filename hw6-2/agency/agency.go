package agency

import (
	"errors"
	"fmt"
	"prjctr_go/hw6-2/entities"
	"prjctr_go/hw6-2/route"
	"prjctr_go/utils"
)

type travelAgency struct {
	route          route.Route
	availableTrips entities.Trips
}

func NewTravelAgency() (*travelAgency, error) {
	trips, err := utils.GetDataFromJson[entities.Trips]("trips.json")
	if err != nil {
		return nil, err
	}

	ta := travelAgency{
		availableTrips: *trips,
	}

	return &ta, nil
}

func (ta *travelAgency) SendTo(direction string, passengers ...*entities.Passenger) error {
	tripInfo := ta.availableTrips.GetTripInfo(direction)
	if tripInfo == nil {
		return errors.New("sorry, no available routes")
	}

	ta.route.Build(*tripInfo)
	ta.send(passengers...)
	return nil
}

func (ta *travelAgency) send(passengers ...*entities.Passenger) {
	for _, t := range ta.route.Transports {
		t.PickUpPassengers(passengers...)
		fmt.Print(t)
		t.DropOffPassengers()
	}
}
