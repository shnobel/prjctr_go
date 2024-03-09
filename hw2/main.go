package main

import (
	"fmt"
	"math/rand"
	"time"
)

const GUESSNUMBER = 1

type Zoo struct {
	keeper  ZooKeeper
	animals []Animal
	cage    Cage
}

func (z *Zoo) Run() {
	fmt.Println("Zoo manager: Welcome to our zoo!")
	time.Sleep(1 * time.Second)
	fmt.Println("Zoo manager: Today we have next animals: ")
	time.Sleep(1 * time.Second)

	for index := range z.animals {
		z.animals[index].PrintInfo()
		time.Sleep(1 * time.Second)
		z.animals[index].Release()
	}

	fmt.Println("Zoo manager: ALARM! ALARM! ALARM!")
	time.Sleep(1 * time.Second)
	fmt.Println("Zoo manager: All animals outside of the cage, zoo keeper!!!")
	time.Sleep(1 * time.Second)
	fmt.Println("Zoo manager: Catch them!!!")
	time.Sleep(1 * time.Second)
	fmt.Printf("Zoo keeper %v: Got you boss!\n", z.keeper.name)

	for index := range z.animals {
		z.keeper.Catch(&z.animals[index], &z.cage)
	}

	fmt.Println("Zoo manager: Good job!")
}

type ZooKeeper struct {
	name string
}

func (z *ZooKeeper) Catch(a *Animal, cage *Cage) {
	fmt.Printf("Zoo keeper %v: Who is here, looks like I see: %v %c\n", z.name, a.name, a.picture)
	for index := 0; !a.inCage; {
		time.Sleep(1 * time.Second)
		guess := rand.Intn(3)
		if guess == GUESSNUMBER {
			a.inCage = true
			cage.animals[index] = *a
			fmt.Printf("Got you. %v %c in the cage\n", a.name, a.picture)
		} else {
			fmt.Println("Missed it!")
		}
	}
}

type Animal struct {
	name    string
	picture rune
	inCage  bool
}

func (a *Animal) PrintInfo() {
	fmt.Printf("Hi, I am: %v %c\n", a.name, a.picture)
}

func (a *Animal) Release() {
	a.inCage = false
}

type Cage struct {
	animals []Animal
}

func main() {
	animalsCount := 10
	animalsPool, err := getRandomAnimals(animalsCount)

	if err != nil {
		fmt.Println("Game Over")
		return
	}

	zoo := Zoo{
		animals: animalsPool,
		keeper: ZooKeeper{
			name: "Rojo",
		},
		cage: Cage{
			animals: make([]Animal, animalsCount),
		},
	}

	zoo.Run()
}

func getRandomAnimals(count int) ([]Animal, error) {
	animals := map[string]rune{
		"RAT":            '\U0001F400',
		"MOUSE":          '\U0001F401',
		"OX":             '\U0001F402',
		"WATER BUFFALO":  '\U0001F403',
		"COW":            '\U0001F404',
		"TIGER":          '\U0001F405',
		"LEOPARD":        '\U0001F406',
		"RABBIT":         '\U0001F407',
		"CAT":            '\U0001F408',
		"DRAGON":         '\U0001F409',
		"CROCODILE":      '\U0001F40A',
		"WHALE":          '\U0001F40B',
		"SNAIL":          '\U0001F40C',
		"SNAKE":          '\U0001F40D',
		"HORSE":          '\U0001F40E',
		"RAM":            '\U0001F40F',
		"GOAT":           '\U0001F410',
		"SHEEP":          '\U0001F411',
		"MONKEY":         '\U0001F412',
		"ROOSTER":        '\U0001F413',
		"CHICKEN":        '\U0001F414',
		"DOG":            '\U0001F415',
		"PIG":            '\U0001F416',
		"BOAR":           '\U0001F417',
		"ELEPHANT":       '\U0001F418',
		"OCTOPUS":        '\U0001F419',
		"SPIRAL SHELL":   '\U0001F41A',
		"BUG":            '\U0001F41B',
		"ANT":            '\U0001F41C',
		"HONEYBEE":       '\U0001F41D',
		"LADY BEETLE":    '\U0001F41E',
		"FISH":           '\U0001F41F',
		"TURTLE":         '\U0001F422',
		"HATCHING CHICK": '\U0001F423',
		"PENGUIN":        '\U0001F427',
	}

	animals_keys := []string{
		"RAT",
		"MOUSE",
		"OX",
		"WATER BUFFALO",
		"COW",
		"TIGER",
		"LEOPARD",
		"RABBIT",
		"CAT",
		"DRAGON",
		"CROCODILE",
		"WHALE",
		"SNAIL",
		"SNAKE",
		"HORSE",
		"RAM",
		"GOAT",
		"SHEEP",
		"MONKEY",
		"ROOSTER",
		"CHICKEN",
		"DOG",
		"PIG",
		"BOAR",
		"ELEPHANT",
		"OCTOPUS",
		"SPIRAL SHELL",
		"BUG",
		"ANT",
		"HONEYBEE",
		"LADY BEETLE",
		"FISH",
		"TURTLE",
		"HATCHING CHICK",
		"PENGUIN",
	}

	if count >= len(animals) || count < 1 {
		return nil, fmt.Errorf("count out of range: %d", count)
	}

	result := make([]Animal, count)

	for i := 0; i < count; i++ {
		index := rand.Intn(len(animals_keys))

		aName := animals_keys[index]
		aPicture := animals[aName]

		result[i] = Animal{
			name:    aName,
			picture: aPicture,
			inCage:  true,
		}
	}
	return result, nil
}
