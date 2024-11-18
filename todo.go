package main

import (
	"errors"
	"fmt"
	"time"
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
