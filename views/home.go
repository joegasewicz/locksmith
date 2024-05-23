package views

import (
	"github.com/joegasewicz/gomek"
	"net/http"
)

type Home struct{}

func (h *Home) Post(w http.ResponseWriter, request *http.Request, d *gomek.Data) {
	panic("implement me")
}

func (h *Home) Put(w http.ResponseWriter, request *http.Request, d *gomek.Data) {
	panic("implement me")
}

func (h *Home) Delete(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	panic("implement me")
}

func (h *Home) Get(w http.ResponseWriter, r *http.Request, data *gomek.Data) {

}
