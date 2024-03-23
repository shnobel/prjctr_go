package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

var (
	nonWordSymbolRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
)

type editor struct {
	fileName string
	content  []string
	words    map[string][]int
}

func NewEditor(fileName string) editor {
	editor := editor{
		fileName: fileName,
	}

	err := editor.readContent()
	checkError(err)
	err = editor.createIndex()
	checkError(err)

	return editor
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

func (e *editor) createIndex() error {
	splitRow := func(row string) []string {
		row = strings.TrimSuffix(row, " ")
		return strings.Split(row, " ")
	}

	cleanString := func(str string) string {
		str = nonWordSymbolRegex.ReplaceAllString(str, "")
		return strings.ToLower(str)
	}

	e.words = make(map[string][]int)

	for rowIndex, row := range e.content {
		rowWords := splitRow(row)
		for _, str := range rowWords {
			str = cleanString(str)
			if !slices.Contains(e.words[str], rowIndex) {
				e.words[str] = append(e.words[str], rowIndex)
			}
		}
	}
	return nil
}

func (e *editor) getRowsSlow(text string) ([]string, error) {
	defer duration(track("getRowsSlow"))
	result := []string{}
	for i, row := range e.content {
		if strings.Contains(row, text) {
			str := formatIndex(i) + row
			result = append(result, str)
		}
	}

	if len(result) == 0 {
		return nil, errors.New("no results")
	}

	return result, nil
}

func (e *editor) getRowsFast(word string) ([]string, error) {
	defer duration(track("getRowsFast"))
	rows, ok := e.words[word]
	if !ok {
		return nil, errors.New("no results")
	}

	result := []string{}
	for _, rowIndex := range rows {
		str := formatIndex(rowIndex) + e.content[rowIndex]
		result = append(result, str)
	}
	return result, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func formatIndex(i int) string {
	return "line " + strconv.Itoa(i+1) + ": "
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
