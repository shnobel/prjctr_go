package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type editor struct {
	fileName string
	content  []string
}

func (e *editor) readContent() error {
	file, err := os.Open(e.fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		e.content = append(e.content, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (e *editor) getRows(text string) []string {

	formatIndex := func(i int) string {
		return "line " + strconv.Itoa(i+1) + ": "
	}

	result := []string{}
	for i, row := range e.content {
		if strings.Contains(row, text) {
			str := formatIndex(i) + row
			result = append(result, str)
		}
	}

	return result
}

func (e *editor) find(text string) error {
	rows := e.getRows(text)
	if len(rows) == 0 {
		return errors.New("no results")
	}
	for _, row := range rows {
		fmt.Println(row)
	}

	return nil
}
