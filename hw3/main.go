package main

import (
	"errors"
	"fmt"
	"os"
	"prjctr_go/utils"
)

const DASH = "-----------------------------------------------------------------------------------------------"

type Quest struct {
	Name        string   `json:"name"`
	Goal        string   `json:"goal"`
	Description string   `json:"desc"`
	Dialogs     []Dialog `json:"dialogs"`
	Tasks       []Task   `json:"tasks"`
}

func (q *Quest) GetDialogBy(id int) Dialog {
	var dialog Dialog
	for i := range q.Dialogs {
		if q.Dialogs[i].ID == id {
			return q.Dialogs[i]
		}
	}
	return dialog
}

func (q *Quest) GetTaskBy(id int) Task {
	var task Task
	for i := range q.Tasks {
		if q.Tasks[i].ID == id {
			task = q.Tasks[i]
			break
		}
	}
	return task
}

func (q *Quest) Start() error {
	if len(q.Dialogs) == 0 {
		return errors.New("empty dialogs")
	}

	startId := 1
	dialog := q.GetDialogBy(startId)
	nextDialogId, nextTaskId := dialog.Proceed()

	for {
		if nextTaskId > 0 {
			task := q.GetTaskBy(nextTaskId)
			nextDialogId = task.Proceed()
		}

		if nextDialogId > 0 {
			dialog = q.GetDialogBy(nextDialogId)
			if dialog.Last {
				dialog.PrintInfo()
				break
			}
			nextDialogId, nextTaskId = dialog.Proceed()
		}
	}

	return nil
}

type Dialog struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	OwnerName string   `json:"owner_name"`
	Text      string   `json:"text"`
	Options   []Option `json:"options"`
	Last      bool     `json:"last"`
}

func (d *Dialog) Proceed() (dialogId, taskId int) {
	d.PrintInfo()
	result := d.GetUserInput()
	optionId := result - 1
	return d.GetNext(optionId)
}

func (d *Dialog) PrintInfo() {
	fmt.Printf("%v\n", DASH)
	fmt.Printf("%50v\n", d.Name)
	if d.OwnerName != "" {
		fmt.Printf("%10v: ", d.OwnerName)
	}

	if d.Text != "" {
		fmt.Printf("%v\n", d.Text)
	}

	fmt.Println()
	for i, option := range d.Options {
		option.PrintInfo(i)
	}
}

func (d *Dialog) GetUserInput() int {
	var result int
	fmt.Println()

	for {
		fmt.Print("Ваш ответ: ")
		fmt.Scanln(&result)
		if result > len(d.Options) || result < 1 {
			fmt.Println("Incorrect answer, try again!")
			continue
		}
		break
	}

	fmt.Println()

	return result
}

func (d *Dialog) GetNext(id int) (dialogId, taskId int) {
	option := d.Options[id]
	switch {
	case option.Answer != "":
		fmt.Println(option.Answer)
	case option.NextTask > 0:
		return -1, option.NextTask
	}
	return option.NextDialog, -1
}

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	NextDialog  int    `json:"next_dialog"`
}

func (t *Task) Proceed() (dialogId int) {
	t.PrintInfo()
	return t.NextDialog
}

func (t *Task) PrintInfo() {
	fmt.Printf("%v\n", DASH)
	fmt.Printf("%50v\n", t.Name)
	fmt.Printf("%v\n", t.Description)
}

type Option struct {
	Text       string `json:"text"`
	NextDialog int    `json:"next_dialog"`
	NextTask   int    `json:"next_task"`
	Answer     string `json:"answer"`
}

func (o *Option) PrintInfo(index int) {
	fmt.Printf("%v: %v\n", index+1, o.Text)
}

func main() {

	var quest *Quest

	quest, err := utils.GetDataFromJson[Quest]("quest.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Printf("%100v\n", quest.Name)
	fmt.Printf("Задание: %v\n", quest.Goal)
	fmt.Printf("Описание: %v\n", quest.Description)

	err = quest.Start()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
