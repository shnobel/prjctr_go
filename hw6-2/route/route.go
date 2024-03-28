package route

import (
	"errors"
	"fmt"
	"prjctr_go/hw6-2/entities"
	"prjctr_go/hw6-2/transport"
)

type Route struct {
	Transports []transport.Transport
}

func (r *Route) Build(ti entities.TripInfo) {
	for _, value := range ti.Path {
		switch value.TransportType {
		case "bus":
			r.AddTransport(&transport.Bus{
				Name:        value.TransportType,
				Origin:      value.From,
				Destination: value.To,
			})
		case "airplane":
			r.AddTransport(&transport.Airplane{
				Name:        value.TransportType,
				Origin:      value.From,
				Destination: value.To,
			})
		case "train":
			r.AddTransport(&transport.Train{
				Name:        value.TransportType,
				Origin:      value.From,
				Destination: value.To,
			})
		}

	}
}

func (r *Route) AddTransport(transport transport.Transport) error {
	if transport == nil {
		return errors.New("please provide transport")
	}
	r.Transports = append(r.Transports, transport)
	return nil
}

func (r *Route) String() {
	for _, t := range r.Transports {
		fmt.Println(t)
	}
}
