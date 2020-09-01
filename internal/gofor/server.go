package gofor

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (g *Gofor) StartServer() {
	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(catchAllHandler)
	http.ListenAndServe(fmt.Sprintf(":%v", g.ListenPort), r)
}

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
	re, err := GoFor.Get(r.URL)
	if err != nil {
		fmt.Println(err)
	}
	if re != nil {
		w.Write(re)
	}
}
