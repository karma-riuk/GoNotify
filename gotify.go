package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gotify/go-api-client/v2/auth"
	"github.com/gotify/go-api-client/v2/client"
	"github.com/gotify/go-api-client/v2/client/message"
	"github.com/gotify/go-api-client/v2/gotify"
	"github.com/gotify/go-api-client/v2/models"
)

const (
	gotifyURL        = "https://notify.karma-riuk.com"
	applicationToken = "AWwrtmj8OH5BxtO"
)

var gotifyClient *client.GotifyREST

func setup() bool {
	log.Println("Setting up")
	myURL, _ := url.Parse(gotifyURL)
	gotifyClient = gotify.NewClient(myURL, &http.Client{})
	versionResponse, err := gotifyClient.Version.GetVersion(nil)
	if err != nil {
		log.Fatal("Could not request version ", err)
		return false
	}
	version := versionResponse.Payload
	log.Println("Found version", *version)
	return true
}

func Notify(r *http.Request) {
	params := message.NewCreateMessageParams()
	params.Body = &models.MessageExternal{
		Title:    "Website visited",
		Message:  fmt.Sprint(r),
		Priority: 5,
	}
	_, err := gotifyClient.Message.CreateMessage(params, auth.TokenAuth(applicationToken))
	if err != nil {
		log.Fatalf("Could not send message %v", err)
		return
	}
	log.Println("Message Sent!")
}
