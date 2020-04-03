package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"

	"sewik/pkg/dom"
	"sewik/pkg/es"
	"sewik/pkg/sewik"
	"sewik/pkg/sync"
	"sewik/pkg/sys"
)

func main() {
	log.SetFlags(0)
	var r map[string]interface{}
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 1. Get cluster info
	res, err := client.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	log.Println(strings.Repeat("~", 37))

	// 2. Index documents concurrently
	wg := sync.LimitingWaitGroup{Limit: 50}
	flag.Parse()
	for event := range sewik.ElementsOf("ZDARZENIE", sys.Filenames(flag.Args(), 10000), 10) {
		catcher(&wg, event, client)
	}
	wg.Wait()
	log.Println(strings.Repeat("-", 37))

	// 3. Search for the indexed documents
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "test",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	log.Println(strings.Repeat("=", 37))
}

func catcher(wg *sync.LimitingWaitGroup, event *dom.Element, client *elasticsearch.Client) {
	doc := es.NewDoc(event, "")
	defer func() {
		wg.Done()
		if r := recover(); r != nil {
			log.Printf("%s\t %s\n%s", doc.ID, strings.ReplaceAll(r.(string), "\n", " "), doc.Body)
			fmt.Printf(`{"id":"%s","recovery":"%s","body":%s}`+"\n", doc.ID, strings.ReplaceAll(r.(string), "\n", " "), doc.Body)
		}
	}()
	wg.Add(1)
	go index(wg, doc, client)
}

func index(wg sync.WaitGroup, doc *es.Document, client *elasticsearch.Client) {
	defer wg.Done()

	req := esapi.IndexRequest{
		Index:      "sewik",
		DocumentID: doc.ID,
		Body:       strings.NewReader(doc.Body),
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), client)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("%s\tError parsing the response body: %s", doc.ID, err)
			fmt.Printf(`{"id":"%s","errorDecodingError":"%s","body":%s}`+"\n", doc.ID, strings.ReplaceAll(err.Error(), "\n", " "), doc.Body)
		} else {
			// Print the response status and indexed document version.
			log.Printf("%s\t[%s] %s", doc.ID, res.Status(), r["error"])
			fmt.Printf(`{"id":"%s","errorHTTP":"%s","body":%s}`+"\n", doc.ID, strings.ReplaceAll(r["error"].(string), "\n", " "), doc.Body)
		}
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("%s\tError parsing the error response body: %s", doc.ID, err)
			fmt.Printf(`{"id":"%s","errorDecodingSuccess":"%s","body":%s}`+"\n", doc.ID, strings.ReplaceAll(err.Error(), "\n", " "), doc.Body)
		} else {
			v := int(r["_version"].(float64))
			//log.Printf("%s\t[%s] %s; version=%d", doc.Id, res.Status(), r["result"], v)
			if v > 1 {
				fmt.Printf(`{"id":"%s","version":%d,"body":%s}`+"\n", doc.ID, v, doc.Body)
			}
		}
	}
}
