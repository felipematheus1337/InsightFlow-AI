package dto

import (
	"errors"
	"fmt"

	"github.com/lib/pq"
)

type CreateEmployeeDTO struct {
	Name  string         `json:"name"`
	Age   int            `json:"age"`
	Tasks pq.StringArray `json:"tasks"`
}

func (e *CreateEmployeeDTO) Validate() error {
	if e.Name == "" {
		return errors.New(fmt.Sprintf("'name', 'string'"))
	}

	if e.Age <= 18 {
		return errors.New(fmt.Sprintf("'age', 'int'"))
	}

	if e.Tasks == nil {
		return errors.New(fmt.Sprintf("'tasks', '[]'"))
	}

	return nil
}
