package mapper

import (
	"github.com/felipematheus1337/InsightFlow-AI/go-crud-ms/dto"
	"github.com/felipematheus1337/InsightFlow-AI/go-crud-ms/schemas"
)

func CreateToSchema(dto dto.CreateEmployeeDTO) *schemas.Employee {
	return &schemas.Employee{
		Name:  dto.Name,
		Age:   dto.Age,
		Tasks: dto.Tasks,
	}
}
