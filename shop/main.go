//go:generate kek
package main

import (
	"fmt"
	"github.com/Azatik1000/distsys-hw/shop/server"
	"github.com/Azatik1000/distsys-hw/shop/storage"
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

func setupDB() (storage.Storage, error) {
	dbIP, present := os.LookupEnv("DB_IP")
	if !present {
		return nil, fmt.Errorf("database ip not specified")
	}

	dbPort, present := os.LookupEnv("DB_PORT")
	if !present {
		return nil, fmt.Errorf("database port not specified")
	}

	dbUser, present := os.LookupEnv("DB_USER")
	if !present {
		return nil, fmt.Errorf("database user not specified")
	}

	dbPassword, present := os.LookupEnv("DB_PASSWORD")
	if !present {
		return nil, fmt.Errorf("database password not specified")
	}

	dbName, present := os.LookupEnv("DB_NAME")
	if !present {
		return nil, fmt.Errorf("database name not specified")
	}

	//time.Sleep(time.Second * 200)

	db, err := storage.NewDB(dbIP, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := setupDB()
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
