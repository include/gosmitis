package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/naoina/toml"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

// Config represents toml's configuration file
type Config struct {
	EmailAccount []struct {
		From string
		To   string
		Pass string
	}
	EmailRoute []struct {
		Endpoint string
	}
}

// Dat var
var Dat Config

// Contact represents body email...
type Contact struct {
	From string `json:"from"`
	Body string `json:"body"`
}

func sendEmail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

	var f, p, t string
	f = Dat.EmailAccount[0].From
	p = Dat.EmailAccount[0].Pass
	t = Dat.EmailAccount[0].To

	decoder := json.NewDecoder(r.Body)

	var c Contact
	var err error
	err = decoder.Decode(&c)
	if err != nil {
		panic(err)
	}

	msg := "From: " + f + "\n" +
		"To: " + t + "\n" +
		"Subject: Hello there\n\n" +
		"From: " + c.From + "\n\n" +
		"Message: " + c.Body

	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", f, p, "smtp.gmail.com"),
		f, []string{t}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	log.Print("sent...")
}

func main() {

	f, err := os.Open("config-apisendemail.toml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	if err := toml.Unmarshal(buf, &Dat); err != nil {
		panic(err)
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/"+Dat.EmailRoute[0].Endpoint, sendEmail)

	log.Fatal(http.ListenAndServe(":8080", router))
}
