package handler

import (
	"net/http"
	"strings"

	"github.com/felipematheus1337/InsightFlow-AI/go-crud-ms/dto"
	"github.com/felipematheus1337/InsightFlow-AI/go-crud-ms/mapper"
	"github.com/felipematheus1337/InsightFlow-AI/go-crud-ms/service"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	service *service.EmployeeService
}

func NewEmployeeHandler(service *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service}
}

func (h *EmployeeHandler) CreateEmployee(ctx *gin.Context) {

	var employeeDTO dto.CreateEmployeeDTO

	if err := ctx.ShouldBind(&employeeDTO); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	errValidate := employeeDTO.Validate()

	if errValidate != nil {
		msgErro := errValidate.Error()
		errStrings := strings.Split(msgErro, ",")
		ErrParamIsRequired(errStrings[0], errStrings[1])
	}

	employee := mapper.CreateToSchema(employeeDTO)

	response, err := h.service.CreateEmployee(employee)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "create-employee", response, http.StatusCreated)
}
