package main

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"

	"net/http"

	"gopkg.in/mgo.v2"

	"go-sse-rest/controllers"
)

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017/sse")
	if err != nil {
		log.Fatal(err)
	}
	return session
}

func main() {
	log.Print("go-sse")

	router := gin.Default()

	ctrl := controllers.NewSseController()

	router.GET("/:user", ctrl.GetUser)

	http.ListenAndServe(":8080", router)
}
