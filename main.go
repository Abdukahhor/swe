package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/abdukahhor/swe/app"
	"github.com/abdukahhor/swe/handlers/rpc"
	"github.com/abdukahhor/swe/storage"
	"google.golang.org/grpc"
)

func main() {
	var (
		addr   = flag.String("addr", ":8087", "ip:port of server")
		dbpath = flag.String("db", "sdb", "Folder path of the embedded database")
	)

	db, err := storage.Connect(*dbpath)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	l, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	s := grpc.NewServer()
	c := app.New(db)
	rpc.Register(s, c)

	go func() {
		sigint := make(chan os.Signal)
		signal.Notify(sigint, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
		log.Println("server received signal ", <-sigint)

		s.GracefulStop()
		os.Exit(0)
	}()

	log.Println("server started at ", *addr)
	err = s.Serve(l)
	if err != nil {
		log.Fatalln(err)
	}
}
