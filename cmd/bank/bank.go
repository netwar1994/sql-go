package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/netwar1994/sql-go/cmd/bank/app"
	"github.com/netwar1994/sql-go/pkg/card"
	"log"
	"net"
	"net/http"
	"os"
)

const defaultPort = "9999"
const defaultHost = ""

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = defaultPort
	}

	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = defaultHost
	}

	log.Printf("Server run on http://%s:%s", host, port)

	if err := execute(net.JoinHostPort(host, port)); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func execute(addr string) (err error) {
	dsn := "postgres://app:pass@localhost:5432/db"
	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		log.Println(err)
		return
	}
	defer pool.Close()

	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Release()

	cardSvc := card.NewService()
	mux := http.NewServeMux()
	application := app.NewServer(cardSvc, mux, ctx, conn)
	application.Init()

	server := &http.Server{
		Addr: addr,
		Handler: application,
	}
	return server.ListenAndServe()
}