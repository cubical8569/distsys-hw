//go:generate kek
package main

import (
	"github.com/Azatik1000/distsys-hw/internal/pkg/server"
	"github.com/Azatik1000/distsys-hw/internal/pkg/storage"
)


// TODO: make host annotation adopt port
// @Title Online Store API

// @Contact.name Azat Kalmykov

// @Host localhost:3333

// @Tag.name products

// @BasePath /
func main() {
	db, err := storage.NewDB()
	if err != nil {
		panic(err)
	}

	s := server.NewServer(db)

	s.Run()
}
