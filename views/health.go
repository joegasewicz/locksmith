package views

import (
	"github.com/joegasewicz/gomek"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	var data = struct {
		Health string
	}{
		Health: "OK",
	}
	gomek.JSON(w, data, http.StatusOK)
}
