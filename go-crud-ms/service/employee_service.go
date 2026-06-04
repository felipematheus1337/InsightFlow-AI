package service

import (
	"time"

	"github.com/felipematheus1337/InsightFlow-AI/go-crud-ms/schemas"
	"gorm.io/gorm"
)

type EmployeeService struct {
	db *gorm.DB
}

func NewEmployeeService(db *gorm.DB) *EmployeeService {
	return &EmployeeService{db}
}

func (e EmployeeService) CreateEmployee(employee *schemas.Employee) (*schemas.EmployeeResponse, error) {

	var response *schemas.EmployeeResponse

	employee.HasAnyTaskDone = len(employee.Tasks) > 0

	if err := e.db.Create(&employee).Error; err != nil {

		return &schemas.EmployeeResponse{}, err
	}

	response = &schemas.EmployeeResponse{
		Name:      employee.Name,
		Age:       employee.Age,
		Tasks:     employee.Tasks,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: employee.DeletedAt,
	}

	return response, nil

}

func (e EmployeeService) listEmployees() ([]*schemas.EmployeeResponse, error) {

	var employees []*schemas.Employee

	if err := e.db.Find(&employees).Error; err != nil {
		return nil, err
	}

	var listEmployeesResponse []*schemas.EmployeeResponse

	for _, emp := range employees {
		listEmployeesResponse = append(listEmployeesResponse, &schemas.EmployeeResponse{
			Name:      emp.Name,
			Age:       emp.Age,
			Tasks:     emp.Tasks,
			CreatedAt: emp.CreatedAt,
			UpdatedAt: emp.UpdatedAt,
			DeletedAt: emp.DeletedAt,
		})
	}

	return listEmployeesResponse, nil

}
