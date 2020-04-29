package main

import (
	"github.com/Azatik1000/distsys-hw/internal/pkg/server"
	"github.com/Azatik1000/distsys-hw/internal/pkg/storage"
)

func main() {
	s := server.Server(&storage.Memory{})
	s.Run()
}
