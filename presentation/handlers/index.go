package handlers
import (
	"fmt"
	"net/http"
	log "github.com/Sirupsen/logrus"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
}
