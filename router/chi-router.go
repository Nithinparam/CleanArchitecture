package router

import (
	"fmt"
	"net/http"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)

}
func (*chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)

}
func (*chiRouter) SERVE(port string) {
	fmt.Println("Chi http server running on port ", port)
	http.ListenAndServe(port, chiDispatcher)

}
