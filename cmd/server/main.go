package main

import (
	"github.com/Azatik1000/distsys-hw/internal/pkg/server"
	"github.com/Azatik1000/distsys-hw/internal/pkg/storage"
)

func main() {
	s := server.NewServer(&storage.Memory{})
	s.Run()
}
