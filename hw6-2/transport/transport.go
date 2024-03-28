package transport

import (
	"fmt"
	"prjctr_go/hw6-2/entities"
)

type Transport interface {
	PickUpPassengers(newPassengers ...*entities.Passenger)
	DropOffPassengers()
}

type Bus struct {
	Name        string
	Origin      string
	Destination string
	Passengers  []*entities.Passenger
}

func (b Bus) String() string {
	var result string = ""
	for _, p := range b.Passengers {
		result += fmt.Sprintf("Passenger %s heading to %s from %s by %s\n", p.Name, b.Destination, b.Origin, b.Name)
	}
	return result
}

func (b *Bus) PickUpPassengers(newPassengers ...*entities.Passenger) {
	if newPassengers != nil {
		b.Passengers = append(b.Passengers, newPassengers...)
		for _, p := range newPassengers {
			fmt.Printf("%s in the bus\n", p.Name)
		}
	} else {
		fmt.Println("No passengers to pick up")
	}
}

func (b *Bus) DropOffPassengers() {
	for _, p := range b.Passengers {
		fmt.Printf("%s out of the bus\n", p.Name)
	}
	b.Passengers = nil
}

type Train struct {
	Name        string
	Origin      string
	Destination string
	Passengers  []*entities.Passenger
}

func (t Train) String() string {
	var result string = ""
	for _, p := range t.Passengers {
		result += fmt.Sprintf("Passenger %s heading to %s from %s by %s\n", p.Name, t.Destination, t.Origin, t.Name)
	}
	return result
}

func (t *Train) PickUpPassengers(newPassengers ...*entities.Passenger) {
	if newPassengers != nil {
		t.Passengers = append(t.Passengers, newPassengers...)
		for _, p := range newPassengers {
			fmt.Printf("%s in the train\n", p.Name)
		}
	} else {
		fmt.Println("No passengers to pick up")
	}
}

func (t *Train) DropOffPassengers() {
	for _, p := range t.Passengers {
		fmt.Printf("%s out of the train\n", p.Name)
	}
	t.Passengers = nil
}

type Airplane struct {
	Name        string
	Origin      string
	Destination string
	Passengers  []*entities.Passenger
}

func (a Airplane) String() string {
	var result string = ""
	for _, p := range a.Passengers {
		result += fmt.Sprintf("Passenger %s heading to %s from %s by %s\n", p.Name, a.Destination, a.Origin, a.Name)
	}
	return result
}

func (a *Airplane) PickUpPassengers(newPassengers ...*entities.Passenger) {
	if newPassengers != nil {
		a.Passengers = append(a.Passengers, newPassengers...)
		for _, p := range newPassengers {
			fmt.Printf("%s in the Airplane\n", p.Name)
		}
	} else {
		fmt.Println("No passengers to pick up")
	}
}

func (a *Airplane) DropOffPassengers() {
	for _, p := range a.Passengers {
		fmt.Printf("%s out of the Airplane\n", p.Name)
	}
	a.Passengers = nil
}
