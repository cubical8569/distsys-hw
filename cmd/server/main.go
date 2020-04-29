package main

import (
	"github.com/Azatik1000/distsys-hw/internal/pkg/server"
	"github.com/Azatik1000/distsys-hw/internal/pkg/storage"
)

func main() {
	db, err := storage.NewDB()
	if err != nil {
		panic(err)
	}

	s := server.NewServer(db)
	s.Run()
}
