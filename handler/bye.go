package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Bye struct {
	l *log.Logger
}

func NewBye(l *log.Logger) *Bye {
	return &Bye{l}
}
func (h *Bye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Bye")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("WELL"))
		return
	}
	fmt.Fprintf(rw, "Bye %s", d)
}
