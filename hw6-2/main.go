package main

import (
	"fmt"
	"os"
	"prjctr_go/hw6-2/agency"
	"prjctr_go/hw6-2/entities"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func main() {
	ta, err := agency.NewTravelAgency()
	checkError(err)
	passenger1 := entities.Passenger{Name: "passenger1"}
	passenger2 := entities.Passenger{Name: "passenger2"}
	passenger3 := entities.Passenger{Name: "passenger3"}
	passenger4 := entities.Passenger{Name: "passenger4"}

	err = ta.SendTo("New York", &passenger1, &passenger2, &passenger3)
	checkError(err)
	err = ta.SendTo("Anchorage", &passenger2)
	checkError(err)
	err = ta.SendTo("Washington", &passenger3)
	checkError(err)
	err = ta.SendTo("Bucharest", &passenger4)
	checkError(err)
}
