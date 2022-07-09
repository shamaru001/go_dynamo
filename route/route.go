package route

import (
	"net/http"

	"github.com/crud/controller"
	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

func init() {
	Router.HandleFunc("/getprofile", controller.GetProfile).Methods("GET")
	Router.HandleFunc("/updateprofile", controller.UpdateProfile).Methods("GET")
	Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
}
