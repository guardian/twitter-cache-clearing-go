package main

import (
	"context"
	"errors"
	"io/ioutil"
	"log"

	"encoding/json"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

// See https://docs.aws.amazon.com/lambda/latest/dg/go-programming-model-handler-types.html

func main() {
	lambda.Start(handle)
}

func handle(ctx context.Context, event events.KinesisEvent) error {
	lc, _ := lambdacontext.FromContext(ctx)

	// Access event example
	log.Printf("Lambda invoked at %s", event.Records)

	// Access context example
	log.Printf("Env vars are: %v", lc.ClientContext.Env)

	// get CAPI ID
	// call CAPI to check if recently updated
	// call Ophan to get shared URLs for content ID
	// purge Twitter for shared URLs
	return nil
}

type content struct {
	id                 string
	webPublicationDate string
	fields             struct {
		lastModified string
	}
}

func getFromCAPI(contentID string) (content, error) {
	var capiContent content

	capiURL := "https://content.guardianapis.com/world/live/2020/feb/27/coronavirus-news-live-updates?api-key=frontend&show-fields=lastModified,webPublicationDate"
	resp, err := http.Get(capiURL)
	if err != nil {
		return capiContent, errors.New("CAPI down")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &capiContent)
	if err != nil {
		return capiContent, errors.New("json unmarshal failed" + err.Error())
	}

	return capiContent, nil
}

func isRecentlyUpdated(webPublicationDate, lastModified string) (bool, error) {
	publicationDate, _ := time.Parse(time.RFC3339, webPublicationDate)
	lastModifiedDate, _ := time.Parse(time.RFC3339, lastModified)

	return lastModifiedDate.Before(publicationDate.Add(5 * time.Minute)), nil
}

func fail() {

}

func getSharedURLs(contentID string) []string {
	return nil
}

func purgeTwitter(URLs []string) error {
	return nil
}
