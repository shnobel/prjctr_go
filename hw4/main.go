package main

import (
	"fmt"
	"os"
)

func main() {
	//HW-4.1
	editor := editor{
		fileName: "text.txt",
	}

	err := editor.readContent()
	checkError(err)

	text := "duplicate"
	err = editor.find(text)
	checkError(err)

	fmt.Println()
	fmt.Println()

	//HW-4.2
	test1 := []person{
		{id: 1},
		{id: 1},
		{id: 2},
	}

	test2 := []person{
		{id: 0},
		{id: 0},
		{id: 1},
		{id: 1},
		{id: 1},
		{id: 2},
		{id: 2},
		{id: 3},
		{id: 3},
		{id: 4},
	}

	test3 := []person{
		{id: 4},
		{id: 3},
		{id: 3},
		{id: 2},
		{id: 2},
		{id: 1},
		{id: 1},
		{id: 1},
		{id: 0},
		{id: 0},
	}

	test4 := []person{}

	t1Result, err := distinct(test1)
	checkError(err)
	fmt.Println(t1Result)

	t2Result, err := distinct(test2)
	checkError(err)
	fmt.Println(t2Result)

	t3Result, err := distinct(test3)
	checkError(err)
	fmt.Println(t3Result)

	t4Result, err := distinct(test4)
	checkError(err)
	fmt.Println(t4Result)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
