package main

import "log"

func main() {
	//Gorilla Mux
	// router := NewRouter()
	// //srv setting server

	// log.Fatal(srv.ListenAndServe())

	//Gin
	eng := GetEngine()
	// srv := &http.Server{
	// 	Addr:         "192.168.90.26:1314",
	// 	Handler:      eng,
	// 	WriteTimeout: 60 * time.Second,
	// 	ReadTimeout:  60 * time.Second,
	// }
	log.Fatal(eng.Run(":1314"))
	// log.Fatal(srv.ListenAndServe())
	//endless in linux
	//endless.ListenAndServe(":1313",eng);
}
