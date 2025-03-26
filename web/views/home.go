package views

import (
	"github.com/joegasewicz/gomek"
	"net/http"
)

type HomeView struct{}

func (h *HomeView) Get(w http.ResponseWriter, request *http.Request, d *gomek.Data) {

}

func (h *HomeView) Post(w http.ResponseWriter, request *http.Request, d *gomek.Data) {
	panic("implement me")
}

func (h *HomeView) Delete(w http.ResponseWriter, request *http.Request, d *gomek.Data) {
	panic("implement me")
}
func (h *HomeView) Put(w http.ResponseWriter, request *http.Request, d *gomek.Data) {
	panic("implement me")
}

func GetHome(w http.ResponseWriter, request *http.Request, d *gomek.Data) {

}
