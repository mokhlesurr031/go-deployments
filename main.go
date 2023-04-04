package main

import (
	"log"

	"github.com/mokhlesur-rahman/golang-basic-crud-api-server/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Println(err)
	}
}
