package handlers

import (
	"Rein/presentation/model"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
)

// IndexHandler Control the health check of the application
func IndexHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	userModel := model.NewUserModel()
	userModel.SetName("Gon")
	userModel.SetLastName("Lux")
	response, err := json.Marshal(userModel)

	if err != nil {
		log.Error(err)
		return
	}
	log.Info(userModel.ToString())
	fmt.Fprintln(res, string(response))
}
