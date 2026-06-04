package schemas

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name           string
	Age            int
	HasAnyTaskDone bool
	Tasks          pq.StringArray `gorm:type:text[]`
}

type EmployeeResponse struct {
	Name      string         `json:"name"`
	Age       int            `json:"age"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	Tags      []string       `json:"tags"`
}
