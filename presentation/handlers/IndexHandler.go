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
func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userModel := model.NewUserModel()
	userModel.SetName("Tracer 123")
	userModel.SetActive(false)

	userModel.SetRole([]string{"admin"})
	response, err := json.Marshal(userModel)

	if err != nil {
		log.Error(err)
		return
	}
	log.Info(userModel.ToString())
	fmt.Fprintln(w, string(response))
}
