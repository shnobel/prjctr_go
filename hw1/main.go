package main

import (
	"fmt"
	"time"
)

type Person struct {
	fullName        string
	birthDate       time.Time
	where           Location
	achievement     string
	achievementYear time.Time
	career          string
}

type Location struct {
	country string
	city    string
}

func (p Person) PrintBio() {
	fmt.Printf(`
	%30q

	Birth: %v, %s, %s
	
	Claim to Fame: %s (%v)
	
	Philosophy: Open-source advocate

	Career: %s
	
	`, p.fullName, p.birthDate.Format("January 2, 2006"), p.where.city, p.where.country, p.achievement, p.achievementYear.Year(), p.career)
}

func main() {
	linus := Person{
		fullName:        "Linus Torvalds",
		birthDate:       time.Date(1969, time.December, 28, 0, 0, 0, 0, time.UTC),
		where:           Location{country: "Finland", city: "Helsinki"},
		achievement:     "Creator of the Linux kernel",
		achievementYear: time.Date(1991, 0, 0, 0, 0, 0, 0, time.UTC),
		career: `Despite Linux's widespread success, Linus Torvalds maintained a humble approach.
	He continued to oversee the development of the Linux kernel and became an influential figure in the open-source community. 
	In 2005, he was inducted into the Hall of Fellows of the Computer History Museum for his creation of the Linux kernel and his leadership in open-source development.`,
	}

	linus.PrintBio()
}
