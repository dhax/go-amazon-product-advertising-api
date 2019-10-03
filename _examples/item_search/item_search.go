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
	res, err := client.ItemSearch(amazon.ItemSearchParameters{
		SearchIndex:    amazon.SearchIndexMusic,
		ResponseGroups: []amazon.ItemSearchResponseGroup{amazon.ItemSearchResponseGroupLarge},
		Keywords:       "Pat Metheny",
	}).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d results found\n\n", res.Items.TotalResults)
	for _, item := range res.Items.Item {
		fmt.Printf(`-------------------------------
[Title] %v
[URL]   %v
`, item.ItemAttributes.Title, item.DetailPageURL)
	}
}
