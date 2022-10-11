package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	pb := pocketbase.New()

	pb.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Static("/", "static")
		return nil
	})

	if err := pb.Start(); err != nil {
		log.Fatal(err)
	}

}
