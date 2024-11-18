package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (tds *Todos) add(title string) {
	td := Todo{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	*tds = append(*tds, td)
}

func (tds *Todos) validate_index(id int) error {
	if id < 0 || id >= len(*tds) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (tds *Todos) delete(id int) error {
	t := *tds
	if err := t.validate_index(id); err != nil {
		return err
	}
	*tds = append(t[:id], t[id+1:]...)
	return nil
}

func (tds *Todos) toggle(id int) error {
	t := *tds
	if err := t.validate_index(id); err != nil {
		return err
	}
	is_completed := t[id].Completed
	if !is_completed {
		completion_time := time.Now()
		t[id].CompletedAt = &completion_time
	}
	t[id].Completed = !is_completed
	return nil
}

func (tds *Todos) edit(id int, new_title string) error {
	t := *tds
	if err := t.validate_index(id); err != nil {
		return err
	}
	t[id].Title = new_title
	return nil
}

func (tds *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for id, t := range *tds {
		completed := "❌"
		completed_at := ""
		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completed_at = t.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(
			strconv.Itoa(id),
			t.Title,
			completed,
			t.CreatedAt.Format(time.RFC1123),
			completed_at,
		)
	}
	table.Render()
}
