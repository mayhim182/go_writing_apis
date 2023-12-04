package todo

// What we wanna do
// Add new tasks
// Mark tasks as complete
// Delete tasks from the list of tasks
// Save the list of tasks as JSON
// Get the tasks from the JSON file

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

// Add function to add new tasks to the list []item
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

// Complete function for setting the done field as true and completed to the current time then
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

// Save function saves the list of tasks in JSON format
func (l *List) Save(fileName string) error {
	json, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, json, 0644)
}

// Get function this function will get the saved tasks list from the directory with the help of the fileName
// and decode and parse that JSON data into a list

func (l *List) Get(fileName string) error {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		//If the given file does not exist
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}
	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}
