package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"

	"sewik/es"
	"sewik/index/sewik"
	"sewik/sync"
)

func main() {
	log.SetFlags(0)

	var r map[string]interface{}

	// Initialize a client with the default settings.
	//
	// An `ELASTICSEARCH_URL` environment variable will be used when exported.
	//
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 1. Get cluster info
	//
	res, err := client.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print client and server version numbers.
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	log.Println(strings.Repeat("~", 37))

	// 2. Index documents concurrently
	wg := sync.LimitingWaitGroup{Limit: 100}
	flag.Parse()
	for event := range sewik.EventChannel(flag.Args(), 5) {
		wg.Add(1)
		go index(&wg, es.NewDoc(event), client)
	}
	wg.Wait()

	log.Println(strings.Repeat("-", 37))

	// 3. Search for the indexed documents
	//
	// Build the request body.
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

	// Perform the search request.
	res, err = client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex("sewik"),
		client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	log.Println(strings.Repeat("=", 37))
}

// Client: 8.0.0-SNAPSHOT
// Server: 8.0.0-SNAPSHOT
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// [201 Created] updated; version=1
// [201 Created] updated; version=1
// -------------------------------------
// [200 OK] 2 hits; took: 5ms
//  * ID=1, map[title:Test One]
//  * ID=2, map[title:Test Two]
// =====================================

func index(wg sync.WaitGroup, doc *es.Document, client *elasticsearch.Client) {
	defer wg.Done()

	//bufio.NewReader(os.Stdin).ReadBytes('\n')

	// Set up the request object.
	req := esapi.IndexRequest{
		Index:      "sewik",
		DocumentID: doc.Id,
		Body:       strings.NewReader(doc.Body),
		Refresh:    "true",
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), client)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("%s\tError parsing the response body: %s", doc.Id, err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("%s\t[%s] %s", doc.Id, res.Status(), r["error"])
		}
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("%s\tError parsing the error response body: %s", doc.Id, err)
		} else {
			log.Printf("%s\t[%s] %s; version=%d", doc.Id, res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
}
