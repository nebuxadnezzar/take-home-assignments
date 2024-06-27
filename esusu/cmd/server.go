package main

import (
	"fmt"
	db "maas/api/database"
	"maas/api/route"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const (
	DEFAULT_PORT = "8080"
)

func main() {
	fmt.Println("Starting...")
	port := DEFAULT_PORT
	l := len(os.Args)
	if l > 1 {
		if p, e := strconv.Atoi(os.Args[1]); e == nil {
			port = fmt.Sprintf("%d", p)
		} else {
			fmt.Fprintf(os.Stderr, "first optional argument must be numeric port number, you passed %s\n", os.Args[1])
		}
	}
	userdbpath := ``
	if l > 2 {
		userdbpath = os.Args[2]
	}
	svr, custom := route.NewEchoWrapperRouter(userdbpath)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	go func() {
		<-sigs
		done <- true
	}()
	go func() {
		svr.Logger.Fatal(svr.Start(":" + port))
	}()
	defer cleanup(custom)
	<-done
}

func cleanup(custom map[string]interface{}) {
	fmt.Printf("Shutting down...\n")
	usersdb := custom[db.USERSDB_KEY].(db.Database)
	fmt.Printf("%v\n", usersdb.Flush())
}
