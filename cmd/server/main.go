//go:generate kek
package main

import (
	"github.com/Azatik1000/distsys-hw/internal/pkg/server"
	"github.com/Azatik1000/distsys-hw/internal/pkg/storage"
)

// @title Online Store API

// @contact.name Azat Kalmykov

// @host localhost:8080
// @BasePath /
func main() {
	db, err := storage.NewDB()
	if err != nil {
		panic(err)
	}

	s := server.NewServer(db)

	s.Run()
}
