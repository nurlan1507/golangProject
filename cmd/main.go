package main

import (
	"github.com/joho/godotenv"
	"testApp"
	"testApp/pkg"
)

func init() {

}
func main() {
	err := godotenv.Load(".env")

	server := &testApp.Server{}

	err = server.RunServer(":4000", pkg.Routes())

	if err != nil {
		return
	}
}
