go-amazon-product-advertising-api
=================================

This is a fork of ngs/go-amazon-product-advertising-api
Since a year or so ngs did not accept or merge any pull requests. So this is a fork which merges all the open pull requests. Because some of them are critical and mandatory!!

12.08.19 - Can't get any contact to ngs. Please find a new maintainer!

[![Build Status](https://travis-ci.org/ngs/go-amazon-product-advertising-api.svg?branch=master)](https://travis-ci.org/ngs/go-amazon-product-advertising-api)
[![GoDoc](https://godoc.org/github.com/ngs/go-amazon-product-advertising-api/amazon?status.svg)](https://godoc.org/github.com/ngs/go-amazon-product-advertising-api/amazon)
[![Go Report Card](https://goreportcard.com/badge/github.com/ngs/go-amazon-product-advertising-api)](https://goreportcard.com/report/github.com/ngs/go-amazon-product-advertising-api)
[![Coverage Status](https://coveralls.io/repos/github/ngs/go-amazon-product-advertising-api/badge.svg?branch=master)](https://coveralls.io/github/ngs/go-amazon-product-advertising-api?branch=master)

Go Client Library for [Amazon Product Advertising API]

How to Use
----------

```sh
go get -u github.com/ngs/go-amazon-product-advertising-api/amazon
```

```go
package main

import (
	"fmt"
	"log"

	"github.com/ngs/go-amazon-product-advertising-api/amazon"
)

func main() {
	client, err := amazon.NewFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.ItemSearch(amazon.ItemSearchParameters{
		SearchIndex: amazon.SearchIndexBooks,
		Keywords:    "Go 言語",
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
```

```sh
export AWS_ACCESS_KEY_ID=${YOUR_AWS_ACCESS_KEY_ID}
export AWS_SECRET_ACCESS_KEY=${YOUR_AWS_SECRET_ACCESS_KEY}
export AWS_PRODUCT_REGION=JP
export AWS_ASSOCIATE_TAG=ngsio-22

go run item_search.go
```

[Amazon Product Advertising API]: https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html

## Author

[Atsushi Nagase]

## License

See [LICENSE]

[Atsushi Nagase]: https://ngs.io
[LICENSE]: LICENSE
