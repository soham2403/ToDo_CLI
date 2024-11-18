package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type todo struct {
	title        string
	completed    bool
	created_at   time.Time
	completed_at *time.Time
}

type todos []todo

func (tds *todos) add(title string) {
	td := todo{
		title:        title,
		completed:    false,
		created_at:   time.Now(),
		completed_at: nil,
	}
	*tds = append(*tds, td)
}

func (tds *todos) validate_index(id int) error {
	if id < 0 || id > len(*tds) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (tds *todos) delete(id int) error {
	t := *tds
	if err := t.validate_index(id); err != nil {
		return err
	}

	*tds = append(t[:id], t[id+1:]...)
	return nil
}

func (tds *todos) toogle(id int) error {
	t := *tds
	if err := t.validate_index(id); err != nil {
		return err
	}

	is_completed := t[id].completed

	if !is_completed {
		completion_time := time.Now()
		t[id].completed_at = &completion_time
	}
	t[id].completed = !is_completed
	return nil
}

func (tds *todos) update(id int, new_title string) error {
	t := *tds
	if err := t.validate_index(id); err != nil {
		return err
	}

	t[id].title = new_title
	return nil
}

func (tds *todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for id, t := range *tds {
		completed := "❌"
		completed_at := ""

		if t.completed {
			completed = "✅"
			if t.completed_at != nil {
				completed_at = t.completed_at.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(id), t.title, completed, t.created_at.Format(time.RFC1123), completed_at)
	}
	table.Render()
}
