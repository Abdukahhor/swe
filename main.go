package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/abdukahhor/swe/app"
	"github.com/abdukahhor/swe/handlers/rpc"
	"github.com/abdukahhor/swe/storage"
)

func main() {
	var (
		addr   = flag.String("addr", ":8087", "ip:port of server")
		dbpath = flag.String("db", "sdb", "Folder path of the embedded database")
	)

	flag.Parse()

	db, err := storage.Connect(*dbpath)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	//init business layer
	c := app.New(db)
	//init rpc
	rpc.Register(c)

	go func() {
		sigint := make(chan os.Signal)
		signal.Notify(sigint, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
		log.Println("server received signal ", <-sigint)

		rpc.Close()
		os.Exit(0)
	}()

	log.Println("server started at ", *addr)
	//serve rpc
	err = rpc.Run(*addr)
	if err != nil {
		log.Fatalln(err)
	}
}
