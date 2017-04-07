package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	youtube "google.golang.org/api/youtube/v3"
	"log"
	"os"
)

type KeyConfig struct {
	Key ClientInfo
}

type ClientInfo struct {
	Id     string
	Secret string
}

var (
	channelId = flag.String("c", "", "channel id")
)

func main() {

	flag.Parse()

	var keyconfig KeyConfig
	_, err := toml.DecodeFile("./config/cnf.tml", &keyconfig)
	if err != nil {
		panic(err)
	}

	log.Printf("Clientid : %s", keyconfig.Key.Id)
	log.Printf("Clientsecret : %s", keyconfig.Key.Secret)

	if keyconfig.Key.Id == "" || keyconfig.Key.Secret == "" {
		log.Fatalf("param error")
		os.Exit(1)
	}

	// OAuth2用のconfig
	config := &oauth2.Config{
		ClientID:     keyconfig.Key.Id,
		ClientSecret: keyconfig.Key.Secret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{youtube.YoutubeScope},
	}

	ctx := context.Background()
	c := newOAuthClient(ctx, config)

	service, err := youtube.New(c)
	if err != nil {
		log.Fatalf("Unable to create YouTube service: %v", err)
	}

	if *channelId == "" {
		log.Fatalf("must channel id")
		os.Exit(1)
	}

	db = MyDB{}
	db.Connect()
	defer db.Close()

	nextPageToken := ""
	for {
		response, err := service.Search.List("snippet").
			ChannelId(*channelId).
			MaxResults(50).
			PageToken(nextPageToken).
			Do()

		if err != nil {
			// The channels.list method call returned an error.
			log.Fatalf("Error making API call to list channels: %v", err.Error())
		}

		for _, item := range response.Items {
			id := item.Id.VideoId
			title := item.Snippet.Title
			publishedAt := item.Snippet.PublishedAt
			thumbnails := item.Snippet.Thumbnails.High.Url
			Insert(Videos{Id: id, Title: title, Published_at: publishedAt, Thumbnails: thumbnails})
		}

		nextPageToken = response.NextPageToken
		if nextPageToken == "" {
			break
		}
	}
}
