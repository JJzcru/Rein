package handlers

import (
	"Rein/domain/interactor"
	"Rein/presentation/mapper"
	"Rein/presentation/model"
	"encoding/json"
	"fmt"
	"net/http"

	"Rein/domain"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
)

// DepartmentHandler Handles the request for departments
type DepartmentHandler struct{}

type DepartmentsModel []*model.DepartmentModel

// NewDepartmentHandler Constructor to create a department handler
func NewDepartmentHandler() *DepartmentHandler {
	return &DepartmentHandler{}
}

// CreateDepartment Add a new department to the system
func (d *DepartmentHandler) CreateDepartment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	departmentModel := model.NewDepartmentModel()
	departmentModel.SetName("ventas")

	ch := make(chan domain.Department, 1)

	go interactor.NewCreateDepartment(mapper.NewDepartmentModelDataMapper().Transform(departmentModel)).Execute(ch)

	departmentModel = mapper.NewDepartmentModelDataMapper().Convert(<-ch)
	response, err := json.Marshal(departmentModel)
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Fprintln(w, string(response))
}

// GetDepartments Add all the departments
func (d *DepartmentHandler) GetDepartments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	ch := make(chan []domain.Department, 1)

	go interactor.NewGetAllDepartments().Execute(ch)
	departmentsModel := mapper.NewDepartmentModelDataMapper().ConvertSlice(<-ch)
	response, err := json.Marshal(departmentsModel)
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Fprintln(w, string(response))
}
