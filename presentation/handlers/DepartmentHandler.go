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
	department := <-ch
	departmentError := department.GetError()

	if departmentError.GetMessage() != "" {
		errorModel := model.NewErrorModel(departmentError.GetMessage())
		response, err := json.Marshal(errorModel)
		if err != nil {
			log.Error(err)
			return
		}
		w.WriteHeader(departmentError.GetCode())
		fmt.Fprintln(w, string(response))
	} else {
		departmentModel = mapper.NewDepartmentModelDataMapper().Convert(department)
		response, err := json.Marshal(departmentModel)
		if err != nil {
			log.Error(err)
			return
		}

		fmt.Fprintln(w, string(response))
	}
}

// GetDepartments return a list of all the deparments
func (d *DepartmentHandler) GetDepartments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	queryValues := r.URL.Query()
	name := queryValues.Get("name")
	if name == "" {
		response, err := getAllDepartments()
		if err != nil {
			log.Error(err)
			return
		}
		fmt.Fprintln(w, string(response))

	} else {
		getDepartmentByName(name, w)
	}
}

func getAllDepartments() ([]byte, error) {
	ch := make(chan []domain.Department, 1)

	go interactor.NewGetAllDepartments().Execute(ch)

	departments := <-ch
	departmentsModel := mapper.NewDepartmentModelDataMapper().ConvertSlice(departments)
	return json.Marshal(departmentsModel)
}

func getDepartmentByName(name string, w http.ResponseWriter) {
	ch := make(chan domain.Department, 1)

	go interactor.NewGetDepartmentByName(name).Execute(ch)

	department := <-ch
	departmentError := department.GetError()

	if departmentError.GetMessage() != "" {
		errorModel := model.NewErrorModel(departmentError.GetMessage())
		response, err := json.Marshal(errorModel)
		if err != nil {
			log.Error(err)
			return
		}
		w.WriteHeader(departmentError.GetCode())
		fmt.Fprintln(w, string(response))
	} else {
		departmentModel := mapper.NewDepartmentModelDataMapper().Convert(department)
		response, err := json.Marshal(departmentModel)
		if err != nil {
			log.Error(err)
			return
		}
		fmt.Fprintln(w, string(response))
	}
}
