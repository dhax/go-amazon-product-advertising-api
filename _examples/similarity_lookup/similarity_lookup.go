package main

import (
	"fmt"
	"log"

	"github.com/jami/go-amazon-product-advertising-api/amazon"
)

func main() {
	client, err := amazon.NewFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.SimilarityLookup(amazon.SimilarityLookupParameters{
		ResponseGroups: []amazon.SimilarityLookupResponseGroup{
			amazon.SimilarityLookupResponseGroupLarge,
		},
		ItemIDs: []string{
			"477418392X",
		},
	}).Do()
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range res.Items.Item {
		fmt.Printf(`-------------------------------
[Title] %v
[URL]   %v
`, item.ItemAttributes.Title, item.DetailPageURL)
	}
}
