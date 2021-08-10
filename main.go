package main

import (
	"log"
	"micro/handler"
	"net/http"
	"os"
)

func main() {
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	// log.Println("hello")
	// 	// d, err := ioutil.ReadAll(r.Body)
	// 	// if err != nil {
	// 	// 	rw.WriteHeader(http.StatusBadRequest)
	// 	// 	rw.Write([]byte("WELL"))
	// 	// 	return
	// 	// }
	// 	// log.Printf("Data %s\n", d)

	// })
	// http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
	// 	log.Println("bye")
	// })
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handler.NewHello(l)
	bb := handler.NewBye(l)
	ph := handler.NewPorducts(l)
	sm := http.NewServeMux()
	sm.Handle("/", ph)
	sm.Handle("/Hello", hh)

	sm.Handle("/Bye", bb)
	//sm.Handle("/goodbye", bb)
	http.ListenAndServe(":8080", sm)
}
