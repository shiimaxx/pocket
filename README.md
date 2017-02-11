# pocket

Pocket API v3 Client

[![travis status](https://travis-ci.org/shiimaxx/pocket.svg?branch=master)](https://travis-ci.org/shiimaxx/pocket.svg?branch=master)

## Installation

```
go get github.com/shiimaxx/pocket
```

## Usage

1. Create an Application
2. Get access token

    ```
    pocket auth -c <cousumer key>
    ```
3. Retrieve Sample

    ```
    client, _ := pocket.NewClient("<consumer key>", "<access token>")

    items, err := client.Retrieve(&pocket.RetrieveOpts{
            Count:    10,
            Favorite: false,
            Sort:     "oldest",
    })

    for _, item := range items.List {
            fmt.Println(item.ResolvedTitle)
            fmt.Println(item.GivenURL)
            fmt.Println(item.ItemID)
    ```
