package main

import (
	"log"

	"github.com/AlbinoDrought/imap-password-logger/password"
	"github.com/emersion/go-imap/server"
)

func main() {
	backend := password.New(func(username, password string) {
		log.Println(username, password)
	})
	s := server.New(backend)
	s.Addr = ":1143"
	s.AllowInsecureAuth = true
	log.Println("Starting IMAP server at localhost:1143")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
