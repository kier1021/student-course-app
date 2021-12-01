package main

import (
	"log"

	"github.com/kier1021/student-course-app/server"
	_ "github.com/lib/pq"
)

func main() {

	apiServer, err := server.NewAPIServer()
	if err != nil {
		log.Fatal(err)
	}

	err = apiServer.Run()

	if err != nil {
		log.Fatal(err)
	}
}
