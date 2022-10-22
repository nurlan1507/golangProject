package main

import (
	"testApp"
	"testApp/pkg"
)

func init() {

}
func main() {
	server := &testApp.Server{}
	err := server.RunServer(":4000", pkg.Routes())

	if err != nil {
		return
	}
}
