package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	w.Write([]byte(fmt.Sprintf("Hello %s", name)))
}
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Привет. Это главная страница")
}
func main() {
	router := mux.NewRouter()
	s := http.StripPrefix("/lol", http.FileServer(http.Dir("./swapi/")))
	//router.HandleFunc("/", Index)
	router.PathPrefix("/lol").Handler(s)
	http.Handle("/", router)
	http.ListenAndServe(":5000", nil)

	fmt.Println("xxx")
	// router := httprouter.New()
	// router.GET("/:name", IndexHandler)
	// listener, err := net.Listen("tcp", "127.0.0.1:1234")
	// if err != nil {
	// 	panic(err)
	// }
	// sertver := &http.Server{
	// 	Handler:      router,
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }
	// log.Fatalln(sertver.Serve(listener))

}
