//go:generate kek
package main

import (
	"github.com/Azatik1000/distsys-hw/internal/pkg/server"
	"github.com/Azatik1000/distsys-hw/internal/pkg/storage"
	"os"
	"os/signal"
	"sync"
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
	defer func() {
		_ = db.Close()
	}()

	s := server.NewServer(db)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.Run()
	}()

	<-signals

	err = s.Shutdown()
	if err != nil {
		panic(err)
	}

	wg.Wait()
}
