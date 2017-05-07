# openapiro-go
Go package for accessing openapi.ro

## Installation

```bash
go get github.com/fionita/openapiro-go
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/fionita/openapiro-go"
)

func main() {
	client, err := openapi.Init(
		&openapi.Config{
			Token: "<YOUR_TOKEN>",
		},
	)
	if err != nil {
		fmt.Printf("%v", err)
	}
	company, err := client.Companies("<CIF>")
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v", company)
}
```
